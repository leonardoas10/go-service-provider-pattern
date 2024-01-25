package service

import (
	"fmt"
	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
)

type JsonPlaceHoldersProvider interface {
	GetJsonPlaceHolders() ([]models.JsonPlaceHolder, int, error)
	GetJsonPlaceHolder(int) (models.JsonPlaceHolder, int, error)
	UpdateJsonPlaceHolder(models.UpdateJsonPlaceHolder) (models.JsonPlaceHolder, int, error)
}

type service struct {
	jsonPlaceHoldersProvider JsonPlaceHoldersProvider
}

func NewService(j JsonPlaceHoldersProvider) *service {
	return &service{
		jsonPlaceHoldersProvider: j,
	}
}

func (s *service) WhoAreThey() ([]models.JsonPlaceHolder, int, error) {
	users, status, err := s.jsonPlaceHoldersProvider.GetJsonPlaceHolders()
	if err != nil {
		fmt.Printf("Error WhoAreThey %s\n", err)
		return users, status, fmt.Errorf("error: %w", err)
	}

	return users, status, nil
}

func (s *service) WhoIs(id int) (models.JsonPlaceHolder, int, error) {
	user, status, err := s.jsonPlaceHoldersProvider.GetJsonPlaceHolder(id)
	if err != nil {
		fmt.Printf("Error WhoIs %s\n", err)
		return user, status, fmt.Errorf("error: %w", err)
	}
	return user, status, nil
}

func (s *service) UpdateJsonPlaceHolder(jsonPlaceholder models.UpdateJsonPlaceHolder) (models.JsonPlaceHolder, int, error)  {
	user, status, err := s.jsonPlaceHoldersProvider.UpdateJsonPlaceHolder(jsonPlaceholder)
	if err != nil {
		fmt.Printf("Error UpdateJsonPlaceHolder %s\n", err)
		return user, status, fmt.Errorf("error: %w", err)
	}
	return user, status, nil
}