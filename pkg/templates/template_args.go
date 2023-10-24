package templates

import (
	"net/http"
	"sandbox/pkg/user"
)

type UserTableRows struct {
	Users []user.User
}

func UserRowPartialTemplate(w http.ResponseWriter, user user.User) {
	Get().ExecuteTemplate(w, "user_row.partial.html", user)
}

func UserTableRowsPartialTemplate(w http.ResponseWriter, users []user.User) {
	Get().ExecuteTemplate(w, "user_table_rows.partial.html", UserTableRows{
		Users: users,
	})
}

func HomePageTemplate(w http.ResponseWriter, users []user.User) {
	Get().ExecuteTemplate(w, "home.page.html", UserTableRows{
		Users: users,
	})
}
