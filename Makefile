.PHONY:
.SILENT:

openapi:
	oapi-codegen -generate types -o internal/port/http/v1/openapi_type.gen.go -package v1 api/openapi/api.yaml
	oapi-codegen -generate chi-server -o internal/port/http/v1/openapi_server.gen.go -package v1 api/openapi/api.yaml

lint:
	golangci-lint run