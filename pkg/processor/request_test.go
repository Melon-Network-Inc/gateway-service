package processor

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUserData(t *testing.T) {
	testContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name     string
		setup    func()
		teardown func()
		args     args
		want     map[string]string
		want1    bool
	}{
		{
			name: "full config",
			setup: func() {
				testContext.Set(UsernameKey, "test")
				testContext.Set(UserIDKey, uint(1))
				testContext.Set(AuthorizationKey, "token")
				testContext.Set(ContextRoleKey, "1")
				testContext.Request = httptest.NewRequest("GET", "http://localhost:8080", nil)
			},
			teardown: func() {
				testContext, _ = gin.CreateTestContext(httptest.NewRecorder())
			},
			args: args{
				ctx: testContext,
			},
			want: map[string]string{
				ContextUserKey:   "test",
				ContextUserIDKey: "1",
				AuthorizationKey: "token",
				ContextRoleKey:   "1",
				ContextClientIP:  "192.0.2.1",
				ContextRequestID: "",
			},
			want1: true,
		},
		{
			name: "registration token",
			setup: func() {
				testContext.Set(RegistrationKey, "registration")
				testContext.Request = httptest.NewRequest("GET", "http://localhost:8080", nil)
			},
			teardown: func() {
				testContext, _ = gin.CreateTestContext(httptest.NewRecorder())
			},
			args: args{
				ctx: testContext,
			},
			want: map[string]string{
				ContextRegistrationTokenKey: "registration",
				ContextRoleKey:              "1",
				ContextClientIP:             "192.0.2.1",
				ContextRequestID:            "",
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.setup()
			got, got1 := GetUserData(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetUserData() got1 = %v, want %v", got1, tt.want1)
			}
			tt.teardown()
		})
	}
}
