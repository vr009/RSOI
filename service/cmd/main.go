package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"service/config"
	"service/internal/delivery"
	"service/internal/repo"
	usecase2 "service/internal/usecase"
	"service/middleware"
)

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

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

	srv := &http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	log.Fatal(srv.ListenAndServe())
}
