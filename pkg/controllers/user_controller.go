package controllers

import (
	"encoding/json"
	"net/http"
	"sandbox/pkg/db"
	"sandbox/pkg/templates"
	"sandbox/pkg/user"
)

type AddUserRequest struct {
	Email string `json:"email"`
}

func ListAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	templates.HomePageTemplate(w, db.Get().UsersAll())
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var body AddUserRequest
	json.NewDecoder(r.Body).Decode(&body)

	user := user.New(body.Email, true)
	db.Get().UserAdd(user)

	templates.UserRowPartialTemplate(w, user)
}

func ToggleUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Get().UserActiveToggle(id)
	templates.UserTableRowsPartialTemplate(w, db.Get().UsersAll())
}

func UserRemoveHandle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Get().UserRemove(id)
	templates.UserTableRowsPartialTemplate(w, db.Get().UsersAll())
}
