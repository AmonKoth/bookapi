package main

import (
	"log"
	"net/http"
)

func main() {
	h := &Handler{store: NewStore()}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", h.ListBooks)
	mux.HandleFunc("GET /books/{id}", h.GetBook)
	mux.HandleFunc("POST /books", h.CreateBook)
	mux.HandleFunc("PUT /books/{id}", h.UpdateBook)
	mux.HandleFunc("DELETE /books/{id}", h.DeleteBook)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
