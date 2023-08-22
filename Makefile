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

## Automation

openapi-generator-version:=6.5.0
openapi-generator-url:=https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/$(openapi-generator-version)/openapi-generator-cli-$(openapi-generator-version).jar
openapi-generator-jar:=bin/openapi-generator-cli.jar
openapi-generator-cli:=java -jar $(openapi-generator-jar)
goimports:=$(shell go env GOPATH)/bin/goimports

generator:=go
services:=balanceplatform configurationwebhook reportwebhook transferwebhook binlookup checkout legalentity management payments payout posterminalmanagement recurring storedvalue transfers
output:=src
templates:=templates/custom

# Generate models (for each service)
models: $(services)

balanceplatform: spec=BalancePlatformService-v2
balanceplatform: serviceName=BalancePlatform
balanceplatform: hasRestServiceError=true
configurationwebhook: spec=BalancePlatformConfigurationNotification-v1
reportwebhook: spec=BalancePlatformReportNotification-v1
transferwebhook: spec=BalancePlatformTransferNotification-v3
binlookup: spec=BinLookupService-v54
checkout: spec=CheckoutService-v70
checkout: serviceName=Checkout
legalentity: spec=LegalEntityService-v3
legalentity: serviceName=LegalEntity
payments: spec=PaymentService-v68
payout: spec=PayoutService-v68
recurring: spec=RecurringService-v68
storedvalue: spec=StoredValueService-v46
storedvalue: serviceName=StoredValue
transfers: spec=TransferService-v3
transfers: serviceName=Transfers
transfers: hasRestServiceError=true
management: spec=ManagementService-v1
management: serviceName=Management
management: hasRestServiceError=true
posterminalmanagement: spec=TfmAPIService-v1
posterminalmanagement: serviceName=PosTerminalManagementApi

# Generate a full client (models and service classes)
$(services): schema $(openapi-generator-jar) $(goimports)
	GO_POST_PROCESS_FILE="$(goimports) -w" $(openapi-generator-cli) generate \
		-i schema/json/$(spec).json \
		-g $(generator) \
		-t $(templates) \
		-o $(output)/$(@) \
		-p packageName=$(@) \
		-c ./templates/config.yaml \
		--global-property apis,models \
		--global-property supportingFiles=client.go \
		--global-property apiTests=false \
		--global-property apiDocs=false \
		--global-property modelDocs=false \
		--skip-validate-spec \
		--enable-post-process-file \
		--inline-schema-name-mappings PaymentDonationRequest_paymentMethod=CheckoutPaymentMethod \
		--additional-properties=serviceName=$(serviceName) \
		--additional-properties=$(if $(hasRestServiceError),hasRestServiceError=true)
	rm -rf $(output)/$(@)/go.{mod,sum}
	rm -rf $(output)/$(@)/.openapi-generator/FILES

# Clone OpenAPI spec (and apply local patches)
schema:
	git clone https://github.com/Adyen/adyen-openapi.git schema
	perl -i -pe 's/"openapi" : "3.[0-9].[0-9]"/"openapi" : "3.0.0"/' schema/json/*.json
	for json in schema/json/*.json; do \
		jq -e 'if has("paths") then .paths[][] |= (.operationId = ."x-methodName") else . end' $$json > "$${json}.tmp"; \
		mv "$${json}.tmp" $$json; \
  	done

# Extract templates (copy them for modifications)
templates: $(openapi-generator-jar)
	$(openapi-generator-cli) author template -g $(generator) -o schema/templates

# Download the generator
$(openapi-generator-jar):
	mkdir -p bin
	wget --quiet -o /dev/null $(openapi-generator-url) -O $(openapi-generator-jar)

# Download the import optimizer (and code formatter)
$(goimports):
	go install golang.org/x/tools/cmd/goimports@latest

# Discard generated artifacts and changed models
clean:
	git checkout src
	git clean -f -d src

.PHONY: templates models $(services)
