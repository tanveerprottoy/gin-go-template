package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/templates-go-gin/internal/app/module/user/entity"
	"github.com/tanveerprottoy/templates-go-gin/internal/app/module/user/handler"
	"github.com/tanveerprottoy/templates-go-gin/internal/app/module/user/repository"
	"github.com/tanveerprottoy/templates-go-gin/internal/app/module/user/service"
	sqlxPkg "github.com/tanveerprottoy/templates-go-gin/pkg/data/sql/sqlx"

	"github.com/go-playground/validator/v10"
)

type Module struct {
	Handler         *handler.Handler
	Service         *service.Service
	Repository      sqlxPkg.Repository[entity.User]
}

func NewModule(db *sqlx.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = repository.NewRepository(db)
	m.MongoRepository = repository.NewRepositoryAlt(db)
	m.Service = service.NewService(m.Repository)
	m.Handler = handler.NewHandler(m.Service, validate)
	return m
}
