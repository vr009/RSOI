package repo

import (
	models2 "service/models"
)

type IRepo interface {
	CreatePerson(person models2.PersonRequest) (models2.Person, models2.StatusCode)
	DeletePerson(person models2.Person) models2.StatusCode
	UpdatePerson(person models2.Person) models2.StatusCode
	GetPersonsList() ([]models2.Person, models2.StatusCode)
	GetPerson(person models2.Person) (models2.Person, models2.StatusCode)
}
