package main

import "time"

const format = "20060102150405" // YYYYMMDDhhmmss

func parseIsosec(id string) (time.Time, error) {
	t, err := time.Parse(format, id)
	return t, err
}
