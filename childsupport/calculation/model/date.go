package model

import "time"

type Date struct {
	value *time.Time
}

func NewDate(time *time.Time) *Date {
	return &Date{value: time}
}

func (d *Date) SubMonths(date *Date) int {
	year := d.value.Year() - date.value.Year()
	return year*12 + int(d.value.Month()) - int(date.value.Month())
}
