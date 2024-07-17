THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY:  start linter-golangci race

start:	### starting app
	go run cmd/app/main.go

linter-golangci: ### check by golangci linter
	golangci-lint run

race: ### checking race conditions
	go run -race cmd/app/main.go