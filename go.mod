module github.com/adyen/adyen-go-api-library/v21

go 1.25.0

retract v21.2.1 // Source-breaking change to webhook.NotificationRequestItem.AdditionalData. Use v21.2.2 or later.

// Maintainers: keep the retract directive above in every subsequent release.
// The go command reads retractions only from the go.mod of the highest
// published version, so dropping it would make v21.2.1 selectable again.

require (
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/stretchr/testify v1.11.1
	golang.org/x/oauth2 v0.36.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
