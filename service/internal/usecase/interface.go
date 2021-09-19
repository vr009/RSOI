package usecase

type IUsecase interface {
	CreatePerson() error
	RemovePerson() error
	UpdatePerson() error
	GetPersonsList() error
	GetPerson() error
}
