package mapper

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/models"
	"bitbucket.org/fabribraguev/api-toolbox/app/infrastructure/data/db/mysql/entity"
)

func BlackListEntityToBlackList(bl *entity.BlackList) *models.BlackList {
	blackList := &models.BlackList{
		ID:      bl.ID,
		Sku:     bl.Sku,
	}
	return blackList
}