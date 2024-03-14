package factory

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/auth"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/todo"
)

type HandlerFactory struct {
	Router   *mux.Router
	Logger   *logrus.Logger
	Database *sqlx.DB
}

func NewHandlerFactory(router *mux.Router, logger *logrus.Logger, db *sqlx.DB) *HandlerFactory {
	return &HandlerFactory{
		Router:   router,
		Logger:   logger,
		Database: db,
	}
}

func (hf *HandlerFactory) RegisterHandlers() {
	todoHandler := todo.NewHandler(hf.Router, hf.Logger, hf.Database)
	todoHandler.RegisterRoutes()

	userHandler := auth.NewHandler(hf.Router, hf.Logger, hf.Database)
	userHandler.RegisterRoutes()
}
