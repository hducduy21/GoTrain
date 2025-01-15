package handlers

import (
	"context"
	"emb/pkg/db"
	"emb/pkg/db/ent/task"
	"emb/pkg/db/ent/user"
	"emb/pkg/utils"
	"net/http"
)

type TaskRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func GetTasksJsonHandle(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Client.Task.Query().WithUsers().All(context.Background())
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	utils.JsonResponseWriter(w, tasks)
}

func CreateTaskJsonHandle(w http.ResponseWriter, r *http.Request) {
	var taskReq TaskRequest

	if err := utils.JsonParse(r, &taskReq); err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	if taskReq.Name == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	task, err := db.Client.Task.Create().SetName(taskReq.Name).Save(context.Background())
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.JsonResponseWriter(w, task)
}

func UpdateTaskJsonHandle(w http.ResponseWriter, r *http.Request) {
	id := utils.GetURLParamNumber(w, r, "id")

	var taskReq TaskRequest

	if err := utils.JsonParse(r, &taskReq); err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	if taskReq.Name == "" {
		http.Error(w, "Task name cannot be empty", http.StatusBadRequest)
		return
	}

	errUpdate := db.Client.Task.UpdateOneID(id).SetName(taskReq.Name).Exec(context.Background())
	if errUpdate != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	task := db.Client.Task.Query().Where(task.ID(id)).WithUsers().OnlyX(context.Background())
	w.WriteHeader(http.StatusOK)
	utils.JsonResponseWriter(w, task)
}

func UpdateTaskStatusJsonHandle(w http.ResponseWriter, r *http.Request) {
	id := utils.GetURLParamNumber(w, r, "id")

	var taskReq TaskRequest

	if err := utils.JsonParse(r, &taskReq); err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	db.Client.Task.UpdateOneID(id).SetDone(taskReq.Done).ExecX(context.Background())
	w.WriteHeader(http.StatusOK)
	utils.JsonResponseWriter(w, taskReq)
}

func JoinTaskJsonHandle(w http.ResponseWriter, r *http.Request) {
	id := utils.GetURLParamNumber(w, r, "id")

	username, _ := r.Context().Value("username").(string)

	user, err := db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	errUpdate := db.Client.User.UpdateOne(user).AddTaskIDs(id).Exec(context.Background())
	if errUpdate != nil {
		http.Error(w, "Failed to join task", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	utils.JsonResponseWriter(w, user)
}

func LeaveTaskJsonHandle(w http.ResponseWriter, r *http.Request) {
	id := utils.GetURLParamNumber(w, r, "id")

	username, _ := r.Context().Value("username").(string)

	user, err := db.Client.User.Query().Where(user.Username(username)).Only(context.Background())
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	errUpdate := db.Client.User.UpdateOne(user).RemoveTaskIDs(id).Exec(context.Background())
	if errUpdate != nil {
		http.Error(w, "Failed to leave task", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	utils.JsonResponseWriter(w, user)
}

func DeleteTaskJsonHanle(w http.ResponseWriter, r *http.Request) {
	id := utils.GetURLParamNumber(w, r, "id")

	errDelete := db.Client.Task.DeleteOneID(id).Exec(context.Background())
	if errDelete != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.JsonResponseWriter(w, nil)
}
