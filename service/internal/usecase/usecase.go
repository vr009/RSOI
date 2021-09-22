package usecase

import (
	"service/internal/models"
	"service/internal/repo"
)

type PersonUsecase struct {
	repo *repo.PersonRepo
}

func NewPersonUsecase(repo *repo.PersonRepo) *PersonUsecase {
	return &PersonUsecase{repo: repo}
}

func CreatePerson(person models.Person) error {
	return nil
}
func RemovePerson(person models.Person) error {
	return nil
}
func UpdatePerson(person models.Person) error {
	return nil
}
func GetPerson(person models.Person) error {
	return nil
}
func GetPersonsList() ([]models.Person, error) {
	return nil, nil
}
