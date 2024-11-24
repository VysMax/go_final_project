package DBManip

import (
	"database/sql"
	"fmt"
)

func (r *Repository) DeleteRow(id string) error {

	res, err := r.Db.Exec("DELETE FROM scheduler WHERE id = :id",
		sql.Named("id", id))
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
