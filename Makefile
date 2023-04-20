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

openapi-generator-version:=6.5.0
openapi-generator-url:=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(openapi-generator-version)/openapi-generator-cli-$(openapi-generator-version).jar
openapi-generator-jar:=bin/openapi-generator-cli.jar
openapi-generator-cli:=java -jar $(openapi-generator-jar)
goimports:=$(GOPATH)/bin/goimports

generator:=go
services:=checkout legalentity payments
output:=src/
templates:=templates/small

# Generate models (for each service)
models: $(services)

checkout: spec=CheckoutService-v70
checkout: service=checkout
legalentity: spec=LegalEntityService-v3
legalentity: service=legalentity
payments: spec=PaymentService-v68
payments: service=payments

# Generate a full client (models and service classes)
$(services): schema $(openapi-generator-jar) $(goimports)
	GO_POST_PROCESS_FILE="$(goimports) -w" $(openapi-generator-cli) generate \
		-i schema/json/$(spec).json \
		-g $(generator) \
		-t $(templates) \
		-o $(output)$(service) \
		-p packageName=$(@) \
		--global-property apiTests=false \
		--global-property apis,models \
		--global-property apiDocs=false \
		--global-property modelDocs=true \
		--git-repo-id adyen-go-api-library/v6 --git-user-id adyen \
		--enable-post-process-file \
		--inline-schema-name-mappings PaymentDonationRequest_paymentMethod=CheckoutPaymentMethod \
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

# Download the import optimizer (and code formatter)
$(goimports):
	go install golang.org/x/tools/cmd/goimports@latest

# Discard generated artifacts and changed models
clean:
	git checkout src/checkout
	git clean -f -d src/checkout


.PHONY: templates models $(services)