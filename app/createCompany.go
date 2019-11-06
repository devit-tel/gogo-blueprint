package app

import (
	"net/http"

	"github.com/devit-tel/goerror/ginresp"
	company2 "github.com/devit-tel/gogo-blueprint/app/inout/company"
	serviceCompany "github.com/devit-tel/gogo-blueprint/service/company"
	"github.com/gin-gonic/gin"
)

func (app *App) CreateCompany(c *gin.Context) {
	input := &company2.CreateCompanyInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	company, err := app.companyService.CreateCompany(c.Request.Context(), &serviceCompany.CreateCompanyInput{Name: input.Name})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &company2.CreateCompanyOutput{
		Company: company2.ToCompanyOutput(company),
	})
}
