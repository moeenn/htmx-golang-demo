package controllers

import (
	"encoding/json"
	"net/http"
	"sandbox/pkg/db"
	"sandbox/pkg/templates"
	"sandbox/pkg/todo"
)

type ListAllTodosTemplateArgs struct {
	Todos []todo.Todo
}

func ListAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	args := ListAllTodosTemplateArgs{
		Todos: db.Get().TodoAll(),
	}

	templates.Get().ExecuteTemplate(w, "home.page.html", args)
}

type AddTodoRequest struct {
	New string `json:"new"`
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	var body AddTodoRequest
	json.NewDecoder(r.Body).Decode(&body)

	todo := todo.New(body.New, false)
	db.Get().TodoAdd(todo)

	templates.Get().ExecuteTemplate(w, "todoItem.partial.html", todo)
}

func ToggleTodoStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Get().TodoToggle(id)

	args := ListAllTodosTemplateArgs{
		Todos: db.Get().TodoAll(),
	}

	templates.Get().ExecuteTemplate(w, "todolist.partial.html", args)
}

func TodoRemoveHandle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Get().TodoRemove(id)

	args := ListAllTodosTemplateArgs{
		Todos: db.Get().TodoAll(),
	}

	templates.Get().ExecuteTemplate(w, "todolist.partial.html", args)
}
