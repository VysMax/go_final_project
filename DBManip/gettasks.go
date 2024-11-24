package DBManip

import (
	"VysMax/models"
	"database/sql"
	"time"
)

const (
	Layout       = "20060102"
	SearchLayout = "02.01.2006"
)

func (r *Repository) GetMultiple(search string) models.TasksList {
	var (
		result models.TasksList
		tasks  []models.Task
		rows   *sql.Rows
		err    error
	)

	parsedSearch, err := time.Parse(SearchLayout, search)

	switch {
	case search == "" && err != nil:
		rows, err = r.Db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT :limit",
			sql.Named("limit", 50))
		if err != nil {
			result.Error = "Не удалось выполнить запрос к базе данных"
			return result
		}
	case search != "" && err == nil:
		date := parsedSearch.Format(Layout)
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

	result.Tasks = tasks

	return result
}
