package usecase

import "service/internal/models"

type IUsecase interface {
	CreatePerson(person models.Person) models.StatusCode
	RemovePerson(person models.Person) models.StatusCode
	UpdatePerson(person models.Person) models.StatusCode
	GetPerson(person models.Person) models.StatusCode
	GetPersonsList() ([]models.Person, models.StatusCode)
}
