package repo

import (
	"service/models"
)

const (
	ADDQUERY    = ""
	DELETEQUERY = ""
	UPDATEQUERY = ""
	GETQUERY    = ""
)

type PersonRepo struct {
}

func NewPersonRepo() *PersonRepo {
	return &PersonRepo{}
}

func (pr *PersonRepo) CreatePerson(person models.PersonRequest) (models.Person, models.StatusCode) {
	return models.Person{}, models.Okay
}
func (pr *PersonRepo) DeletePerson(person models.Person) models.StatusCode {
	return models.Okay
}
func (pr *PersonRepo) UpdatePerson(person models.Person) models.StatusCode {
	return models.Okay
}
func (pr *PersonRepo) GetPerson(person models.Person) (models.Person, models.StatusCode) {
	return models.Person{}, models.Okay
}
func (pr *PersonRepo) GetPersonsList() ([]models.Person, models.StatusCode) {
	return nil, models.Okay
}
