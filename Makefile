APP_NAME=go-cache-manager

.PHONY: gen
gen: gen-setup
	@go build -o bin/protoc-gen-$(APP_NAME) main.go
	@cd ./protos && PATH="${PWD}/bin:${PATH}" buf generate

.PHONY: test
test: gen
	@go test -v ./...

.PHONY: gen-setup
gen-setup:
	@if [ -x "$$(command -v buf)" ]; then \
		: ; \
	else \
	    echo "buf could not be found! Installing..."; \
	    go install github.com/bufbuild/buf/cmd/buf@latest; \
	fi
