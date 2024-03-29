package criteria

import (
	"ahp-be/internal/alternative"
	"ahp-be/internal/criteria"
	"ahp-be/internal/criteria/dto"
	"ahp-be/internal/criteria/usecase"
	"ahp-be/pkg/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler(repo criteria.Repository, repoAlternative alternative.Repository, db *gorm.DB) *Handler {
	service := usecase.New(repo, repoAlternative, db)
	return &Handler{service}
}

// GetCriteria
// @Summary Get All Criteria Alternative
// @Description Get ALl Criteria Alternative
// @Tags Criteria
// @Accept json
// @Produce json
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/criteria [get]
func (h *Handler) GetCriteria(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.FindCriteria(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// CheckConsistencyRatio
// @Summary Check Consistency Ratio
// @Description Check Consistency Ratio
// @Tags Criteria
// @Accept json
// @Produce json
// @Success 200 {object} dto.CheckConsistencyRatioResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/criteria/check [get]
func (h *Handler) CheckConsistencyRatio(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.CheckConsistencyRatio(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetScores
// @Summary Get Scores
// @Description Get Scores
// @Tags Criteria
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/scores/{collection_id} [get]
func (h *Handler) GetScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.FindScoresByCollectionID(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetFinalScores
// @Summary Get Final Scores
// @Description Get Final Scores
// @Tags Criteria
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/final_scores/{collection_id} [get]
func (h *Handler) GetFinalScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.FindFinalScoresByCollectionID(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// UpdateCriteriaAlternative
// @Summary Update Criteria Alternative
// @Description Update Criteria Alternative
// @Tags Criteria
// @Accept json
// @Produce json
// @Param request body dto.CriteriaUpdateRequest true "request body"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/criteria [patch]
func (h *Handler) UpdateCriteriaAlternative(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.CriteriaUpdateRequest)
	if err := c.Bind(&payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.UpdateCriteria(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// CalculateAlternativeToPoint
// @Summary Calculate Alternative to Point
// @Description Calculate Alternative to Point
// @Tags Criteria
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/point/calculate/{collection_id} [get]
func (h *Handler) CalculateAlternativeToPoint(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.CalculateAlternativeToPoint(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)

}

// CalculateScores
// @Summary Calculate Scores
// @Description Calculate Scores
// @Tags Criteria
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/scores/calculate/{collection_id} [get]
func (h *Handler) CalculateScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.CalculateScore(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// CalculateFinalScores
// @Summary Calculate Final Scores
// @Description Calculate Final Scores
// @Tags Criteria
// @Accept json
// @Produce json
// @Param collection_id path string true "collection_id path"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/final_scores/calculate/{collection_id} [get]
func (h *Handler) CalculateFinalScores(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.AHPByCollectionIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.CalculateFinalScore(ctx, &payload.CollectionID)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}
