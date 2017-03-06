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
	todos      map[string]Todo
	r          *render.Render
}

// NewTodoController creates a new TodoController
func NewTodoController() *TodoControllerImpl {
	return &TodoControllerImpl{
		controller: "TodoController",
		todos: map[string]Todo{
			"1": Todo{1, "Write some Go", true, time.Now()},
			"2": Todo{2, "Release on Github", false, time.Now()},
			"3": Todo{3, "Order Go Gopher mascot", false, time.Now()},
		},
	}
}

// Register registers controller methods with router
func (tc *TodoControllerImpl) Register(router *mux.Router, r *render.Render) {
	tc.r = r
	router.HandleFunc("/todos/{id:[0-9]+}", tc.update).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", tc.findOne)
	router.HandleFunc("/todos", tc.create).Methods("POST")
	router.HandleFunc("/todos", tc.find)
}

// find finds all Todo's in tc.todos
func (tc *TodoControllerImpl) find(w http.ResponseWriter, r *http.Request) {
	todos := make([]Todo, 0, len(tc.todos))
	for _, todo := range tc.todos {
		todos = append(todos, todo)
	}
	tc.r.JSON(w, http.StatusOK, todos)
}

// findOne finds Todo in tc.todo with given ID
func (tc *TodoControllerImpl) findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo, found := tc.todos[vars["id"]]

	if !found {
		tc.r.JSON(w, http.StatusNotFound, []int{})
		return
	}

	tc.r.JSON(w, http.StatusOK, todo)
	return
}

// create creates a new Todo and appends it to tc.todos
func (tc *TodoControllerImpl) create(res http.ResponseWriter, req *http.Request) {
	var newTodo Todo
	dec := json.NewDecoder(req.Body)
	dec.Decode(&newTodo)

	nextID := len(tc.todos) + 1
	newTodo.ID = nextID

	if newTodo.Due == (Todo{}.Due) {
		newTodo.Due = time.Now().AddDate(0, 0, 10) // add 10 days
	}

	tc.todos[strconv.Itoa(nextID)] = newTodo
	tc.r.JSON(res, http.StatusCreated, tc.todos[strconv.Itoa(nextID)])
}

// update updates Todo in tc.todos
func (tc *TodoControllerImpl) update(res http.ResponseWriter, req *http.Request) {
	var todo Todo
	dec := json.NewDecoder(req.Body)
	dec.Decode(&todo)

	tc.todos[strconv.Itoa(todo.ID)] = todo
	tc.r.JSON(res, http.StatusOK, todo)
}
