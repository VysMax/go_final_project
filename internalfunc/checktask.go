package internalfunc

import (
	"time"

	"VysMax/models"
)

// const (
// 	Layout = "20060102"
// )

func CheckTaskFields(newData models.Task) (models.Task, string) {
	var errMessage string

	if newData.Title == "" {
		errMessage = "Не указан заголовок задачи"
		return newData, errMessage
	}
	if newData.Date == "" {
		newData.Date = time.Now().Format(Layout)
	}

	_, err := time.Parse(Layout, newData.Date)
	if err != nil {
		errMessage = "Дата указана в неправильном формате"
		return newData, errMessage
	}

	if newData.Date < time.Now().Format(Layout) {
		switch newData.Repeat == "" {
		case true:
			newData.Date = time.Now().Format(Layout)
		case false:
			newData.Date, err = NextDate(time.Now(), newData.Date, newData.Repeat)
			if err != nil {
				errMessage = "Правило повторения указано в неправильном формате"
			}
		}
	}

	return newData, errMessage

	// res, err := database.DB.Db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
	// 	sql.Named("date", newTask.Date),
	// 	sql.Named("title", newTask.Title),
	// 	sql.Named("comment", newTask.Comment),
	// 	sql.Named("repeat", newTask.Repeat))
	// if err != nil {
	// 	newID.Error = "Не удалось добавить запись в базу данных"
	// }

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	newID.Error = fmt.Sprintf("Не удалось найти запись под номером %d в базе данных", id)
	// 	return newID
	// }
	// newID.ID = id
	// return newID
}
