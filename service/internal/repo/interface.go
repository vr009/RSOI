package repo

import "service/internal/models"

type IUsecase interface {
	CreatePerson(person models.Person) models.StatusCode
	DropPerson(person models.Person) models.StatusCode
	UpdatePerson(person models.Person) models.StatusCode
	GetPersonsList(person models.Person) models.StatusCode
	GetPerson(person models.Person) models.StatusCode
}
