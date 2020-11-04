package usecase

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/models"
)

type BlackListUsecase interface {
	ShowBlackList(option string) ([]*models.BlackList, error)}

