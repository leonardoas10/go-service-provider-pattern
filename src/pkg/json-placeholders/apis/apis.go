package api

import (
	handlersJsonPlaceHolders "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/handlers"
	jsonPlaceHoldersProvider "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider/json-placeholders"
	mongoProvider "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider/mongo"
	structs "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/structs"
	customMiddlewares "github/leonardoas10/go-provider-pattern/src/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JsonplaceholdersGroup(e *echo.Group)  {
	e.GET("/:id", func(c echo.Context) error {
		return handlersJsonPlaceHolders.GetJsonPlaceHolder(jsonPlaceHoldersProvider.NewProvider(), c)
	}, customMiddlewares.ParamsValidatorMiddleware("id"), middleware.Logger())

	e.GET("", func(c echo.Context) error {
		return handlersJsonPlaceHolders.GetJsonPlaceHolders(mongoProvider.NewProvider(), c)
	}, middleware.Logger())

	e.PUT("", func(c echo.Context) error {
		return handlersJsonPlaceHolders.UpdateJsonPlaceHolder(jsonPlaceHoldersProvider.NewProvider(), c)
	}, customMiddlewares.BodyValidatorMiddleware(&structs.RequestUpdateJsonPlaceHolder{}), middleware.Logger())

	e.GET("/concurrent-change-titles", func(c echo.Context) error {
		return handlersJsonPlaceHolders.ConcurrentChangeTitles(jsonPlaceHoldersProvider.NewProvider(), c)
	}, middleware.Logger())
}