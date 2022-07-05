package service

import "github.com/gin-gonic/gin"

type paymentServiceAccessor interface {
	HandleRequest(ctx *gin.Context)
}

type accountServiceAccessor interface {
	HandleRequest(ctx *gin.Context)
}

// Accessor is interface which defines all functions used on service routing
// in the system. Any other packages should use this Accessor instead of using
// eg. service directly.
type Accessor interface {
	paymentServiceAccessor
	accountServiceAccessor
}
