package db

import (
	"sandbox/pkg/user"
)

var db *Database

type Database struct {
	Users []user.User
}

func Get() *Database {
	if db == nil {
		db = &Database{
			Users: []user.User{},
		}
	}
	return db
}

func (db Database) UsersAll() []user.User {
	return db.Users
}

func (db *Database) UserAdd(item user.User) {
	db.Users = append(db.Users, item)
}

func (db *Database) UserActiveToggle(id string) {
	for i := 0; i < len(db.Users); i++ {
		if db.Users[i].Id == id {
			db.Users[i].Active = !db.Users[i].Active
		}
	}
}

func (db Database) findUserById(id string) int {
	for i, user := range db.Users {
		if user.Id == id {
			return i
		}
	}
	return -1
}

func (db *Database) UserRemove(id string) {
	index := db.findUserById(id)
	if index == -1 {
		return
	}

	db.Users = append(db.Users[:index], db.Users[index+1:]...)
}
