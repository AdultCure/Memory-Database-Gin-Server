package service

import "http-type-one/repository"

type AccountIntegrationService struct {
	repo repository.AccountIntegration
}

func NewAccountIntegrationService(repo repository.AccountIntegration) *AccountIntegrationService {
	return &AccountIntegrationService{repo: repo}
}
