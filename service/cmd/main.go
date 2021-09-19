package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//repom := repo.NewPGRepo()
	//repom.InitDB()
	//defer repom.Close()
	//usecase := usecase.NewChatUseCase(repom)
	//handler := delivery.NewChatHandler(usecase)

	r := mux.NewRouter()

	//r.HandleFunc("/chat/chatlist", handler.GetChatList)
	//r.HandleFunc("/chat/conv", handler.GetChat).Methods("GET")
	//r.HandleFunc("/chat/put/message", handler.PostMessage)

	srv := &http.Server{Handler: r, Addr: ":8000"}

	log.Fatal(srv.ListenAndServe())
}
