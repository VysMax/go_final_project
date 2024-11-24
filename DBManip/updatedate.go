package DBManip

import (
	"VysMax/models"
	"database/sql"
	"fmt"
)

func (r *Repository) UpdateDate(task models.Task) error {

	res, err := r.Db.Exec("UPDATE scheduler SET date = :date WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("id", task.ID))
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
