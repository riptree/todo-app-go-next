package util

import (
	"time"
	"todo-app/internal/package/global"
)

func ParseDate(date *string) (*time.Time, error) {
	if date == nil {
		return nil, nil
	}

	d, err := time.Parse(global.DateFormat, *date)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
