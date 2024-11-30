package internalfunc

import (
	"time"

	"VysMax/models"
)

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
}
