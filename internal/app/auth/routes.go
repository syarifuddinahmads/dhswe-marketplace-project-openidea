package auth

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	userService Service
}

func NewHandler(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) service {

	// Initialize your todo service
	userService := NewService(repository.NewRepository(db))

	// Create a new instance of the service
	svc := service{
		logger:      lg,
		router:      r,
		userService: userService,
	}

	// Register routes
	svc.RegisterRoutes()
	// Register MiddlewareLogger middleware
	svc.router.Use(utils.MiddlewareLogger)

	return svc
}

func (s *service) RegisterRoutes() {
	s.router.HandleFunc("/login", s.Login).Methods("POST")
	s.router.HandleFunc("/register", s.Register).Methods("POST")
}
