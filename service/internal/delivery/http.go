package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"service/internal/middleware"
	"service/internal/models"
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
	person := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !middleware.PersonIsValid(person) {
		w.Write([]byte("Data is not full\n"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = ph.usecase.CreatePerson(person); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
