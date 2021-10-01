package usecase

import (
	models2 "service/models"
)

type IUsecase interface {
	GetPersonsList() ([]models2.Person, models2.StatusCode)
	CreatePerson(person models2.Person) (models2.Person, models2.StatusCode)
	RemovePerson(person models2.Person) models2.StatusCode
	UpdatePerson(person models2.Person) models2.StatusCode
	GetPerson(person models2.Person) (models2.Person, models2.StatusCode)
}
