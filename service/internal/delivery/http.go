package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"service/internal/usecase"
	"service/models"
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

func (ph *PersonHandler) GetPersonsList(w http.ResponseWriter, r *http.Request) {
	users, status := ph.usecase.GetPersonsList()
	if status == models.Okay {
		body, err := json.Marshal(users)
		if err != nil {
			Response(w, models.InternalError, "", nil)
			return
		}
		Response(w, models.Okay, "All Persons", body)
		return
	}
	Response(w, models.BadRequest, "", nil)
}

func (ph *PersonHandler) AddPerson(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		Response(w, models.InternalError, "", nil)
		return
	}
	NewPerson, status := ph.usecase.CreatePerson(person)
	if status != models.Okay {
		Response(w, models.BadRequest, "Invalid data", nil)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/api/v1/persons/"+strconv.Itoa(NewPerson.ID))
}

func (ph *PersonHandler) GetPerson(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	vars := mux.Vars(r)
	var (
		err    error
		status models.StatusCode
	)
	person.ID, err = strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	person, status = ph.usecase.GetPerson(person)
	if status == models.Okay {
		body, err := json.Marshal(person)
		if err != nil {
			Response(w, models.InternalError, "", nil)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	body, err := json.Marshal(models.ErrorResponse{Message: "Not found person for id"})
	w.Write(body)
}

func (ph *PersonHandler) RemovePerson(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	vars := mux.Vars(r)
	var (
		err    error
		status models.StatusCode
	)
	person.ID, err = strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	status = ph.usecase.RemovePerson(person)
	if status != models.Okay {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	w.Header().Set("Description", "Person for ID was removed")
}

func (ph *PersonHandler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	vars := mux.Vars(r)
	var (
		err    error
		status models.StatusCode
	)
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	person.ID, err = strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status = ph.usecase.UpdatePerson(person)

	switch status {
	case models.Okay:
		body, _ := json.Marshal(person)
		Response(w, status, "Person for ID was updated", body)
		return
	case models.NotFound:
		body, _ := json.Marshal(models.ErrorResponse{Message: "Not Found Person for ID"})
		Response(w, status, "Not found Person for ID", body)
		return
	case models.BadRequest:
		body, _ := json.Marshal(models.ValidationErrorResponse{Message: "Invalid data"})
		Response(w, status, "Invalid data", body)
		return
	}
}
