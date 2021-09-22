package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"service/internal/models"
	"service/internal/usecase"
	"service/middleware"
	"strconv"
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !middleware.PersonIsValid(person) {
		// TODO middleware
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if code := ph.usecase.CreatePerson(person); code != models.Okay {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Location", strconv.Itoa(person.ID))
	w.WriteHeader(http.StatusCreated)
}

func (ph *PersonHandler) RemovePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if code := ph.usecase.RemovePerson(models.Person{ID: id}); code != models.Okay {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (ph *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	person := models.Person{}
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if code := ph.usecase.UpdatePerson(person); code != models.Okay {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (ph *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {

}

func (ph *PersonHandler) GetPersonsList(w http.ResponseWriter, r *http.Request) {

}
