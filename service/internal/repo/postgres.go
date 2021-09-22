package repo

import (
	"service/internal/models"
	"service/internal/usecase"
)

type PersonRepo struct {
	uc *usecase.IUsecase
}

func NewPersonRepo(uc *usecase.IUsecase) *PersonRepo {
	return &PersonRepo{uc: uc}
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
