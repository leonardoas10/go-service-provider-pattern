# Go Provider Pattern

I wrote a simple project to practice `Provider Pattern in Golang`, Go is suitable for this, thanks to the implicit interface implementation and packages structure and naming, it could give us the chance to separate in a proper way the logic of external apis as the case of `placeholders`.

Using a structure of files like `provider` `service`, where:

-   Provider pkg have the external logic
-   Service pkg have the internal logic, for example, business logic.

Insure of decoupling the app.. but

### What is the benefit of separate provider and service?

Service package never know how Provider package gonna find the data, it only matters if Provider package give the struct that service package are waiting for.

So, if the external provider change, we only have to change the code of Provider package... adjusting the response to match Service interface.

For that, exist an interface (Service pkg) referencing a couple of functions that needs to match if we want an instance of Service pkg.

Always decoupling...

# How to Run App

-   Create .env file in src/cmd/main directory
-   go run main.go inside src/cmd/main directory

or

-   Install [ Docker Engine ](https://docs.docker.com/engine/install/) :fire:
-   Build image `docker build -t go-provider-pattern .`
-   Run container `docker run -dp 5000:3000 --name go-provider-pattern go-provider-pattern`
-   Go to the app [ App ](http://127.0.0.1:5000/json-placeholders)
-   Start reading code, interpreting functionalities and programming: smile:
