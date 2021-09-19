package usecase

import "service/internal/models"

type IUsecase interface {
	CreatePerson(person models.Person) error
	RemovePerson(person models.Person) error
	UpdatePerson(person models.Person) error
	GetPerson(person models.Person) error
	GetPersonsList() ([]models.Person, error)
}
