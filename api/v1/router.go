package router

import (
	kerusakanController "master-data/api/v1/kerusakan"
	teknisiController "master-data/api/v1/teknisi"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	teknisiController teknisiController.Controller,
	kerusakanController kerusakanController.Controller,
) {

	// teknisi
	teknisiV1 := e.Group("v1/teknisi")
	teknisiV1.POST("/add", teknisiController.InsertTeknisi)
	teknisiV1.GET("", teknisiController.GetAllTeknisi)
	teknisiV1.PUT("/update", teknisiController.UpdateTeknisi)
	teknisiV1.PUT("/add/antrian", teknisiController.AddAntrian)
	teknisiV1.PUT("/erase/antrian", teknisiController.EraseAntrian)
	teknisiV1.DELETE("/id/:id", teknisiController.DeleteTeknisi)

	// kerusakan
	kerusakanV1 := e.Group("v1/kerusakan")
	kerusakanV1.POST("/add", kerusakanController.InsertKerusakan)
	kerusakanV1.GET("", kerusakanController.GetAllKerusakan)
	kerusakanV1.GET("/id/:id", kerusakanController.GetKerusakan)
	kerusakanV1.PUT("/update", kerusakanController.UpdateKerusakan)
	kerusakanV1.DELETE("/id/:id", kerusakanController.DeleteKerusakan)

}
