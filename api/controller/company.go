package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/go-pg/pg"

	"github.com/jeroldleslie/golang-web-scraper/api/model"
)

type CompanyController struct{
	DB	*pg.DB
}

// Get Single Company Info
func (c *CompanyController) GetCompany(ctx echo.Context) error {
	cvrid := ctx.Param("cvrid")
	companyInfo := &model.CompanyInfo{CvrId: cvrid}
	err := c.DB.Select(companyInfo)
	if(err != nil){
		return ctx.String(http.StatusNotFound, "record not found")
	}
	return ctx.JSON(http.StatusCreated, companyInfo)
}

//Get All Companies Info
func (c *CompanyController) GetCompaniesInfo(ctx echo.Context) error {
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")

	if limit == "" || offset == "" {
		return ctx.String(http.StatusUnprocessableEntity,"limit,offset query parameter is mandatory")
	}
	var companies []model.CompanyInfo
	_, err := c.DB.Query(&companies, "select * from company_info limit ? offset ?", limit, offset)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "something went wrong")
	}

	return ctx.JSON(http.StatusOK, companies)
}
