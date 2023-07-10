package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/go-gin-template/internal/app/template/module/auth"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg/constant"
	"github.com/tanveerprottoy/go-gin-template/pkg/response"
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
		e, err := m.Service.Authorize(ctx)
		if err != nil {
			// send error response
			response.RespondError(http.StatusForbidden, err, ctx)
			// need to abort the middleware chain
			ctx.Abort()
			return
		}
		ctx.Set(constant.KeyAuthUser, e)
		ctx.Next()
	}
}
