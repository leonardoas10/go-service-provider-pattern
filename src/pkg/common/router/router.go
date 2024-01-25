package router

import (
	apis "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/apis"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	jsonplaceholdersGroup := e.Group("/json-placeholders")

	//set main routes
	apis.JsonplaceholdersGroup(jsonplaceholdersGroup)
	
	return e
}

