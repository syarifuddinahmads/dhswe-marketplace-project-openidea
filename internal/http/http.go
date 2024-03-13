package http

import (
	"github.com/labstack/echo/v4"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/auth"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/product"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/user"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/factory"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1Group := e.Group("v1")
	auth.NewHandler(f).Route(v1Group.Group("/user"))
	product.NewHandler(f).Route(v1Group.Group("/products"))
	user.NewHandler(f).Route(v1Group.Group("/users"))
}
