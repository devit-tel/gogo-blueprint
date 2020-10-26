# gogo-blueprint 🔥

<p align="left">
  <a href="https://github.com/devit-tel/gogo-blueprint"><img alt="GitHub Actions status" src="https://github.com/devit-tel/gogo-blueprint/workflows/go-unit-test/badge.svg"></a>
</p>

simple project - implement api service by golang

#### Blueprint
- [x]  Support Jaeger
- [x]  MongoDB
- [x]  Unit Test & Integration Test (DB)
- [x]  Logstash to ElasticSearch
- [x]  Implement Application, Service, Repository, domain Layer
- [x]  run debug mode
- [ ]  Kubenetes Deployment
- [ ]  Gitlab CI

---
### Folder Structure

```
  /app                        # application layer
    /inout                    # api input / output
      /company
      	company.go           
	craeteCompany.go      # create company input / output
    app.go                    # setup api handlers
    createCompany.go          # api
    createCompany_test.go     # api testing

  /config                     # load config (from env)
    config.go

  /deployment                 # kubernetes config for deploy

  /development                # development tools (db, jaeger, local env)
    docker-compose.yml        # docker mongodb, elasticsearch + kibana, jaeger
    local.env                 # local env for integration test

  /domain                     # business logic layer
    /company
      company.go
      company_test.go

  /external                   # external service layer
  
  /lib                        # internal library
  
  /repository                 # repository layer
    /company
      /mocks                  # repository mocks for testing
      /store                  # repository implement interface
      repository.go           # repository interface & repository errors

  /service                    # service layer for control domains
    /company
      /withtracer	      # implement jaeger trace by wrap service interface
      /mocks                  # mock service for testing
      create.go
 
  main.go                     # initial application
  setup.go                    # load and setup dependency
```


---

### Pre-Require

Mockery
```
GO111MODULE=off go get github.com/vektra/mockery/.../
```

### Installation

```
  git clone https://github.com/devit-tel/gogo-blueprint.git
  cd gogo-blueprint
  go mod download
```



### Testing 
unit testing command

```
  go test ./...
```

integrating testing command

```
  go test ./... -tags integration
```


### Generate Mocks

generate mocks from interfaces for unit testing

```
  go generate ./...
```


### Local development
development in local start mongodb, elasticsearch + kibana, jaeger

```
cd development
source ./local.env
docker-compose up -d
```

### Tracing with Jaeger
please see in the example code implement jaeger wrap service ```service/company/withtracer```


### Others

- Uber golang style guide [link](https://github.com/uber-go/guide)
- Practical Go: Real world advice for writing maintainable Go programs [link](https://dave.cheney.net/practical-go/presentations/qcon-china.html?fbclid=IwAR2_D2Y2HXVYUNiG3LctB0kF64YKzGUatcIHm_sLYwm9SEqEKWAd76G7NAU)
- How to run debug go code in vscode [link](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)