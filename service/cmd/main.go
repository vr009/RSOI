package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"service/internal/delivery"
	"service/internal/repo"
	usecase2 "service/internal/usecase"
)

func main() {
	repom := repo.NewPersonRepo()
	usecase := usecase2.NewPersonUsecase(repom)
	handler := delivery.NewPersonHandler(usecase)

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/persons", handler.GetPersonsList).Methods(http.MethodGet)
		api.HandleFunc("/persons", handler.AddPerson).Methods(http.MethodPost)
		api.HandleFunc("/persons/{id:[0-9]+}", handler.GetPerson).Methods(http.MethodGet)
		api.HandleFunc("/persons/{id:[0-9]+}", handler.UpdatePerson).Methods(http.MethodPatch)
		api.HandleFunc("/persons/{id:[0-9]+}", handler.RemovePerson).Methods(http.MethodDelete)
	}

	srv := &http.Server{Handler: r, Addr: ":8080"}
	log.Fatal(srv.ListenAndServe())
}
