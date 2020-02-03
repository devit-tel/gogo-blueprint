package withtracer

import (
	"context"

	"github.com/devit-tel/goerror"
	"github.com/opentracing/opentracing-go"

	domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
	service "github.com/devit-tel/gogo-blueprint/service/company"
)

type ServiceWithTracer struct {
	service service.Service
}

func Wrap(service service.Service) service.Service {
	return &ServiceWithTracer{
		service: service,
	}
}

func (swt *ServiceWithTracer) CreateCompany(ctx context.Context, input *service.CreateCompanyInput) (*domainCompany.Company, goerror.Error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.company.CreateCompany")
	defer sp.Finish()
	sp.LogKV("name", input.Name)

	c, err := swt.service.CreateCompany(ctx, input)
	if err != nil {
		sp.LogKV("error", err)
		return c, err
	}

	sp.LogKV("company", c)
	return c, err
}
