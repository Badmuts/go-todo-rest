package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

type TodoControllerImpl struct {
	controller string
	todos      []Todo
}

func NewTodoController() *TodoControllerImpl {
	return &TodoControllerImpl{
		controller: "TodoController",
		todos: []Todo{
			Todo{1, "Write some Go", true},
			Todo{2, "Release on Github", false},
			Todo{3, "Order Go Gopher mascot", false},
		},
	}
}

func (tc *TodoControllerImpl) Register(router *mux.Router) {
	router.HandleFunc("/todos/{id:[0-9]+}", tc.findOne)
	router.HandleFunc("/todos", tc.create).Methods("POST")
	router.HandleFunc("/todos", tc.find)
}

func (tc *TodoControllerImpl) find(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(tc.todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (tc *TodoControllerImpl) findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	var todo Todo

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range tc.todos {
		if item.ID == id {
			todo = item
		}
	}

	if (Todo{}) == todo {
		w.Write([]byte("{}"))
		http.NotFound(w, r)
		return
	}

	js, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (tc *TodoControllerImpl) create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("stub create"))
}
