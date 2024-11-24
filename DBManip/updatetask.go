package DBManip

import (
	"VysMax/models"
	"database/sql"
	"fmt"
)

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
	if numRows == 0 || err != nil {
		err = fmt.Errorf("task not found")
		return err
	}

	return nil
}
