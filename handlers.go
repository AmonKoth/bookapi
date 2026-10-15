package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	store *Store
}

func (handler *Handler) ListBooks(writer http.ResponseWriter, request *http.Request) {
	books := handler.store.GetAll()
	WriteJSON(writer, http.StatusOK, books)
}

func (handler *Handler) GetBook(writer http.ResponseWriter, request *http.Request) {
	id, err := ParseID(request)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, "invalid book ID")
		return
	}
	book, found := handler.store.Get(id)
	if !found {
		WriteError(writer, http.StatusNotFound, "book not found")
		return
	}
	WriteJSON(writer, http.StatusOK, book)
}

func (handler *Handler) CreateBook(writer http.ResponseWriter, request *http.Request) {
	var book Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		WriteError(writer, http.StatusBadRequest, "invalid JSON")
		return
	}
	if book.Author == "" || book.Title == "" {
		WriteError(writer, http.StatusBadRequest, "title and author required")
		return
	}
	created := handler.store.Create(book)
	WriteJSON(writer, http.StatusCreated, created)
}

func (handler *Handler) UpdateBook(writer http.ResponseWriter, request *http.Request) {
	bookID, err := ParseID(request)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, "invalid book id")
		return
	}
	var book Book
	if err = json.NewDecoder(request.Body).Decode(&book); err != nil {
		WriteError(writer, http.StatusBadRequest, "invalid JSON")
		return
	}
	if book.Author == "" || book.Title == "" {
		WriteError(writer, http.StatusBadRequest, "title and author required")
		return
	}
	updated, found := handler.store.Update(bookID, book)
	if !found {
		WriteError(writer, http.StatusNotFound, "book could not be found")
		return
	}
	WriteJSON(writer, http.StatusOK, updated)
}

func (handler *Handler) DeleteBook(writer http.ResponseWriter, request *http.Request) {
	bookID, err := ParseID(request)
	if err != nil {
		WriteError(writer, http.StatusBadRequest, "invalid book id")
		return
	}
	success := handler.store.Delete(bookID)
	if !success {
		WriteError(writer, http.StatusNotFound, "book cannot be found or deleted")
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
