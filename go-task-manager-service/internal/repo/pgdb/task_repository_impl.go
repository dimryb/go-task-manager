package pgdb

import (
	"go-task-manager-service/internal/entity"
	"go-task-manager-service/internal/repo/pgdb/models"
	"gorm.io/gorm"
	"time"
)

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) Create(task *entity.Task) error {
	dbTask := models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
	}
	err := r.DB.Create(&dbTask).Error
	if err == nil {
		task.ID = dbTask.ID
		task.CreatedAt = dbTask.CreatedAt
		task.UpdatedAt = dbTask.UpdatedAt
	}
	return err
}

func (r *taskRepository) GetFiltered(status, priority, dueDate, title string) ([]entity.Task, error) {
	var tasks []entity.Task
	db := r.DB

	if status != "" {
		db = db.Where("status = ?", status)
	}
	if priority != "" {
		db = db.Where("priority = ?", priority)
	}
	if dueDate != "" {
		db = db.Where("due_date = ?", dueDate)
	}
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}

	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetById(id uint) (entity.Task, error) {
	var dbTask models.Task
	err := r.DB.First(&dbTask, id).Error
	if err != nil {
		return entity.Task{}, err
	}

	task := entity.Task{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		Status:      dbTask.Status,
		Priority:    dbTask.Priority,
		DueDate:     dbTask.DueDate,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
	return task, nil
}

func (r *taskRepository) Update(task entity.Task) error {
	dbTask := models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
		UpdatedAt:   time.Now(),
	}

	err := r.DB.Model(&dbTask).Where("id = ?", task.ID).Updates(dbTask).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) Delete(id uint) error {
	result := r.DB.Where("id = ?", id).Delete(&models.Task{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
