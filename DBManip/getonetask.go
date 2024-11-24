package DBManip

import (
	"VysMax/models"
	"database/sql"
)

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
