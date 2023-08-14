package subCriteria

import "github.com/labstack/echo/v4"

func (h *Handler) Route(g *echo.Group) {
	g.GET("/sub-criteria", h.GetSubCriteria)
	g.PATCH("/sub-criteria", h.UpdateSubCriteria)
	g.GET("/sub-criteria/check/:mode", h.CheckConsistencyRatio)
}
