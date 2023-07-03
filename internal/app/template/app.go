package template

import (
	"github.com/tanveerprottoy/go-gin-template/internal/app/template/module/auth"
	"github.com/tanveerprottoy/go-gin-template/internal/app/template/module/user"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg/constant"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg/middleware"
	"github.com/tanveerprottoy/go-gin-template/internal/pkg/router"
	"github.com/tanveerprottoy/go-gin-template/pkg/data/sqlxpkg"

	"github.com/go-playground/validator/v10"
)

// App struct
type App struct {
	DBClient    *sqlxpkg.Client
	gin         *pkg.Gin
	Middlewares []any
	AuthModule  *auth.Module
	UserModule  *user.Module
	Validate    *validator.Validate
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initMiddlewares() {
	authMiddleWare := middleware.NewAuthMiddleware(a.AuthModule.Service)
	a.Middlewares = append(a.Middlewares, authMiddleWare)
}

func (a *App) initModules() {
	a.UserModule = user.NewModule(a.DBClient.DB, a.Validate)
	a.AuthModule = auth.NewModule(a.UserModule.Service)
}

func (a *App) initModuleRouters() {
	m := a.Middlewares[0].(*middleware.AuthMiddleware)
	router.RegisterUserRoutes(a.gin.Engine, constant.V1, a.UserModule, m)
}

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.gin = pkg.NewGin()
	a.initModules()
	a.initMiddlewares()
	a.initModuleRouters()
}

// Run app
func (a *App) Run() {
	a.gin.Engine.Run(":8080")
}

// Run app
func (a *App) RunTLS() {
	a.gin.Engine.Run(":443", "cert.crt", "key.key")
}
