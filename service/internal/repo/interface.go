package repo

import "service/internal/models"

type IUsecase interface {
	CreatePerson(person models.Person) error
	DropPerson(person models.Person) error
	UpdatePerson(person models.Person) error
	GetPersonsList(person models.Person) error
	GetPerson(person models.Person) error
}
