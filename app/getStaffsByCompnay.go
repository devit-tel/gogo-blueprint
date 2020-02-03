package app

import (
	"net/http"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/devit-tel/gogo-blueprint/app/inout/staff"
	serviceStaff "github.com/devit-tel/gogo-blueprint/service/staff"
)

func (app *App) GetStaffsByCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.getStaffsByCompany",
	)
	defer span.Finish()

	input := &staff.GetStaffsByCompanyInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staffs, err := app.staffService.GetStaffsByCompany(ctx,
		&serviceStaff.GetStaffsByCompanyInput{
			CompanyId: input.CompanyId,
			Offset:    input.Offset,
			Limit:     input.Limit,
		})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff.GetStaffsByCompanyOutput{
		Staffs: staff.ToStaffsOutput(staffs),
	})
}
