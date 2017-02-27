package controllers

import (
	"net/http"
	"strconv"

	"time"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Todo struct {
	ID   int       `json:"id"`
	Todo string    `json:"todo"`
	Done bool      `json:"done"`
	Due  time.Time `json:"due"`
}

type TodoControllerImpl struct {
	controller string
	todos      []Todo
	r          *render.Render
}

// NewTodoController creates a new TodoController
func NewTodoController() *TodoControllerImpl {
	return &TodoControllerImpl{
		controller: "TodoController",
		todos: []Todo{
			Todo{1, "Write some Go", true, time.Now()},
			Todo{2, "Release on Github", false, time.Now()},
			Todo{3, "Order Go Gopher mascot", false, time.Now()},
		},
	}
}

// Register registers controller methods with router
func (tc *TodoControllerImpl) Register(router *mux.Router, r *render.Render) {
	tc.r = r
	router.HandleFunc("/todos/{id:[0-9]+}", tc.findOne)
	router.HandleFunc("/todos", tc.create).Methods("POST")
	router.HandleFunc("/todos", tc.find)
}

func (tc *TodoControllerImpl) find(w http.ResponseWriter, r *http.Request) {
	tc.r.JSON(w, http.StatusOK, tc.todos)
}

func (tc *TodoControllerImpl) findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range tc.todos {
		if item.ID == id {
			tc.r.JSON(w, http.StatusOK, item)
			return
		}
	}

	tc.r.JSON(w, http.StatusNotFound, nil)
	return
}

func (tc *TodoControllerImpl) create(res http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)

	var newTodo Todo
	dec.Decode(&newTodo)

	nextID := len(tc.todos) + 1
	newTodo.ID = nextID

	if newTodo.Due == (Todo{}.Due) {
		newTodo.Due = time.Now().AddDate(0, 0, 10) // add 10 days
	}

	tc.todos = append(tc.todos, newTodo)
	tc.r.JSON(res, http.StatusCreated, newTodo)
}
