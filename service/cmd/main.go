package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"service/config"
	"service/internal/delivery"
	"service/internal/repo"
	usecase2 "service/internal/usecase"
	"service/middleware"
)

func main() {
	connString, err := config.GetConnectionString()
	if err != nil {
		panic(err.Error())
	}
	conn, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		panic(err.Error())
	}

	repom := repo.NewPersonRepo(conn)
	usecase := usecase2.NewPersonUsecase(repom)
	handler := delivery.NewPersonHandler(usecase)

	r := mux.NewRouter()
	r.Use(middleware.CORSMiddleware)
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
