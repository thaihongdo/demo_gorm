package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model

	Content string `json:"content"`
	UserID  string `json:"user_id"`
}

func (obj *Task) Get(id uint) (*Task, error) {
	err := db.
		Where("id = ?", id).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (obj *Task) GetByUser(userID uint) ([]*Task, error) {
	var list []*Task
	err := db.Where("user_id = ?", userID).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (obj *Task) Add() (*Task, error) {
	if err := db.Create(&obj).Error; err != nil {
		return nil, err
	}
	obj, err = obj.Get(obj.ID)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
func (obj *Task) Update(id uint) (*Task, error) {
	var tmpObj Task
	err := db.Where("id = ?", id).First(&tmpObj).Error
	if err != nil {
		return nil, err
	}
	db.Model(&tmpObj).Update(obj)
	//response
	resObj, err := tmpObj.Get(id)
	if err != nil {
		return nil, err
	}
	return resObj, nil
}
func (obj *Task) Delete(id uint) (*Task, error) {
	resObj, err := obj.Get(id)
	if err != nil {
		return nil, err
	}
	if resObj.ID > 0 {
		if err := db.Where("id = ?", resObj.ID).Delete(&resObj).Error; err != nil {
			return nil, err
		}
		return resObj, nil
	}
	return nil, errors.New("Task does not exist")
}
