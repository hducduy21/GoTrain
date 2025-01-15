package handlers

import (
	"context"
	"emb/pkg/db"
	"emb/pkg/db/ent"
	"emb/pkg/db/ent/task"
	"emb/pkg/db/ent/user"
	"emb/pkg/tmpl"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TaskTmpl struct {
	Task *ent.Task
}

func CreateTaskHandle(w http.ResponseWriter, r *http.Request) {
	taskName := r.FormValue("newtask")
	if taskName == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	task, err := db.Client.Task.Create().SetName(taskName).SetDone(false).Save(context.Background())
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "Task", task)
}

func UpdateTaskHandle(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, errConvrt := strconv.Atoi(taskIDStr)
	if errConvrt != nil {
		http.Error(w, "Task ID cannot be empty", http.StatusBadRequest)
		return
	}

	taskName := r.URL.Query().Get("name")
	if taskName == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	db.Client.Task.UpdateOneID(taskID).SetName(taskName).ExecX(context.Background())
	w.WriteHeader(http.StatusOK)
}

func UpdateTaskStatusHandle(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, errConvrt := strconv.Atoi(taskIDStr)
	if errConvrt != nil {
		http.Error(w, "Task ID cannot be empty", http.StatusBadRequest)
		return
	}

	task, _ := db.Client.Task.Query().Where(task.ID(taskID)).Only(context.Background())
	db.Client.Task.UpdateOneID(taskID).SetDone(!task.Done).ExecX(context.Background())
	w.WriteHeader(http.StatusOK)
}

func DeleteTaskHandle(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, errConvrt := strconv.Atoi(taskIDStr)

	if errConvrt != nil {
		http.Error(w, "Task ID cannot be empty", http.StatusBadRequest)
		return
	}
	db.Client.Task.DeleteOneID(taskID).ExecX(context.Background())
	w.WriteHeader(http.StatusOK)
}

func JoinTaskHandle(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, errConvrt := strconv.Atoi(taskIDStr)
	if errConvrt != nil {
		http.Error(w, "Task ID cannot be empty", http.StatusBadRequest)
		return
	}
	username, _ := r.Context().Value("username").(string)

	user, err := db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
	db.Client.User.UpdateOne(user).AddTaskIDs(taskID).ExecX(context.Background())
	if err != nil {
		http.Error(w, "Failed to join task", http.StatusInternalServerError)
		return
	}

	task, err := db.Client.Task.Query().Where(task.ID(taskID)).WithUsers().Only(context.Background())
	if err != nil {
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}
	tmpl.Tmpl.ExecuteTemplate(w, "Task", task)
}

func LeaveTaskHandle(w http.ResponseWriter, r *http.Request) {
	taskIDStr := chi.URLParam(r, "id")
	taskID, errConvrt := strconv.Atoi(taskIDStr)
	if errConvrt != nil {
		http.Error(w, "Task ID cannot be empty", http.StatusBadRequest)
		return
	}
	username, _ := r.Context().Value("username").(string)

	user, err := db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
	db.Client.User.UpdateOne(user).RemoveTaskIDs(taskID).ExecX(context.Background())
	if err != nil {
		http.Error(w, "Failed to join task", http.StatusInternalServerError)
		return
	}
	task, err := db.Client.Task.Query().Where(task.ID(taskID)).WithUsers().Only(context.Background())
	if err != nil {
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}
	tmpl.Tmpl.ExecuteTemplate(w, "Task", task)
}
