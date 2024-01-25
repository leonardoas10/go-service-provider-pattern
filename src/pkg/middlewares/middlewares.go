package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
			// Create a new instance of the expected struct for each request
			expectedStructType := reflect.TypeOf(expectedStruct)
			newExpectedStruct := reflect.New(expectedStructType.Elem()).Interface()

			// Read the content of the request body
			body, err := ioutil.ReadAll(c.Request().Body)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// Create a new reader with the body content
			bodyReader := bytes.NewReader(body)

			// Bind the new reader to the new instance of the expected struct
			if err := json.NewDecoder(bodyReader).Decode(&newExpectedStruct); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// Use the validator to check if the incoming struct is valid
			if err := validate.Struct(newExpectedStruct); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			// Print the struct to the console
			fmt.Printf("Request Body: %+v\n", newExpectedStruct)

			// Create a new request with the modified body reader
			newRequest := c.Request().WithContext(c.Request().Context())
			newRequest.Body = ioutil.NopCloser(bytes.NewReader(body))

			// Set the new request in the context
			c.SetRequest(newRequest)

			// Call the next handler in the chain
			return next(c)
		}
	}
}



