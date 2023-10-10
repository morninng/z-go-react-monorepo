gen-server-model: ## genarete model codes
	@oapi-codegen -generate types -package dto -o server/internal/dto/dto_gen.go api/openapi.yml

gen-client-api:
	@docker run -it --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.2.1 generate \
		-g typescript-axios \
		-i /local/api/openapi.yml \
		-o /local/client/src/api \
		-p modelPropertyNaming=camelCase \
		--type-mappings=set=Array
