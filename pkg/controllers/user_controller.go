package controllers

import (
	"encoding/json"
	"net/http"
	"sandbox/pkg/db"
	"sandbox/pkg/templates"
	"sandbox/pkg/user"
)

type ListAllUsersTemplateArgs struct {
	Users []user.User
}

func ListAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	args := ListAllUsersTemplateArgs{
		Users: db.Get().UsersAll(),
	}

	templates.Get().ExecuteTemplate(w, "home.page.html", args)
}

type AddUserRequest struct {
	Email string `json:"email"`
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var body AddUserRequest
	json.NewDecoder(r.Body).Decode(&body)

	user := user.New(body.Email, true)
	db.Get().UserAdd(user)

	templates.Get().ExecuteTemplate(w, "user_row.partial.html", user)
}

func ToggleUserStatusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Get().UserActiveToggle(id)

	args := ListAllUsersTemplateArgs{
		Users: db.Get().UsersAll(),
	}

	templates.Get().ExecuteTemplate(w, "user_table_rows.partial.html", args)
}

func UserRemoveHandle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db.Get().UserRemove(id)

	args := ListAllUsersTemplateArgs{
		Users: db.Get().UsersAll(),
	}

	templates.Get().ExecuteTemplate(w, "user_table_rows.partial.html", args)
}
