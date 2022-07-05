package service

import "github.com/gin-gonic/gin"

type service struct {
	paymentServiceAccessor;
	accountServiceAccessor;
}

func New(ctx *gin.Context) Accessor {
	paymentService := newPaymentService()
	accountService := newAccountService()
	return &service{
		paymentServiceAccessor: paymentService,
		accountServiceAccessor: accountService
	}
}