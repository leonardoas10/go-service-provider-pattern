# Go Provider Pattern

I wrote a simple project to practice `Provider Pattern in Golang`, Go is suitable for this, thanks to the implicit interface implementation and packages structure and naming, it could give us the chance to separate in a proper way the logic of external apis as the case of `placeholders`.

Using a structure of files like `provider` `service`, where:

-   Provider pkg have the external logic
-   Service pkg have the internal logic, for example, business logic.

Insure of decoupling the app.. but

### What is the benefit of separate provider and service?

Service package never know how Provider package gonna find the data, it only matters if Provider package give the struct that service package are waiting for.

So, if the external provider change, we only have to change the code of Provider package... adjusting the response to match Service interface.

###### For more details

Look inside Service pkg, exist an interface referencing a couple of functions that needs to match if we want an instance of it. The Provider pkg, needs to accomplish the requirements of interface for use the internal logic/business logic.

# How to Run App

### With Go

-   Create .env file in src/cmd/main directory
-   go mod download
-   go run main.go inside src/cmd/main directory

## Or Docker

-   Install [ Docker Engine ](https://docs.docker.com/engine/install/) :fire:

Then, two options:

### 1. Docker

-   Build image `docker build -t go-provider-pattern .`
-   Run container `docker run -v $(pwd):/app -dp 3005:3000 go-provider-pattern`
-   Go to the app [ App ](http://127.0.0.1:3005/json-placeholders)

### 2. Docker Compose

-   Build image `docker-compose build .`
-   Run container `docker-compose up -d`
-   Go to the app [ App ](http://127.0.0.1:3005/json-placeholders)

### Start reading code, interpreting functionalities and programming: smile:
