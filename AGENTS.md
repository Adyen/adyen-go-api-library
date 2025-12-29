## Overview

This is the Adyen Go API Library, providing Go developers with an easy way to interact with the Adyen API. The library is a wrapper around the Adyen API, with services and models generated from OpenAPI specifications.

## Code Generation

A significant portion of this library, particularly the API services and models, is automatically generated.

- **Engine**: We use [OpenAPI Generator](https://openapi-generator.tech/) with custom [Mustache](https://mustache.github.io/) templates to convert Adyen's OpenAPI specifications into Go code.
- **Templates**: The custom templates are located in the `templates/custom` directory. These templates are tailored to fit our custom HTTP client, service structure, and model definitions.
- **Automation**:
    - **Centralized**: The primary generation process is managed in a separate repository, [`adyen-sdk-automation`](https://github.com/Adyen/adyen-sdk-automation). Changes to the OpenAPI specs trigger a GitHub workflow in that repository, which generates the code and opens Pull Requests in this library.
    - **Local**: For development and testing, you must use the [`adyen-sdk-automation`](https://github.com/Adyen/adyen-sdk-automation) repository.

### Local Code Generation

To test new features or changes to the templates, you must run the generation process from a local clone of the `adyen-sdk-automation` repository.

1.  **Clone the automation repository**:
    ```bash
    git clone https://github.com/Adyen/adyen-sdk-automation.git
    ```

2.  **Link this library**: The automation project needs to target your local clone of `adyen-go-api-library`. From inside the `adyen-sdk-automation` directory, run the following commands. This will replace the `go/repo` directory with a symlink to your local project.
    ```bash
    rm -rf go/repo
    ln -s /path/to/your/adyen-go-api-library go/repo
    ```

3.  **Run the generator**: You can now run the Gradle commands to generate code.
    - **To generate all services for the Go library**:
      ```bash
      ./gradlew :go:services
      ```
    - **To generate a single service (e.g., Checkout)**:
      ```bash
      ./gradlew :go:checkout
      ```
    - **To clean the repository before generating**:
      ```bash
      ./gradlew :go:cleanRepo :go:checkout
      ```

## Core Components

- **`adyen.NewClient`**: The main entry point in `src/adyen/api.go`. It creates a new client instance that provides access to all API services.
- **`common.Config`**: The central struct for configuring the library, including API key, environment, basic authentication credentials, and other settings.
- **`common.HTTPClient`**: The underlying client responsible for making HTTP requests to the Adyen API. A custom `http.Client` can be injected for advanced configurations like proxies.
- **`src/{service}`**: Each directory in `src` (e.g., `src/checkout`, `src/management`) represents an API service. These packages contain:
    - The generated API service files (e.g., `api_payments.go`).
    - The generated request and response models (e.g., `model_payment_request.go`).

## Development Workflow

This project uses standard Go tooling and a `Makefile` to streamline development tasks.

### Setup

To install development dependencies and prepare the environment, ensure you have Go installed and run:
```bash
go mod tidy
```

### Running Tests

You can run the unit tests using the `Makefile` target, which is the standard practice for this project and is used in the CI pipeline.
```bash
make test
```
Alternatively, you can run the tests directly using the Go toolchain:
```bash
go test ./...
```

## Release Process

The release process is automated via GitHub Actions. When a release is triggered:
1.  A script determines the next version number (major, minor, or patch).
2.  The `VERSION` file and other configuration files are updated.
3.  A pull request is created with the version bump.
4.  Once merged, a GitHub release is created with the corresponding git tag.
