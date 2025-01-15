package main

import (
	"emb/pkg/db"
	"emb/pkg/handlers"
	"emb/pkg/tmpl"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db.Init()

	errTempl := tmpl.ParseTemplates()
	if errTempl != nil {
		log.Panic(errTempl)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/static"))))

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) { tmpl.Tmpl.ExecuteTemplate(w, "Register", nil) })
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) { tmpl.Tmpl.ExecuteTemplate(w, "Login", nil) })
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Delete("/logout", handlers.LogoutHandler)

	r.Get("/home", handlers.JwtAuthMiddleware(handlers.DirectToMainPage))

	// r.Post("/task", handlers.JwtAuthMiddleware(handlers.CreateTaskHandle))
	// r.Delete("/task/{id}", handlers.JwtAuthMiddleware(handlers.DeleteTaskHandle))
	// r.Patch("/task/{id}/name", handlers.JwtAuthMiddleware(handlers.UpdateTaskHandle))
	// r.Put("/task/{id}/join", handlers.JwtAuthMiddleware(handlers.HandleJoinTask))
	// r.Put("/task/{id}/leave", handlers.JwtAuthMiddleware(handlers.HandleLeaveTask))
	// r.Patch("/task/{id}/status", handlers.JwtAuthMiddleware(handlers.UpdateTaskStatusHandle))

	r.Get("/api/tasks", handlers.JwtAuthJsonMiddleware(handlers.GetTasksJsonHandle))
	r.Post("/api/tasks", handlers.JwtAuthJsonMiddleware(handlers.CreateTaskJsonHandle))
	r.Post("/api/tasks/{id}", handlers.JwtAuthJsonMiddleware(handlers.UpdateTaskJsonHandle))
	r.Patch("/api/tasks/{id}/status", handlers.JwtAuthJsonMiddleware(handlers.UpdateTaskStatusJsonHandle))
	r.Patch("/api/tasks/{id}/join", handlers.JwtAuthJsonMiddleware(handlers.JoinTaskJsonHandle))
	r.Patch("/api/tasks/{id}/leave", handlers.JwtAuthJsonMiddleware(handlers.LeaveTaskJsonHandle))
	r.Delete("/api/tasks/{id}", handlers.JwtAuthJsonMiddleware(handlers.DeleteTaskJsonHanle))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "frontend/src/index.html") })
	http.ListenAndServe(":3000", r)
}
