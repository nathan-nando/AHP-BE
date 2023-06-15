package collection

import "github.com/labstack/echo/v4"

func (h *Handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PATCH("", h.Update)
	g.DELETE("/:id", h.Delete)
}
