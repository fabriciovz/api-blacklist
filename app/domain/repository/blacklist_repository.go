package repository

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/models"
)

type BlackListRepository interface {
	ShowBlackList(option string) ([]*models.BlackList, error)
}
