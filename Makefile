BUF?=buf
PLANTUML?=plantuml

all: gen/python/sekura_grpc.py sekura-gateway

gen:
	mkdir -p gen

doc:
	mkdir -p doc

gen/%: api/v1/sekura.proto gen
	poetry run $(BUF) generate

doc/index.html: gen/openapiv2/sekura/v1/sekura.swagger.yaml doc
	docker run --rm -w /local -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/$< -g html -o /local/doc

doc/schemas.plantuml: gen/openapiv2/sekura/v1/sekura.swagger.yaml doc
	docker run --rm -w /local -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/$< -g plantuml -o /local/doc

doc/schemas.png: doc/schemas.plantuml
	$(PLANTUML) $<

.PHONY: documentation
documentation: doc/index.html doc/schemas.plantuml doc/schemas.png

.PHONY: sekura-gateway
sekura-gateway:
	docker build --target gateway -t sekura-gateway gateway

.PHONY: watch-doc
watch-doc:
	iwatch -c "PLANTUML=~/.local/bin/plantuml BUF=~/.local/bin/buf make documentation" -e close_write -t "sekura.proto" api/v1

watch-lint:
	iwatch -c "~/.local/bin/buf lint" -e close_write -t "sekura.proto" api/v1
