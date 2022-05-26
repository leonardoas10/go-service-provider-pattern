package main

import (
	"github/leonardoas10/go-provider-pattern/src/pkg/common/env"
	router "github/leonardoas10/go-provider-pattern/src/pkg/common/router"
)

func main() {
	port := env.GetEnvVariable("PORT")

	e := router.New()
	e.Logger.Fatal(e.Start(port))
}
