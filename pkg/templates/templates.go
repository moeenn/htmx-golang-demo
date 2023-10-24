package templates

import (
	"html/template"
	"io"
	"sandbox/pkg/user"
)

var tmpl *template.Template

/* templates will be parsed once at package first import */
func init() {
	if tmpl == nil {
		tmpl = template.Must(template.ParseGlob("./pkg/templates/views/**/*.html"))
	}
}

type UserTableRows struct {
	Users []user.User
}

/* all pages and partials will have their own typed functions for execution */
func UserRowPartialTemplate(w io.Writer, user user.User) {
	tmpl.ExecuteTemplate(w, "user_row.partial.html", user)
}

func UserTableRowsPartialTemplate(w io.Writer, users []user.User) {
	tmpl.ExecuteTemplate(w, "user_table_rows.partial.html", UserTableRows{
		Users: users,
	})
}

func HomePageTemplate(w io.Writer, users []user.User) {
	tmpl.ExecuteTemplate(w, "home.page.html", UserTableRows{
		Users: users,
	})
}
