package alternative

import (
	"ahp-be/internal/alternative"
	"ahp-be/internal/alternative/dto"
	"ahp-be/internal/alternative/usecase"
	"ahp-be/pkg/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler(repo alternative.Repository, db *gorm.DB) *Handler {
	service := usecase.New(repo, db)
	return &Handler{service}
}

// GetsByCollectionID
// @Summary Get All Alternatives By CollectionID
// @Description Get All Alternatives By CollectionID
// @Tags alternative
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} dto.FindAlternativesByIDCollectionResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative/collection/{id} [get]
func (h *Handler) GetsByCollectionID(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.FindAlternativeByIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.FindByCollectionID(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetByID
// @Summary Get Collection By AlternativeID
// @Description Get Collection By AlternativeID
// @Tags alternative
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} dto.FindAlternativeByIDResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative/{id} [get]
func (h *Handler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.FindAlternativeByIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	result, err := h.service.FindByID(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Create godoc
// @Summary Create alternative
// @Description Create alternative
// @Tags alternative
// @Accept  json
// @Produce  json
// @Param request body dto.CreateAlternativeRequest true "request body"
// @Success 200 {object} dto.CreateAlternativeResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative [post]
func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.CreateAlternativeRequest)

	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Create(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Update godoc
// @Summary Update alternative
// @Description Update alternative
// @Tags alternative
// @Accept  json
// @Produce  json
// @Param request body dto.UpdateAlternativeRequest true "request body"
// @Success 200 {object} dto.UpdateAlternativeResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative [patch]
func (h *Handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.UpdateAlternativeRequest)
	if err := c.Bind(&payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Update(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Delete godoc
// @Summary Delete alternative
// @Description Delete alternative
// @Tags alternative
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Success 200 {object}  dto.DeleteAlternativeResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /alternative/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.DeleteAlternativeRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Delete(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}
