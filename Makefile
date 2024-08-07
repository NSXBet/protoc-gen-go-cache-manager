APP_NAME=go-cache-manager

.PHONY: gen
gen:
	@go build -o bin/protoc-gen-$(APP_NAME) main.go
	@cd ./protos && PATH="${PWD}/bin:${PATH}" buf generate

.PHONY: test
test: gen
	@go test -v ./...
