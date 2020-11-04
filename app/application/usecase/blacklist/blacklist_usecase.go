package blacklist

import (
	"bitbucket.org/fabribraguev/api-toolbox/app/domain/repository"
)

type blackListUsecase struct {
	blackListRepository repository.BlackListRepository
}

func NewBlackListUsecase(blackListRepository repository.BlackListRepository) *blackListUsecase {
	return &blackListUsecase{blackListRepository: blackListRepository}
}

func (b *blackListUsecase) ShowBlackList(option string) (string, error) {

	return b.blackListRepository.ShowBlackList(option)
}

func (b *blackListUsecase) DeleteExcludeItems()  error {

	return b.blackListRepository.RemoveExcludeItems()
}
