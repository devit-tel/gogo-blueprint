package withtracer

import (
	"context"

	"github.com/devit-tel/goerror"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
	service "github.com/devit-tel/gogo-blueprint/service/staff"
)

type ServiceWithTracer struct {
	service service.Service
}

func Wrap(service service.Service) service.Service {
	return &ServiceWithTracer{
		service: service,
	}
}

func (swt *ServiceWithTracer) GetStaffsByCompany(ctx context.Context, input *service.GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.GetStaffsByCompany")
	defer sp.Finish()
	sp.LogFields(
		log.String("companyId", input.CompanyId),
		log.Int64("offset", input.Limit),
		log.Int64("companyId", input.Offset),
	)

	staffs, err := swt.service.GetStaffsByCompany(ctx, input)
	if err != nil {
		sp.LogKV("error", err)
		return staffs, err
	}

	sp.LogKV("staffs", staffs)
	return staffs, err
}

func (swt *ServiceWithTracer) CreateStaff(ctx context.Context, input *service.CreateStaffInput) (*domainStaff.Staff, goerror.Error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.CreateStaff")
	defer sp.Finish()
	sp.LogKV("input", input)

	staff, err := swt.service.CreateStaff(ctx, input)
	if err != nil {
		sp.LogKV("error", err)
		return staff, err
	}

	sp.LogKV("staff", staff)
	return staff, err
}

func (swt *ServiceWithTracer) UpdateStaff(ctx context.Context, input *service.UpdateStaffInput) (*domainStaff.Staff, goerror.Error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.UpdateStaff")
	defer sp.Finish()
	sp.LogKV("input", input)

	staff, err := swt.service.UpdateStaff(ctx, input)
	if err != nil {
		sp.LogKV("error", err)
		return staff, err
	}

	sp.LogKV("staff", staff)
	return staff, err
}
