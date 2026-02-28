package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Task represents a to-do item
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Store holds our in-memory tasks
type Store struct {
	mu     sync.Mutex
	tasks  map[int]Task
	nextID int
}

func NewStore() *Store {
	return &Store{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func main() {
	_ = NewStore()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go CRUD App!")
	})

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
