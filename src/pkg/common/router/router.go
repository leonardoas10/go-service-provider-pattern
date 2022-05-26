package router

import (
	"fmt"
	apis "github/leonardoas10/go-provider-pattern/src/pkg/common/apis"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	fmt.Println("llego")
	jsonplaceholdersGroup := e.Group("/json-placeholders")

	//set all middlewares
	// middlewares.SetMainMiddleWares(e)
	// middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	apis.JsonplaceholdersGroup(jsonplaceholdersGroup)

	// //set groupRoutes
	// api.AdminGroup(adminGroup)
	return e
}

