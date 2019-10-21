############################
# STEP 1 setup for build project
############################
FROM golang:1.13.2-alpine AS builder
# Install git
# Git is required for fetching the dependencies
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/devit-tel/gogo-blueprint/
COPY . .
# Enable go module
ENV GO111MODULE=on
# Fetch dependencies
# Using go get
RUN go mod download
# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/gogo-blueprint

############################
# STEP 2 build a small image
############################
FROM scratch
# Expose port
EXPOSE 8080
# Copy our static executable.
COPY --from=builder /go/bin/gogo-blueprint /go/bin/gogo-blueprint
# Run the hello binary
ENTRYPOINT ["/go/bin/gogo-blueprint"]
