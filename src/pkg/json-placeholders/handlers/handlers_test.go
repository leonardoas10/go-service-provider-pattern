package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	jsonPlaceHoldersProvider "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider/json-placeholders"
	mongoProvider "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/provider/mongo"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetJsonPlaceHoldersWithJsonPlaceHoldersProvider(t *testing.T) {
	// Create a new echo instance
	e := echo.New()

	// Create a request using http.NewRequest
	req := httptest.NewRequest(http.MethodGet, "/json-placeholders", nil)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Create an echo context
	c := e.NewContext(req, rec)

	// Call the handler function
	if assert.NoError(t, GetJsonPlaceHolders(jsonPlaceHoldersProvider.NewProvider(), c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Unmarshal the response body into a JsonPlaceHolder struct
		var jsonResponse []models.JsonPlaceHolder
		err := json.Unmarshal([]byte(rec.Body.String()), &jsonResponse)

		// Assert that unmarshaling was successful
		assert.NoError(t, err)

		// Assert that the length of the array is greater than 0
		assert.Greater(t, len(jsonResponse), 0)

		// Assert that the element at the chosen index matches the JsonPlaceHolder struct
		assert.IsType(t, models.JsonPlaceHolder{}, jsonResponse[0])
	}
}

func TestGetJsonPlaceHoldersWithMongoProvider(t *testing.T) {
	// Create a new echo instance
	e := echo.New()

	// Create a request using http.NewRequest
	req := httptest.NewRequest(http.MethodGet, "/json-placeholders", nil)

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Create an echo context
	c := e.NewContext(req, rec)

	// Call the handler function
	if assert.NoError(t, GetJsonPlaceHolders(mongoProvider.NewProvider(), c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		// Unmarshal the response body into a JsonPlaceHolder struct
		var jsonResponse []models.JsonPlaceHolder
		err := json.Unmarshal([]byte(rec.Body.String()), &jsonResponse)

		// Assert that unmarshaling was successful
		assert.NoError(t, err)

		// Assert that the length of the array is greater than 0
		assert.Greater(t, len(jsonResponse), 0)

		// Assert that the element at the chosen index matches the JsonPlaceHolder struct
		assert.IsType(t, models.JsonPlaceHolder{}, jsonResponse[0])
	}
}
