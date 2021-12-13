package loan

import (
	"master-data/api/common"
	"master-data/api/v1/kerusakan/request"
	"master-data/api/v1/kerusakan/response"
	"master-data/business"
	"master-data/business/kerusakan"
	"master-data/util/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	kerusakanService kerusakan.Service
}

func NewController(kerusakanService kerusakan.Service) *Controller {
	return &Controller{
		kerusakanService,
	}
}

// Insert Kerusakan
func (controller Controller) InsertKerusakan(c echo.Context) error {
	InsertKerusakan := new(request.InsertKerusakanRequest)

	if err := c.Bind(InsertKerusakan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(InsertKerusakan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	kerusakanSpec := InsertKerusakan.ToUpsertKerusakanSpec()
	err := controller.kerusakanService.InsertKerusakan(kerusakanSpec)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(business.ErrDuplicate))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

// Get Kerusakan
func (controller Controller) GetAllKerusakan(c echo.Context) error {

	allKerusakan, err := controller.kerusakanService.FindAllKerusakan()
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	successResponse := common.NewSuccessResponse(response.NewKerusakanDataResponse(allKerusakan))

	return c.JSON(http.StatusOK, successResponse)
}

func (controller Controller) GetKerusakan(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	kerusakan, err := controller.kerusakanService.FindKerusakanByID(idInt)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	successResponse := common.NewSuccessResponse(response.NewKerusakanResponse(kerusakan))

	return c.JSON(http.StatusOK, successResponse)
}

// Update Kerusakan
func (controller Controller) UpdateKerusakan(c echo.Context) error {
	updateKerusakan := new(request.UpdateKerusakanRequest)

	if err := c.Bind(updateKerusakan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(updateKerusakan); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	kerusakanSpec := updateKerusakan.ToUpdateKerusakanSpec()
	isUpdated, err := controller.kerusakanService.UpdateKerusakan(*kerusakanSpec)
	if err != nil && !isUpdated {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

// Delete Kerusakan
func (controller Controller) DeleteKerusakan(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	isSuccess, err := controller.kerusakanService.DeleteKerusakan(idInt)
	if err != nil || !isSuccess {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}
