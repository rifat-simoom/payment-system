include .env
export

.PHONY: openapi
openapi: openapi_http

.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh card internal/card/ports ports


