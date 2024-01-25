package api

import (
	handlersJsonPlaceHolders "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/handlers"
	structs "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/structs"
	customMiddlewares "github/leonardoas10/go-provider-pattern/src/pkg/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JsonplaceholdersGroup(e *echo.Group)  {
	e.GET("", handlersJsonPlaceHolders.GetJsonPlaceHolders, middleware.Logger())
	e.POST("", handlersJsonPlaceHolders.GetJsonPlaceHolder,  customMiddlewares.BodyValidatorMiddleware(&structs.RequestId{}), middleware.Logger())
}