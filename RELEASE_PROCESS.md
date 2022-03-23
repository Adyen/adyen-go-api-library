## Steps for Release of this library

Release is done on GitHub on the `main` branch all development is done on the `develop` branch only

**Step 1**: Update the exact library version to be released in `LibVersion` constant on `src/common/configuration.go`

**Step 2**: If a open-api spec update was done, make sure the to follow [GENERATING_MODELS.md](/GENERATING_MODELS.md) and make sure appropriate versions and URLS are updated on `src/adyen/api.go`

**Step 3**: Run `make build` to ensure there are no build errors

**Step 4**: For patch and minor version releases go to step 6, For major version releases go to step 5

**Step 5**: Find and replace `github.com/adyen/adyen-go-api-library/v<current major version>` with `github.com/adyen/adyen-go-api-library/v<new major version>` through out the project including test files and readme. For example `github.com/adyen/adyen-go-api-library/v5` will become `github.com/adyen/adyen-go-api-library/v6`

**Step 6**: create a PR form `develop` to `main` and wait for all tests to pass and for approvals

**Step 7**: once ready merge the PR to main and create a GitHub release with changelog and a new tag corresponding to the version
