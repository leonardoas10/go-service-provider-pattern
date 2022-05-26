package jsonplaceholders

import (
	"encoding/json"
	env "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/common/env"
	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"
)

type provider struct {}

func NewProvider() *provider  {
	return &provider{}
}

func (p *provider) GetJsonPlaceHolders() ([]models.JsonPlaceHolder, int, error) {
	res, err := http.Get(env.GetEnvVariable("URL"))
	
	if err != nil {
		return []models.JsonPlaceHolder{}, 500, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)


	var users []models.JsonPlaceHolder
	if err := json.Unmarshal([]byte(bodyString), &users); err != nil {
        panic(err)
    }

	return users, res.StatusCode, nil
}

func (p *provider) GetJsonPlaceHolder(id int) (models.JsonPlaceHolder, int, error) {
	res, err := http.Get(env.GetEnvVariable("URL") + strconv.Itoa(id))
	
	if err != nil {
		return models.JsonPlaceHolder{}, 500, err
	}

	r := new(models.JsonPlaceHolder)
	errors := json.NewDecoder(res.Body).Decode(r)

	if errors != nil {
		return models.JsonPlaceHolder{}, 500, err
	}

	return models.JsonPlaceHolder{
		UserId: r.UserId,
		Id: r.Id,
		Title: r.Title,
		Completed: r.Completed,
	}, res.StatusCode, nil
}
