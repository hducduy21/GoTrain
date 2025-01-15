package handlers

import (
	"context"
	"emb/pkg/auth"
	"emb/pkg/db"
	"emb/pkg/db/ent"
	"emb/pkg/db/ent/user"
	"emb/pkg/services"
	"emb/pkg/tmpl"
	"log"
	"net/http"
	"time"
)

type Error struct {
	ErrorMessage string
}

type PageTmpl struct {
	User  *ent.User
	Auth  bool
	Tasks []*ent.Task
}

func GetUserFromContext(ctx context.Context) (*ent.User, error) {
	username, _ := ctx.Value("username").(string)
	return db.Client.User.Query().WithTasks().Where(user.Username(username)).Only(ctx)
}

func DirectToMainPage(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.GetAllTask()
	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "Error", Error{ErrorMessage: "Failed to get tasks"})
		return
	}
	tmpl.Tmpl.ExecuteTemplate(w, "Base", PageTmpl{Tasks: tasks})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "Login", Error{ErrorMessage: "Please fill the form!"})
		return
	}
	username := r.FormValue("username")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repassword := r.FormValue("repassword")
	if password != repassword {
		tmpl.Tmpl.ExecuteTemplate(w, "Login", Error{ErrorMessage: "Password and Re-password must be the same!"})
		return
	}

	log.Printf("Username: %s, Name: %s, Email: %s, Password: %s, Repassword: %s", username, name, email, password, repassword)

	_, err := db.Client.User.
		Create().
		SetUsername(username).
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(context.Background())

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "Register", Error{ErrorMessage: "Failed to create user"})
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "Login", nil)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "Login", Error{ErrorMessage: "Please fill the form!"})
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Printf("Username: %s, Password: %s", username, password)

	user, err := db.Client.User.Query().Where(user.Username(username)).Only(context.Background())

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "Login", Error{ErrorMessage: "Invalid username or password!"})
		return
	}

	if user.Password != password {
		tmpl.Tmpl.ExecuteTemplate(w, "Login", Error{ErrorMessage: "Invalid username or password!"})
		return
	}

	token, errGen1 := auth.GeneratedToken(username)
	refreshToken, errGen2 := auth.GeneratedRefreshToken()
	if (errGen1 != nil) || (errGen2 != nil) {
		tmpl.Tmpl.ExecuteTemplate(w, "Login", Error{ErrorMessage: "Internal server error!"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(20 * time.Minute),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	})

	DirectToMainPage(w, r)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Minute),
	})
	tmpl.Tmpl.ExecuteTemplate(w, "Login", nil)
}
