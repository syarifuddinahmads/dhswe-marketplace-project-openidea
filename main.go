package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	database "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/db"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/factory"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/http"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/middleware"
)

// load env configuration
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	database.CreateConnectionDB()

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":8000"))

}
