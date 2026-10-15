package main

import "sync"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Store struct {
	mutex  sync.Mutex
	books  map[int]Book
	nextID int
}
