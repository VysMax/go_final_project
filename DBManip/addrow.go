package DBManip

import (
	"VysMax/models"
	"database/sql"
)

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
