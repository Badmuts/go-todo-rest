package models

import (
	"time"

	"github.com/imdario/mergo"
)

type Todo struct {
	ID   int       `json:"id"`
	Todo string    `json:"todo"`
	Done bool      `json:"done"`
	Due  time.Time `json:"due"`
}

type Todos struct {
	todos map[int]Todo
}

func NewTodos() *Todos {
	return &Todos{
		todos: map[int]Todo{
			1: Todo{1, "Write some Go", true, time.Now()},
			2: Todo{2, "Release on Github", false, time.Now()},
			3: Todo{3, "Order Go Gopher mascot", false, time.Now()},
		},
	}
}

func (t *Todos) Add(todo Todo) {
	t.todos[todo.ID] = todo
}

func (t *Todos) Remove(ID int) {
	delete(t.todos, ID)
}

func (t *Todos) Get(ID int) (Todo, bool) {
	todo, found := t.todos[ID]
	return todo, found
}

func (t *Todos) Update(oldTodo Todo, newTodo Todo) Todo {
	mergo.Merge(&newTodo, oldTodo)
	t.todos[oldTodo.ID] = newTodo
	return newTodo
}

func (t *Todos) Flatten() []Todo {
	// TODO: fix iteration https://blog.golang.org/go-maps-in-action#TOC_7.
	todos := make([]Todo, 0, len(t.todos))
	for _, todo := range t.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (t *Todos) Length() int {
	return len(t.todos)
}
