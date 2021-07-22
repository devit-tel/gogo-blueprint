module github.com/devit-tel/gogo-blueprint

go 1.13

require (
	github.com/caarlos0/env/v6 v6.1.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/devit-tel/goerror v0.0.0-20200305073127-b2f4401d8848
	github.com/devit-tel/gotime v0.0.0-20191011035800-6db99a579f46
	github.com/devit-tel/goxid v0.0.0-20191015090949-84147034cdac
	github.com/gin-gonic/gin v1.5.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/toorop/gin-logrus v0.0.0-20190701131413-6c374ad36b67
	github.com/uber-go/atomic v0.0.0-00010101000000-000000000000 // indirect
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.mongodb.org/mongo-driver v1.5.1
	go.uber.org/atomic v1.5.1 // indirect
)

replace github.com/uber-go/atomic => github.com/uber-go/atomic v1.4.0
