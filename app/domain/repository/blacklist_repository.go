package repository

type BlackListRepository interface {
	ShowBlackList(option string) (string, error)
}
