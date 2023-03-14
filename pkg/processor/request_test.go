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
		name  string
		setup func()
		args  args
		want  map[string]string
		want1 bool
	}{
		{
			name: "TestGetUserData",
			setup: func() {
				testContext.Set(UsernameKey, "test")
				testContext.Set(UserIDKey, 1)
				testContext.Set(AuthorizationKey, "token")
				testContext.Set(ContextRoleKey, "1")
				testContext.Request = httptest.NewRequest("GET", "http://localhost:8080", nil)
			},
			args: args{
				ctx: testContext,
			},
			want: map[string]string{
				ContextUserKey:    "test",
				ContextUserIDKey:  "1",
				AuthorizationKey:  "token",
				"UserRole":        "1",
				"X-Forwarded-For": "192.0.2.1",
				"X-Request-ID":    "",
			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, got1 := GetUserData(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetUserData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
