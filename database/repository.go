package database

import (
	"database/sql"
	"fmt"
	"time"

	"VysMax/internalfunc"
	"VysMax/models"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Db: db,
	}
}

func (r *Repository) GetSingle(id string) (models.Task, error) {
	var task models.Task

	task.ID = id

	row := r.Db.QueryRow("SELECT id, date, title, comment, repeat FROM scheduler WHERE id = :id",
		sql.Named("id", id))

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)

	if err != nil {
		return task, err
	}
	return task, nil
}

func (r *Repository) GetMultiple(search string) models.TasksList {
	var (
		result models.TasksList
		tasks  []models.Task
		rows   *sql.Rows
		err    error
	)

	parsedSearch, err := time.Parse(internalfunc.SearchLayout, search)

	switch {
	case search == "" && err != nil:
		rows, err = r.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit",
			sql.Named("limit", 50))
		if err != nil {
			result.Error = "Не удалось выполнить запрос к базе данных"
			return result
		}
	case search != "" && err == nil:
		date := parsedSearch.Format(internalfunc.Layout)
		rows, err = r.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date = :date ORDER BY date LIMIT :limit",
			sql.Named("date", date),
			sql.Named("limit", 50))
		if err != nil {
			result.Error = "Не удалось выполнить запрос к базе данных"
			return result
		}
	default:
		search = "%" + search + "%"
		rows, err = r.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE :search OR comment LIKE :search ORDER BY date LIMIT :limit",
			sql.Named("search", search),
			sql.Named("limit", 50))
		if err != nil {
			result.Error = "Не удалось выполнить запрос к базе данных"
			return result
		}
	}

	for rows.Next() {
		task := models.Task{}

		err = rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			result.Error = "Не удалось выполнить запрос к базе данных"
			return result
		}
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		result.Error = "Не удалось выполнить запрос к базе данных"
		return result
	}

	result.Tasks = tasks

	return result
}

func (r *Repository) AddRow(newTask models.Task) (int64, error) {
	res, err := r.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", newTask.Date),
		sql.Named("title", newTask.Title),
		sql.Named("comment", newTask.Comment),
		sql.Named("repeat", newTask.Repeat))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) UpdateTask(updatedTask models.Task) error {

	res, err := r.Db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", updatedTask.Date),
		sql.Named("title", updatedTask.Title),
		sql.Named("comment", updatedTask.Comment),
		sql.Named("repeat", updatedTask.Repeat),
		sql.Named("id", updatedTask.ID))

	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if numRows == 0 {
		err = fmt.Errorf("task not found")
		return err
	}

	return nil
}

func (r *Repository) DeleteRow(id string) error {

	res, err := r.Db.Exec("DELETE FROM scheduler WHERE id = :id",
		sql.Named("id", id))
	if err != nil {
		return err
	}

	numRows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if numRows == 0 {
		err = fmt.Errorf("task not found")
		return err
	}

	return nil
}
