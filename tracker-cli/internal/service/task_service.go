package service

import (
	"errors"
	"log"
	"tracker-cli/internal/model"
	"tracker-cli/internal/storage"
)

type TaskService struct {
	taskStorage *storage.TaskStorage
}

func NewTaskService(taskStorage *storage.TaskStorage) *TaskService {
	return &TaskService{taskStorage: taskStorage}
}

func (ts *TaskService) NewDB() {
	ts.taskStorage.NewSQliteDB()
}

func (ts *TaskService) Migrate() {
	if err := ts.taskStorage.DB().AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}

func (ts *TaskService) CreateTask(name, description string) error {

	var task model.Task

	if name == "" {
		return errors.New("task name is required")
	}

	task.Name = name

	if description != "" {
		task.Description = description
	}

	return ts.taskStorage.DB().Create(&task).Error
}

func (ts *TaskService) UpdateTask(id uint, name, description string) error {

	var task model.Task

	if err := ts.taskStorage.DB().First(&task, id).Error; err != nil {
		return err
	}

	if name != "" {
		task.Name = name
	}

	if description != "" {
		task.Description = description
	}

	return ts.taskStorage.DB().Save(&task).Error
}

func (ts *TaskService) UpdateTaskByID(id uint, status string) error {

	var task model.Task

	if err := ts.taskStorage.DB().First(&task, id).Error; err != nil {
		return err
	}

	taskStatus := model.TaskStatus(status)
	if taskStatus != model.Todo && taskStatus != model.InProgress && taskStatus != model.Done {
		return errors.New("invalid status")
	}

	return ts.taskStorage.DB().Model(&task).Update("status", taskStatus).Error

}

func (ts *TaskService) UpdateTaskByName(name, status string) error {

	var task model.Task

	if err := ts.taskStorage.DB().Where("name = ?", name).First(&task).Error; err != nil {
		return err
	}

	taskStatus := model.TaskStatus(status)
	if taskStatus != model.Todo && taskStatus != model.InProgress && taskStatus != model.Done {
		return errors.New("invalid status")
	}

	return ts.taskStorage.DB().Model(&task).Update("status", taskStatus).Error

}

func (ts *TaskService) DeleteTaskByID(id uint) error {

	var task model.Task

	if err := ts.taskStorage.DB().First(&task, id).Error; err != nil {
		return err
	}

	return ts.taskStorage.DB().Delete(&task).Error
}

func (ts *TaskService) DeleteTaskByName(name string) error {

	var task model.Task

	if err := ts.taskStorage.DB().Where("name = ?", name).First(&task).Error; err != nil {
		return err
	}

	return ts.taskStorage.DB().Delete(&task).Error
}

func (ts *TaskService) GetTasks() ([]model.Task, error) {

	var tasks []model.Task

	if err := ts.taskStorage.DB().Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskService) GetTaskByID(id uint) (model.Task, error) {

	var task model.Task

	if err := ts.taskStorage.DB().First(&task, id).Error; err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) GetTaskByName(name string) (model.Task, error) {

	var task model.Task

	if err := ts.taskStorage.DB().Where("name = ?", name).First(&task).Error; err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (ts *TaskService) GetTasksByStatus(status string) ([]model.Task, error) {
	var tasks []model.Task

	taskStatus := model.TaskStatus(status)
	if taskStatus != model.Todo && taskStatus != model.InProgress && taskStatus != model.Done {
		return nil, errors.New("invalid status")
	}

	if err := ts.taskStorage.DB().Where("status = ?", status).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
