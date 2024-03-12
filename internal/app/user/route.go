package user

import (
	"github.com/labstack/echo/v4"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/middleware"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get, middleware.Authentication)
	g.GET("/:id", h.GetByID, middleware.Authentication)
}
