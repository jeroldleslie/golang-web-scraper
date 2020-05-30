package main
import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jeroldleslie/golang-web-scraper/api/controller"
	"github.com/jeroldleslie/golang-web-scraper/api/db"

)
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// DB Connection
	db := db.GetDB()

	companyController := &controller.CompanyController{
		DB : db,
	}

	// Routes
	e.GET("/company/:cvrid", companyController.GetCompany)
	e.GET("/company/_search", companyController.GetCompaniesInfo)



	// Server
	e.Start(":1323")
}