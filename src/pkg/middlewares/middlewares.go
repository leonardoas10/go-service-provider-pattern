package middlewares

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func BodyValidatorMiddleware(expectedStruct interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Create a new instance of the expected struct
			expectedStructType := reflect.TypeOf(expectedStruct)
			newExpectedStruct := reflect.New(expectedStructType.Elem()).Interface()

			// Read the content of the request body and bind it to the new instance of the expected struct
			if err := c.Bind(newExpectedStruct); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// Use the validator to check if the incoming struct is valid
			if err := validate.Struct(newExpectedStruct); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// Print the struct to the console
			fmt.Printf("Request Body: %+v\n", newExpectedStruct)

			// Call the next handler in the chain
			return next(c)
		}
	}
}
