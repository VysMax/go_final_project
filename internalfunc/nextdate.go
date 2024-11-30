package internalfunc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Layout       = "20060102"
	SearchLayout = "02.01.2006"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("repeat field is empty")
	}
	initialTime, err := time.Parse(Layout, date)
	if err != nil {
		return "", fmt.Errorf("incorrect date")
	}
	splitRepeat := strings.Split(repeat, " ")
	switch splitRepeat[0] {
	case "d":
		if len(splitRepeat) != 2 {
			return "", fmt.Errorf("incorrect format: d has one number")
		}
		daysNumber, err := strconv.Atoi(splitRepeat[1])
		if err != nil {
			return "", fmt.Errorf("incorrect format")
		}
		if daysNumber > 400 || daysNumber < 1 {
			return "", fmt.Errorf("incorrect format: number must be between 1 and 400")
		}
		newTime := initialTime.AddDate(0, 0, daysNumber)
		for newTime.Before(now) {
			newTime = newTime.AddDate(0, 0, daysNumber)
		}
		return newTime.Format(Layout), nil
	case "y":
		if len(splitRepeat) > 1 {
			return "", fmt.Errorf("incorrect format")
		}

		var newTime time.Time

		if initialTime.Year() < now.Year() {
			yearsDiff := now.Year() - initialTime.Year()
			newTime = initialTime.AddDate(yearsDiff, 0, 0)
		}

		if newTime.Before(now) {
			newTime = initialTime.AddDate(1, 0, 0)
		}

		return newTime.Format(Layout), nil
	case "w":
		if len(splitRepeat) != 2 {
			return "", fmt.Errorf("incorrect format")
		}

		if initialTime.Before(now) {
			initialTime = now
		}

		acceptedNumbers := make(map[int]struct{})

		repeatNumbers := strings.Split(splitRepeat[1], ",")
		for _, v := range repeatNumbers {
			if v == "7" {
				v = "0"
			}

			repeatNumber, err := strconv.Atoi(v)
			if err != nil {
				return "", fmt.Errorf("incorrect format")
			}

			if repeatNumber < 0 || repeatNumber > 6 {
				return "", fmt.Errorf("incorrect format: number must between 1 and 7")
			}

			_, ok := acceptedNumbers[repeatNumber]
			if ok {
				return "", fmt.Errorf("incorrect format: a number for weekday can be used only once")
			}

			acceptedNumbers[repeatNumber] = struct{}{}
		}

		newTime := initialTime

		for i := 1; i <= 7; i++ {
			newTime = initialTime.AddDate(0, 0, i)
			_, ok := acceptedNumbers[int(newTime.Weekday())]
			if ok {
				return newTime.Format(Layout), nil
			}
		}
		return "", fmt.Errorf("unknown error")

	case "m":
		if len(splitRepeat) != 2 && len(splitRepeat) != 3 {
			return "", fmt.Errorf("incorrect format: after m query must contain numbers for days (obligatory) and months (optionally)")
		}

		if initialTime.Before(now) {
			initialTime = now
		}

		acceptedDays := make(map[int]struct{})
		acceptedMonths := make(map[int]struct{})

		repeatDays := strings.Split(splitRepeat[1], ",")
		for _, v := range repeatDays {
			repeatDay, err := strconv.Atoi(v)
			if err != nil {
				return "", fmt.Errorf("incorrect format")
			}

			if repeatDay < -2 || repeatDay > 31 || repeatDay == 0 {
				return "", fmt.Errorf("incorrect format: after m query must contain number for day/days between 1 and 31, or -1, or -2")
			}

			_, ok := acceptedDays[repeatDay]
			if ok {
				return "", fmt.Errorf("incorrect format: a number for day can be used only once")
			}

			acceptedDays[repeatDay] = struct{}{}
		}

		switch {
		case len(splitRepeat) == 2:
			for i := 0; i <= 2; i++ {
				acceptedMonths[int(initialTime.Month())+i] = struct{}{}
			}
		default:
			repeatMonths := strings.Split(splitRepeat[2], ",")
			for _, v := range repeatMonths {
				repeatMonth, err := strconv.Atoi(v)
				if err != nil {
					return "", fmt.Errorf("incorrect format")
				}

				if repeatMonth < 1 || repeatMonth > 12 {
					return "", fmt.Errorf("incorrect format: after m and days query may contain numbers for month/months between 1 and 12")
				}

				_, ok := acceptedMonths[repeatMonth]
				if ok {
					return "", fmt.Errorf("incorrect format: a number for month can be used only once")
				}

				acceptedMonths[repeatMonth] = struct{}{}
			}
		}

		newTime := initialTime
		var isDay, isMonth bool
		var nextMonthFirstDay time.Time

		for i := 1; i <= 366; i++ {
			newTime = initialTime.AddDate(0, 0, i)
			_, isDay = acceptedDays[newTime.Day()]
			_, isMonth = acceptedMonths[int(newTime.Month())]
			if isDay && isMonth {
				return newTime.Format(Layout), nil
			}
			if newTime.Day() >= 27 {
				for j := 1; j <= 2; j++ {
					nextMonthFirstDay = newTime.AddDate(0, 0, j)
					if nextMonthFirstDay.Month() == newTime.Month()+1 {
						_, isDay = acceptedDays[-1*j]
						if isDay && isMonth {
							return newTime.Format(Layout), nil
						}
					}

				}
			}
		}
		return "", fmt.Errorf("unknown error")
	default:
		return "", fmt.Errorf("incorrect format")
	}
}
