package main

import "slices"

func NewStore() *Store {
	return &Store{books: make(map[int]Book), nextID: 1}
}

func (s *Store) GetAll() []Book {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	allBooks := make([]Book, 0, len(s.books))
	for _, b := range s.books {
		allBooks = append(allBooks, b)
	}
	slices.SortFunc(allBooks, func(a, b Book) int { return a.ID - b.ID })
	return allBooks
}

func (s *Store) Get(id int) (Book, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	book, found := s.books[id]
	return book, found
}

func (s *Store) Create(book Book) Book {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	book.ID = s.nextID
	s.books[s.nextID] = book
	s.nextID++
	return book
}

func (s *Store) Update(id int, book Book) (Book, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, found := s.books[id]
	if !found {
		return book, false
	}
	book.ID = id
	s.books[id] = book
	return book, true
}

func (s *Store) Delete(id int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, found := s.books[id]
	if !found {
		return false
	}

	delete(s.books, id)
	return true
}
