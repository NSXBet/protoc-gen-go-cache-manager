APP_NAME=go-cache-manager

.PHONY: gen
gen:
	@go build -o bin/protoc-gen-$(APP_NAME) main.go
	@cd ./protos && PATH="${PATH}:${PWD}/bin" buf generate
