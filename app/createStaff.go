package app

import (
	"net/http"

	"github.com/opentracing/opentracing-go"

	staff2 "github.com/devit-tel/gogo-blueprint/app/inout/staff"
	serviceStaff "github.com/devit-tel/gogo-blueprint/service/staff"

	"github.com/devit-tel/goerror/ginresp"
	"github.com/gin-gonic/gin"
)

func (app *App) CreateStaff(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.createStaff",
	)
	defer span.Finish()

	input := &staff2.CreateStaffInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staff, err := app.staffService.CreateStaff(ctx,
		&serviceStaff.CreateStaffInput{
			Name:      input.Name,
			CompanyId: input.CompanyId,
			Tel:       input.Tel,
		})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff2.CreateStaffOutput{
		Staff: staff2.ToStaffOutput(staff),
	})
}
