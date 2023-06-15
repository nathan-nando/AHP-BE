package collection

import (
	"ahp-be/internal/collection"
	"ahp-be/internal/collection/dto"
	"ahp-be/internal/collection/usecase"
	"ahp-be/pkg/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler(repo collection.Repository, db *gorm.DB) *Handler {
	service := usecase.New(repo, db)
	return &Handler{service}
}

// Get
// @Summary Get All Collections
// @Description Get All Collections
// @Tags collection
// @Accept json
// @Produce json
// @Success 200 {object} dto.FindCollectionsResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection [get]
func (h *Handler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.service.Finds(ctx)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// GetByID
// @Summary Get Collection By CollectionID
// @Description Get Collection By CollectionID
// @Tags collection
// @Accept json
// @Produce json
// @Param id path string true "id path"
// @Success 200 {object} dto.FindCollectionByIDResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection/{id} [get]
func (h *Handler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.FindCollectionByIDRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		res := response.ErrorBuilder(&response.ErrorConstant.Validation, err)
		return res.Send(c)
	}

	fmt.Printf("%+v", payload)

	result, err := h.service.FindByID(ctx, payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(result).Send(c)
}

// Create godoc
// @Summary Create collection
// @Description Create collection
// @Tags collection
// @Accept  json
// @Produce  json
// @Param request body dto.CreateCollectionRequest true "request body"
// @Success 200 {object} dto.CreateCollectionResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection [post]
func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(dto.CreateCollectionRequest)

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
// @Summary Update collection
// @Description Update collection
// @Tags collection
// @Accept  json
// @Produce  json
// @Param request body dto.UpdateCollectionRequest true "request body"
// @Success 200 {object} dto.UpdateCollectionResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection [patch]
func (h *Handler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.UpdateCollectionRequest)
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
// @Summary Delete collection
// @Description Delete collection
// @Tags collection
// @Accept  json
// @Produce  json
// @Param id path string true "id path"
// @Success 200 {object}  dto.DeleteCollectionResponseDoc
// @Failure 400 {object} response.errorResponse
// @Failure 404 {object} response.errorResponse
// @Failure 500 {object} response.errorResponse
// @Router /collection/{id} [delete]
func (h *Handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(dto.DeleteCollectionRequest)
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
