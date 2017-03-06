package controllers

import (
	"net/http"
	"strconv"

	"time"

	"encoding/json"

	"github.com/badmuts/go-todo-rest/models"
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
	todos      *models.Todos
	r          *render.Render
}

// NewTodoController creates a new TodoController
func NewTodoController() *TodoControllerImpl {
	return &TodoControllerImpl{
		controller: "TodoController",
		todos:      models.NewTodos(),
	}
}

// Register registers controller methods with router
func (tc *TodoControllerImpl) Register(router *mux.Router, r *render.Render) {
	tc.r = r
	router.HandleFunc("/todos/{id:[0-9]+}", tc.delete).Methods("DELETE")
	router.HandleFunc("/todos/{id:[0-9]+}", tc.update).Methods("PUT")
	router.HandleFunc("/todos/{id:[0-9]+}", tc.findOne)
	router.HandleFunc("/todos", tc.create).Methods("POST")
	router.HandleFunc("/todos", tc.find)
}

// find finds all Todo's in tc.todos
func (tc *TodoControllerImpl) find(w http.ResponseWriter, r *http.Request) {
	tc.r.JSON(w, http.StatusOK, tc.todos.Flatten())
}

func (tc *TodoControllerImpl) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	tc.todos.Remove(id)
	tc.r.JSON(w, http.StatusOK, tc)
}

// findOne finds Todo in tc.todo with given ID
func (tc *TodoControllerImpl) findOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	todo, found := tc.todos.Get(id)

	if !found {
		tc.r.JSON(w, http.StatusNotFound, tc)
		return
	}

	tc.r.JSON(w, http.StatusOK, todo)
	return
}

// create creates a new Todo and appends it to tc.todos
func (tc *TodoControllerImpl) create(res http.ResponseWriter, req *http.Request) {
	var newTodo models.Todo
	dec := json.NewDecoder(req.Body)
	dec.Decode(&newTodo)

	newTodo.ID = tc.todos.Length() + 1

	if newTodo.Due == (Todo{}.Due) {
		newTodo.Due = time.Now().AddDate(0, 0, 10) // add 10 days
	}

	tc.todos.Add(newTodo)
	todo, _ := tc.todos.Get(newTodo.ID)

	tc.r.JSON(res, http.StatusCreated, todo)
}

// update updates Todo in tc.todos
func (tc *TodoControllerImpl) update(res http.ResponseWriter, req *http.Request) {
	var todo models.Todo
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	dec := json.NewDecoder(req.Body)
	dec.Decode(&todo)

	oldTodo, _ := tc.todos.Get(id)

	tc.r.JSON(res, http.StatusOK, tc.todos.Update(oldTodo, todo))
}
