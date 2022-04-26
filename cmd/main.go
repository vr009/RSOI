package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	log.Print("STARTING")
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}
	connString, err := config.GetConnectionString()
	if err != nil {
		connString = "user=postgres password=postgres host=postgres port=5432 dbname=postgres"
	}
	conn, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		print(err.Error())
	}
	repom := repo.NewPersonRepo(conn)
	usecase := usecase2.NewPersonUsecase(repom)
	handler := delivery.NewPersonHandler(usecase)

	m := middleware.NewMetricsMiddleware()
	m.Register("library_system" + os.Getenv("VER"))

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

	api.Use(m.LogMetrics)
	api.PathPrefix("/api/metrics").Handler(promhttp.Handler())

	http.Handle("/", r)
	srv := &http.Server{Handler: r, Addr: fmt.Sprintf(":%s", port)}
	log.Print("Server running at ", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
