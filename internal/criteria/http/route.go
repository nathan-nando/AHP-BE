package criteria

import "github.com/labstack/echo/v4"

func (h *Handler) Route(g *echo.Group) {
	g.GET("/criteria", h.GetCriteria)
	g.GET("/criteria/check", h.CheckConsistencyRatio)
	g.PATCH("/criteria", h.UpdateCriteriaAlternative)
	g.GET("/scores/:collection_id", h.GetScores)
	g.GET("/final_scores/:collection_id", h.GetFinalScores)
	g.GET("/point/calculate/:collection_id", h.CalculateAlternativeToPoint)
	g.GET("/scores/calculate/:collection_id", h.CalculateScores)
	g.GET("/final_scores/calculate/:collection_id", h.CalculateFinalScores)
}
