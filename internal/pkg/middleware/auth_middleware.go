package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/go-gin-template/internal/app/template/module/auth"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg/constant"
)

type AuthMiddleware struct {
	Service *auth.Service
}

func NewAuthMiddleware(s *auth.Service) *AuthMiddleware {
	m := new(AuthMiddleware)
	m.Service = s
	return m
}

// AuthUserMiddleWare
func (m *AuthMiddleware) AuthUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := m.Service.Authorize(ctx)
		if payload == nil {
			// need to abort the middleware chain
			ctx.Abort()
			return
		}
		ctx.Set(constant.KeyAuthUser, payload)
		ctx.Next()
	}
}
