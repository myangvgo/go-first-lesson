package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

// Create implements store.Store
func (ms *MemStore) Create(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[book.Id]; ok {
		return mystore.ErrExist
	}

	newBook := *book
	ms.books[book.Id] = &newBook

	return nil
}

// Delete implements store.Store
func (ms *MemStore) Delete(id string) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.books[id]; !ok {
		return mystore.ErrExist
	}

	delete(ms.books, id)
	return nil
}

// Get implements store.Store
func (ms *MemStore) Get(id string) (mystore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	t, ok := ms.books[id]

	if ok {
		return *t, nil
	}

	return mystore.Book{}, mystore.ErrNotFound
}

// GetAll implements store.Store
func (ms *MemStore) GetAll() ([]mystore.Book, error) {
	ms.RLock()
	defer ms.RUnlock()

	allBooks := make([]mystore.Book, 0, len(ms.books))
	for _, book := range ms.books {
		allBooks = append(allBooks, *book)
	}

	return allBooks, nil
}

// Update implements store.Store
func (ms *MemStore) Update(book *mystore.Book) error {
	ms.Lock()
	defer ms.Unlock()

	oldBook, ok := ms.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	newBook := *oldBook
	if book.Name != "" {
		newBook.Name = book.Name
	}

	if book.Authors != nil {
		newBook.Authors = book.Authors
	}

	if book.Press != "" {
		newBook.Press = book.Press
	}

	ms.books[book.Id] = &newBook

	return nil
}
