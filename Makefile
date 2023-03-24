build:
	@echo "Building Adyen Go API library"
	go build -o bin/main main.go

run:
	@echo "Building and running Adyen Go API library"
	go run main.go

test:
	@echo "Testing Adyen Go API library"
	go test ./...

verify: build run test

# Automation

openapi-generator-version:=6.4.0
openapi-generator-url:=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(openapi-generator-version)/openapi-generator-cli-$(openapi-generator-version).jar
openapi-generator-jar:=bin/openapi-generator-cli.jar
openapi-generator-cli:=java -jar $(openapi-generator-jar)

generator:=go
services:=checkout
output:=checkout
templates:=templates/go-v6.4.0

# Generate models (for each service)
models: $(services)

checkout: spec=CheckoutService-v70
checkout: service=checkout

# Generate a full client (models and service classes)
checkout: schema $(openapi-generator-jar)
	GO_POST_PROCESS_FILE="gofmt -w" $(openapi-generator-cli) generate \
		-i schema/json/$(spec).json \
		-g $(generator) \
		-t $(templates) \
		-o $(output) \
		-p packageName=$(@) \
		--global-property apiTests=false \
		--git-repo-id adyen-go-api-library/v6 --git-user-id adyen \
		--enable-post-process-file \
		--additional-properties=useOneOfDiscriminatorLookup=true \
		--additional-properties=serviceName=$@
	rm -rf $(output)/go.{mod,sum}

# Checkout spec (and patch version)
schema:
	git clone https://github.com/Adyen/adyen-openapi.git schema
	perl -i -pe 's/"openapi" : "3.[0-9].[0-9]"/"openapi" : "3.0.0"/' schema/json/*.json

# Extract templates (copy them for modifications)
templates: $(openapi-generator-jar)
	$(openapi-generator-cli) author template -g $(generator) -o schema/templates

# Download the generator
$(openapi-generator-jar):
	wget --quiet -o /dev/null $(openapi-generator-url) -O $(openapi-generator-jar)

# Discard generated artifacts and changed models
clean:
	rm -rf $(output)
	git checkout src/checkout
	git clean -f -d src/checkout


.PHONY: templates models $(services)
