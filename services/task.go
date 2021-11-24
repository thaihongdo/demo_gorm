package services

import (
	"errors"
	"time"
	"togo_pre/models"
)

type TaskReq struct {
	ID          uint      `json:"id"`
	Content     string    `json:"content"`
	UserID      uint      `json:"user_id"`
	CreatedTask time.Time `json:"created_task"`
}

type TaskRes struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func toTaskRes(task *models.Task) *TaskRes {
	var taskRes = &TaskRes{
		ID:      task.ID,
		Content: task.Content,
	}
	return taskRes
}

func (obj *TaskReq) Add() (bool, error) {
	now := time.Now().Format("2006-01-02")
	//check total task of user
	totalTaskOfUser, err := (&models.Task{}).GetByUser(obj.UserID, now)
	if err != nil {
		return false, err
	}
	if totalTaskOfUser > 10 {
		return false, errors.New("Can not create more task for day")
	}
	model := models.Task{
		Content:     obj.Content,
		UserID:      obj.UserID,
		CreatedTask: now,
	}
	_, err = model.Add()
	if err != nil {
		return false, err
	}
	return true, nil
}
