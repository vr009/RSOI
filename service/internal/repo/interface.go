package repo

type IUsecase interface {
	CreatePerson() error
	DropPerson() error
	UpdatePerson() error
	GetPersonsList() error
	GetPerson() error
}
