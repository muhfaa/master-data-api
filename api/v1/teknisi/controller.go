package loan

import (
	"master-data/api/common"
	"master-data/api/v1/teknisi/request"
	"master-data/api/v1/teknisi/response"
	"master-data/business"
	"master-data/business/teknisi"
	"master-data/util/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	teknisiService teknisi.Service
}

func NewController(teknisiService teknisi.Service) *Controller {
	return &Controller{
		teknisiService,
	}
}

// Insert Teknisi
func (controller Controller) InsertTeknisi(c echo.Context) error {
	insertTeknisi := new(request.InsertTeknisiRequest)

	if err := c.Bind(insertTeknisi); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(insertTeknisi); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	teknisiSpec := insertTeknisi.ToUpsertTeknisiSpec()
	err := controller.teknisiService.InsertTeknisi(*teknisiSpec)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(business.ErrDuplicate))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

// Get Teknisi
func (controller Controller) GetAllTeknisi(c echo.Context) error {

	allTeknisi, err := controller.teknisiService.FindAllTeknisi()
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	successResponse := common.NewSuccessResponse(response.NewTeknisiDataResponse(allTeknisi))

	return c.JSON(http.StatusOK, successResponse)
}

// Update Teknisi
func (controller Controller) UpdateTeknisi(c echo.Context) error {
	updateTeknisi := new(request.UpdateTeknisiRequest)

	if err := c.Bind(updateTeknisi); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateTeknisi); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	teknisiSpec := updateTeknisi.ToUpdateTeknisiSpec()
	isUpdated, err := controller.teknisiService.UpdateTeknisi(*teknisiSpec)
	if err != nil && !isUpdated {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

// Update Antrian
func (controller Controller) AddAntrian(c echo.Context) error {
	updateAntrian := new(request.UpdateJumlahAntrianRequest)

	if err := c.Bind(updateAntrian); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateAntrian); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	antrianSpec := updateAntrian.ToUpdateJumlahAntrianSpec()
	isUpdated, err := controller.teknisiService.AddAntrian(antrianSpec.ID, antrianSpec.Version)
	if err != nil && !isUpdated {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponse(response.NewTeknisiUpdateResponse(antrianSpec.ID)))
}

func (controller Controller) EraseAntrian(c echo.Context) error {
	updateAntrian := new(request.UpdateJumlahAntrianRequest)

	if err := c.Bind(updateAntrian); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateAntrian); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	antrianSpec := updateAntrian.ToUpdateJumlahAntrianSpec()
	isUpdated, err := controller.teknisiService.EraseAntrian(antrianSpec.ID, antrianSpec.Version)
	if err != nil && !isUpdated {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponse(response.NewTeknisiUpdateResponse(antrianSpec.ID)))
}

// Delete Teknisi
func (controller Controller) DeleteTeknisi(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	isSuccess, err := controller.teknisiService.DeleteTeknisi(idInt)
	if err != nil || !isSuccess {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}
