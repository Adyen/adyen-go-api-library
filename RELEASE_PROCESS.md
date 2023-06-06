## Steps for Release of this library

Releases are done on GitHub on the `main` branch. All development is done on the `develop` branch only.

**Step 0**: Update the exact API version to be supported (ie `CheckoutAPIVersion = v69`) in `src/adyen/api.go`

**Step 1**: Update the exact library version to be released in `LibVersion` constant in `src/common/configuration.go`

**Step 2**: If an open-api spec update was done, make sure the to follow [GENERATING_MODELS.md](/GENERATING_MODELS.md) and make sure appropriate versions and URLs are updated on `src/adyen/api.go`

**Step 3**: Run `make build` to ensure that there are no build errors

**Step 4**: For patch and minor version releases go to step 6, For major version releases go to step 5

**Step 5**: Find and replace `github.com/adyen/adyen-go-api-library/v<current major version>` with `github.com/adyen/adyen-go-api-library/v<new major version>` throughout the project including test files and README. For example `github.com/adyen/adyen-go-api-library/v3` will become `github.com/adyen/adyen-go-api-library/v4`

**Step 6**: create a PR form `develop` to `main` and wait for all tests to pass and for approvals

**Step 7**: once ready merge the PR to main and create a GitHub release with changelog and a new tag corresponding to the version
