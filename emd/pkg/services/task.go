package services

import (
	"context"
	"emb/pkg/db"
	"emb/pkg/db/ent"
)

func GetAllTask() ([]*ent.Task, error) {
	tasks, err := db.Client.Task.Query().WithUsers().All(context.Background())
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
