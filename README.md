# gogo-blueprint 🔥
simple project - implement api service by golang

#### Blueprint
- [x]  Support Jaeger
- [x]  MongoDB
- [x]  Unit Test & Integration Test (DB)
- [x]  Logstash to ElasticSearch
- [x]  Implement Application, Service, Repository, domain Layer
- [ ]  Kubenetes Deployment

---
### Folder Structure

```
  /app                        # application layer
    /inout                    # api input / output
      /company
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
      repository.go           # repository interface

  /service                    # service layer for control domains
    /company
      /mocks                  # mock service for testing
      create.go
 
  main.go                     # initial application
  setup.go                    # load and setup dependency
```


---

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
tracing operation with jaeger

```go
	if span := jaegerstart.StartNewSpan(CONTEXT, OPERATION_NAME); span != nil {
		defer span.Finish()
	}
```


### Others

- uber golang style guide [link](https://github.com/uber-go/guide)

