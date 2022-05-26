package main

import (
	"github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/common/env"
	"github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/common/http"
)

func main() {
	port := env.GetEnvVariable("PORT")

	http.StartHttp(port)
}
