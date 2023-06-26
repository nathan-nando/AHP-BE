package alternative

import "github.com/labstack/echo/v4"

func (h *Handler) Route(g *echo.Group) {
	g.GET("/collection/:id", h.GetsByCollectionID)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PATCH("", h.Update)
	g.DELETE("/:id", h.Delete)
}
