package web

import (
	"github.com/badmuts/go-todo-rest/controllers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type Server struct {
	*negroni.Negroni
}

func NewServer() *Server {
	r := render.New()
	router := mux.NewRouter()

	todoController := controllers.NewTodoController()
	todoController.Register(router, r)

	server := Server{negroni.Classic()}
	server.UseHandler(router)

	return &server
}
