package delivery

import (
	"net/http"
	"service/internal/usecase"
)

type PersonHandler struct {
	usecase usecase.IUsecase
}

func NewPersonHandler(ucase usecase.IUsecase) *PersonHandler {
	return &PersonHandler{
		usecase: ucase,
	}
}

func (ph *PersonHandler) AddPerson(w http.ResponseWriter, r *http.Request) {

}

func (ph *PersonHandler) RemovePerson(w http.ResponseWriter, r *http.Request) {

}

func (ph *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {

}

func (ph *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {

}

func (ph *PersonHandler) GetPersonsList(w http.ResponseWriter, r *http.Request) {

}
