package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg/router"
)

// Gin Engine container struct
// Engine is the framework's instance, it contains the muxer, middleware and configuration settings.
// Create an instance of Engine, by using New() or Default()
type Gin struct {
	Engine *gin.Engine
}

func NewGin() *Gin {
	g := &Gin{}
	g.Engine = gin.Default()
	// SetTrustedProxies set a list of network origins (IPv4 addresses,
	// IPv4 CIDRs, IPv6 addresses or IPv6 CIDRs) from which to trust
	// request's headers that contain alternative client IP when
	// `(*gin.Engine).ForwardedByClientIP` is `true`.
	g.Engine.SetTrustedProxies(nil)
	router.RegisterGlobalMiddlewares(g.Engine)
	return g
}
