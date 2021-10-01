package usecase

import (
	"service/internal/repo"
	"service/models"
)

type PersonUsecase struct {
	repo repo.IRepo
}

func NewPersonUsecase(repo repo.IRepo) *PersonUsecase {
	return &PersonUsecase{repo: repo}
}

func (pu *PersonUsecase) GetPersonsList() ([]models.Person, models.StatusCode) {
	return pu.repo.GetPersonsList()
}

func (pu *PersonUsecase) CreatePerson(person models.PersonRequest) (models.Person, models.StatusCode) {
	return pu.repo.CreatePerson(person)
}
func (pu *PersonUsecase) RemovePerson(person models.Person) models.StatusCode {
	return pu.repo.DropPerson(person)
}
func (pu *PersonUsecase) UpdatePerson(person models.Person) models.StatusCode {
	return pu.repo.UpdatePerson(person)
}
func (pu *PersonUsecase) GetPerson(person models.Person) (models.Person, models.StatusCode) {
	return pu.repo.GetPerson(person)
}
