package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tanveerprottoy/go-gin-template/internal/app/template/module/user"
	"github.com/tanveerprottoy/go-gin-template/internal/app/template/module/user/entity"
	"github.com/tanveerprottoy/go-gin-template/pkg/httppkg"
	"github.com/tanveerprottoy/go-gin-template/pkg/jwtpkg"
)

type Service struct {
	userService *user.Service
}

func NewService(userService *user.Service) *Service {
	s := new(Service)
	s.userService = userService
	return s
}

func (s *Service) Authorize(ctx *gin.Context) (entity.User, error) {
	var u entity.User
	splits, err := httppkg.ParseAuthToken(ctx)
	if err != nil {
		return u, err
	}
	tokenBody := splits[1]
	claims, err := jwtpkg.VerifyToken(tokenBody)
	if err != nil {
		return u, err
	}
	// find user
	u, err = s.userService.ReadOneInternal(claims.Payload.Id)
	if err != nil {
		return u, err
	}
	return u, nil
}
