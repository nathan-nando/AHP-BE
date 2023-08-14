package subCriteria

import (
	"ahp-be/internal/model"
	"ahp-be/internal/subCriteria"
	"ahp-be/internal/subCriteria/dto"
	"ahp-be/internal/subCriteria/usecase"
	"ahp-be/pkg/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler(repo subCriteria.Repository, db *gorm.DB) *Handler {
	service := usecase.New(repo, db)
	return &Handler{service}
}

// GetSubCriteria
// @Summary Get All Sub Criteria
// @Description Get All Criteria
// @Tags SubCriteria
// @Accept json
// @Produce json
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/sub-criteria [get]
func (h *Handler) GetSubCriteria(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.FindSubCriteria(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// UpdateSubCriteria
// @Summary Update Sub Criteria by Mode
// @Description Update Sub Criteria by Mode
// @Tags SubCriteria
// @Accept json
// @Produce json
// @Param request body model.SubCriteria true "request body"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/sub-criteria [patch]
func (h *Handler) UpdateSubCriteria(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.SubCriteria)
	if err := c.Bind(&payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.UpdateSubCriteria(ctx, payload)

	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// CheckConsistencyRatio
// @Summary Check Consistency Ratio
// @Description Check Consistency Ratio
// @Tags SubCriteria
// @Accept json
// @Produce json
// @Param mode path string true "mode"
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /ahp/sub-criteria/check/{mode} [get]
func (h *Handler) CheckConsistencyRatio(c echo.Context) error {
	ctx := c.Request().Context()

	mode := new(dto.CheckCRSubCriteria)

	if err := c.Bind(mode); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(mode); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.CheckConsistencyRatio(ctx, mode)

	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}
