package db

import (
	"sandbox/pkg/todo"
)

var db *Database

type Database struct {
	Todos []todo.Todo
}

func Get() *Database {
	if db == nil {
		db = &Database{
			Todos: []todo.Todo{},
		}
	}
	return db
}

func (db Database) TodoAll() []todo.Todo {
	return db.Todos
}

func (db *Database) TodoAdd(item todo.Todo) {
	db.Todos = append(db.Todos, item)
}

func (db *Database) TodoToggle(id string) {
	for i := 0; i < len(db.Todos); i++ {
		if db.Todos[i].Id == id {
			db.Todos[i].Complete = !db.Todos[i].Complete
		}
	}
}

func (db Database) findTodoById(id string) int {
	for i, todo := range db.Todos {
		if todo.Id == id {
			return i
		}
	}

	return -1
}

func (db *Database) TodoRemove(id string) {
	index := db.findTodoById(id)
	if index == -1 {
		// TODO: handle error
		return
	}

	db.Todos = append(db.Todos[:index], db.Todos[index+1:]...)
}
