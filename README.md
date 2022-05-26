# Go Provider Pattern

I wrote a simple project to practice `Provider Pattern in Golang`, Go is suitable for this, thanks to the implicit interface implementation and packages structure and naming, it could give us the chance to separate in a proper way the logic of external apis as the case of `placeholders`.

Using a structure of files like `provider` `service`, where in the last one, i create a interface references a couple of func that needs to match if we want an instance of `service`

### What is the benefit of separate provider and service?

Service package never know how Provider package gonna find the data, it only matters if Provider package give the struct that service package are waiting for.

So, if the external provider change, we only have to change the code of Provider package... adjusting the response to match Service interface.

Always decoupling...
