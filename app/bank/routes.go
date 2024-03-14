package bank

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	bankRepo "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/bank/repository"
	bankService "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/bank/service"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	bankService bankService.Service
}

func NewHandler(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) service {

	// Initialize your todo service
	bankSvc := bankService.NewService(bankRepo.NewRepository(db))

	// Create a new instance of the service
	svc := service{
		logger:      lg,
		router:      r,
		bankService: bankSvc,
	}

	// Register routes
	svc.RegisterRoutes()

	// Register MiddlewareLogger middleware
	svc.router.Use(utils.MiddlewareLogger)

	return svc
}

func (s *service) RegisterRoutes() {

	fmt.Println("Tesss !!!")

	// s.router.HandleFunc("/todos", s.Get()).Methods("GET")
	// s.router.HandleFunc("/todos/{id}", s.Get()).Methods("GET")
	s.router.HandleFunc("/bank-account", s.Create()).Methods("POST")
	// s.router.HandleFunc("/todos/{id}", s.Update()).Methods("PUT")
	// s.router.HandleFunc("/todos/{id}", s.Delete()).Methods("DELETE")
}
