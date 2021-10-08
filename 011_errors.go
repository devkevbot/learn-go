package main

import (
	"fmt"
	"time"
)

type customError struct {
	occuredAt time.Time
	value     string
}

func (e customError) Error() string {
	return fmt.Sprintf("[%v]: %s", e.occuredAt, e.value)
}

func causeAnError() error {
	return customError{time.Now(), "Something went wrong."}
}

func _011Errors() {
	if err := causeAnError(); err != nil {
		fmt.Println(err)
	}
}
