package main

import (
	"context"

	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"

	"master-data/modules/mysql"

	teknisiController "master-data/api/v1/teknisi"
	teknisiService "master-data/business/teknisi"
	teknisiModule "master-data/modules/teknisi"

	kerusakanController "master-data/api/v1/kerusakan"
	kerusakanService "master-data/business/kerusakan"
	kerusakanModule "master-data/modules/kerusakan"

	config "master-data/config"

	masterDataApp "master-data/api/v1"

	"github.com/labstack/gommon/log"
)

func newDatabaseConnection() *sqlx.DB {
	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.GetConfig().MySQL.User,
		config.GetConfig().MySQL.Password,
		config.GetConfig().MySQL.Host,
		config.GetConfig().MySQL.Port,
		config.GetConfig().MySQL.Name)

	db := mysql.NewDatabaseConnection(uri)

	return db
}

func newTeknisiService(dbsql *sqlx.DB) teknisiService.Service {
	teknisiRepoSQL := teknisiModule.NewMySQLRepository(dbsql)

	return teknisiService.NewService(teknisiRepoSQL)
}

func newKerusakanService(dbsql *sqlx.DB) kerusakanService.Service {
	kerusakanRepoSQL := kerusakanModule.NewMySQLRepository(dbsql)

	return kerusakanService.NewService(kerusakanRepoSQL)
}

func main() {

	// database init
	dbsql := newDatabaseConnection()

	// service init
	teknisiService := newTeknisiService(dbsql)
	kerusakanService := newKerusakanService(dbsql)

	// controller init
	teknisiControllerV1 := teknisiController.NewController(teknisiService)
	kerusakanControllerV1 := kerusakanController.NewController(kerusakanService)

	e := echo.New()

	//register API path and handler
	masterDataApp.RegisterPath(
		e,
		*teknisiControllerV1,
		*kerusakanControllerV1)

	// run server
	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfig().BackendPort)

		if err := e.Start(address); err != nil {
			log.Error("failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the server with
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("failed to grafefully shutting down server", err)
	}

}
