package blacklist

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/models"
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/repository"
)

type blackListUsecase struct {
	blackListRepository repository.BlackListRepository
}

func NewBlackListUsecase(blackListRepository repository.BlackListRepository) *blackListUsecase {
	return &blackListUsecase{blackListRepository: blackListRepository}
}

func (b *blackListUsecase) ShowBlackList(option string) ([]*models.BlackList, error) {

	return b.blackListRepository.ShowBlackList(option)
}