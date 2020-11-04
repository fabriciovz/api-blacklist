package usecase

type BlackListUsecase interface {
	ShowBlackList(option string) (string, error)
	DeleteExcludeItems()  error
}

