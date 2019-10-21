package main

import (
	"io"
	"log"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/devit-tel/gogo-blueprint/app"
	"github.com/devit-tel/gogo-blueprint/config"
	companyRepo "github.com/devit-tel/gogo-blueprint/repository/company/store"
	staffRepo "github.com/devit-tel/gogo-blueprint/repository/staff/store"
	companyService "github.com/devit-tel/gogo-blueprint/service/company"
	staffService "github.com/devit-tel/gogo-blueprint/service/staff"
	"github.com/devit-tel/goxid"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"
)

func setupJaeger(appConfig *config.Config) (opentracing.Tracer, io.Closer) {
	transport, err := zipkin.NewHTTPTransport(
		appConfig.JaegerEndpoint,
		zipkin.HTTPBatchSize(50),
		zipkin.HTTPLogger(jaeger.StdLogger),
	)
	if err != nil {
		log.Fatalf("Cannot initialize HTTP transport: %v", err)
	}

	return jaeger.NewTracer(
		"GoGoBoilerplate",
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(transport),
	)
}

func newApp(appConfig *config.Config) *app.App {
	xid := goxid.New()

	companyStore := companyRepo.New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	company := companyService.New(xid, companyStore)

	staffStore := staffRepo.New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	staff := staffService.New(xid, staffStore, companyStore)

	return app.New(staff, company)
}

func setupLog() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}

func setupHookToLogstash(logger *logrus.Logger, appConfig *config.Config) {
	hook, err := logrustash.NewHook("udp", appConfig.LogstashEndpoint, appConfig.AppName)
	if err != nil {
		log.Fatal(err)
	}

	logger.Hooks.Add(hook)
}
