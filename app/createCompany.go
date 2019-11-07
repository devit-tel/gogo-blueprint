package app

import (
	"net/http"

	"github.com/devit-tel/goerror/ginresp"
	company "github.com/devit-tel/gogo-blueprint/app/inout/company"
	serviceCompany "github.com/devit-tel/gogo-blueprint/service/company"
	"github.com/gin-gonic/gin"
)

func (app *App) CreateCompany(c *gin.Context) {
	input := &company.CreateCompanyInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	newCompany, err := app.companyService.CreateCompany(c.Request.Context(), &serviceCompany.CreateCompanyInput{Name: input.Name})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &company.CreateCompanyOutput{
		Company: company.ToCompanyOutput(newCompany),
	})
}
