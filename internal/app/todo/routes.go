package todo

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
	toDoService Service
}

func NewHandler(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) service {

	// Initialize your todo service
	todoSvc := NewService(repository.NewRepository(db))

	// Create a new instance of the service
	svc := service{
		logger:      lg,
		router:      r,
		toDoService: todoSvc,
	}

	// Register routes
	svc.RegisterRoutes()
	// Register MiddlewareLogger middleware
	svc.router.Use(utils.MiddlewareLogger)

	return svc
}

func (s *service) RegisterRoutes() {
	s.router.HandleFunc("/todos", s.GetAll()).Methods("GET")
	s.router.HandleFunc("/todos/{id}", s.Get()).Methods("GET")
	s.router.HandleFunc("/todos", s.Create()).Methods("POST")
	s.router.HandleFunc("/todos/{id}", s.Update()).Methods("PUT")
	s.router.HandleFunc("/todos/{id}", s.Delete()).Methods("DELETE")
}
