package http

import (
	"github.com/labstack/echo/v4"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/auth"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/product"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/app/user"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/factory"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	auth.NewHandler(f).Route(e.Group("/auth"))
	product.NewHandler(f).Route(e.Group("/products"))
	user.NewHandler(f).Route(e.Group("/users"))
}
