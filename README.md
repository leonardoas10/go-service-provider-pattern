# Go Service Provider Pattern

This project explores the Provider Pattern in Go, leveraging its built-in support for implicit interfaces and modular package structures. By segregating the logic of external APIs, such as placeholders, into the `provider package` and internal application logic into the `service package`, the application achieves a cleaner separation of concerns.

#### Benefits of Separation

Separating the provider and service functionalities offers the following advantages:

Decoupling: The service package remains agnostic to how data is retrieved, focusing solely on the expected data structure provided by the provider.

Flexibility: Adapting to changes in external providers becomes straightforward, as adjustments are confined to the provider package, ensuring minimal impact on the service.

#### Implementation Details

Provider Package: Contains the logic for interfacing with external APIs or data sources, adhering to the predefined interface expected by the service.

Service Package: Houses the core application logic, including business rules, and relies on the defined interface to interact with the provider seamlessly.

# How to Run App

#### With Go

-   Create .env file in src/cmd/main directory
-   go mod download
-   go run main.go inside src/cmd/main directory

#### Or Docker

-   Install [ Docker Engine ](https://docs.docker.com/engine/install/) :fire:
-   In `root` execute the following...
-   Build image `docker-compose build`
-   Run container `docker-compose up -d`
-   Go to the app [ App ](http://127.0.0.1:3005/json-placeholders)

# How to test it? - Using Testing pkg 🧪

-   Set APP_ENV variable, in the root execute `export APP_ENV=.env.test`
-   Execute `go test ./...`

#### use test coverage (indicates the proportion of your code that is covered by tests)

-   Execute `go test -coverprofile=coverage.out ./...`
-   Execute `go tool cover -html=coverage.out -o coverage.html`
-   Load the coverage.html in the browser

### Start reading code, interpreting functionalities and programming :smile:
