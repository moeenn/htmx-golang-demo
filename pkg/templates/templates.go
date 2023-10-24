package templates

import (
	"html/template"
	"net/http"
	"sandbox/pkg/user"
)

var tmpl *template.Template

func init() {
	if tmpl == nil {
		var err error
		tmpl, err = template.ParseGlob("./pkg/templates/views/**/*.html")
		if err != nil {
			panic("failed to load template: " + err.Error())
		}
	}
}

type UserTableRows struct {
	Users []user.User
}

func UserRowPartialTemplate(w http.ResponseWriter, user user.User) {
	tmpl.ExecuteTemplate(w, "user_row.partial.html", user)
}

func UserTableRowsPartialTemplate(w http.ResponseWriter, users []user.User) {
	tmpl.ExecuteTemplate(w, "user_table_rows.partial.html", UserTableRows{
		Users: users,
	})
}

func HomePageTemplate(w http.ResponseWriter, users []user.User) {
	tmpl.ExecuteTemplate(w, "home.page.html", UserTableRows{
		Users: users,
	})
}
