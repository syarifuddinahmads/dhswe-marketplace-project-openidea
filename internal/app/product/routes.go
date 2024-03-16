package product

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/repository"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

type service struct {
	logger         *logrus.Logger
	router         *mux.Router
	productService Service
}

func NewHandler(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) service {

	// Initialize your todo service
	productService := NewService(repository.NewRepository(db))

	// Create a new instance of the service
	svc := service{
		logger:         lg,
		router:         r,
		productService: productService,
	}

	// Register routes
	svc.RegisterRoutes()
	// Register MiddlewareLogger middleware
	svc.router.Use(utils.MiddlewareLogger)

	return svc
}

func (s *service) RegisterRoutes() {
	s.router.HandleFunc("/product", s.IndexProduct()).Methods("GET")
	s.router.HandleFunc("/product/{id}", s.ShowProduct()).Methods("GET")
	s.router.HandleFunc("/product", s.StoreProduct).Methods("POST")
	s.router.HandleFunc("/product/{id}", s.UpdateProduct).Methods("PATCH")
	s.router.HandleFunc("/product/{id}", s.DeleteProduct).Methods("DELETE")
}
