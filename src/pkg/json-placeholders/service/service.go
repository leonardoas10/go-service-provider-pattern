package service

import (
	"fmt"
	models "github/leonardoas10/go-provider-pattern/src/pkg/json-placeholders/models"
)

type JsonPlaceHoldersProvider interface {
	GetJsonPlaceHolders() ([]models.JsonPlaceHolders, int, error)
	GetJsonPlaceHolder(int) (models.JsonPlaceHolders, int, error)
}

type service struct {
	jsonPlaceHoldersProvider JsonPlaceHoldersProvider
}

func NewService(j JsonPlaceHoldersProvider) *service {
	return &service{
		jsonPlaceHoldersProvider: j,
	}
}

func (s *service) WhoAreThey() ([]models.JsonPlaceHolders, int, error) {
	users, status, err := s.jsonPlaceHoldersProvider.GetJsonPlaceHolders()
	if err != nil {
		return users, status, fmt.Errorf("WhatToWear: %w", err)
	}

	return users, status, nil
}

func (s *service) WhoIs(id int) (models.JsonPlaceHolders, int, error) {
	user, status, err := s.jsonPlaceHoldersProvider.GetJsonPlaceHolder(id)
	if err != nil {
		return user, status, fmt.Errorf("WhatToWear: %w", err)
	}

	return user, status, nil
}