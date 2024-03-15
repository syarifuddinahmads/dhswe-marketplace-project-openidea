package factory

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/auth"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/bank_account"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/todo"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/middleware"
)

type HandlerFactory struct {
	Router        *mux.Router
	Logger        *logrus.Logger
	Database      *sqlx.DB
	VersionPrefix string
}

func NewHandlerFactory(router *mux.Router, logger *logrus.Logger, db *sqlx.DB, versionPrefix string) *HandlerFactory {
	return &HandlerFactory{
		Router:        router,
		Logger:        logger,
		Database:      db,
		VersionPrefix: versionPrefix,
	}
}

func (hf *HandlerFactory) RegisterHandlers() {
	v1AuthRouter := hf.Router.PathPrefix("/" + hf.VersionPrefix + "/").Subrouter()
	userHandler := auth.NewHandler(v1AuthRouter, hf.Logger, hf.Database)
	userHandler.RegisterRoutes()

	v1Router := hf.Router.PathPrefix("/" + hf.VersionPrefix + "/").Subrouter()
	v1Router.Use(middleware.JwtMiddleware())
	todoHandler := todo.NewHandler(v1Router, hf.Logger, hf.Database)
	todoHandler.RegisterRoutes()

	v1BankRouter := hf.Router.PathPrefix("/" + hf.VersionPrefix + "/").Subrouter()
	v1BankRouter.Use(middleware.JwtMiddleware())
	bankAccountHandler := bank_account.NewHandler(v1BankRouter, hf.Logger, hf.Database)
	bankAccountHandler.RegisterRoutes()
}
