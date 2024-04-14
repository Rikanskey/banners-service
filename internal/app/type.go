package app

import "time"

type (
	Banner struct {
		Id       int
		Content  map[string]any
		Created  time.Time
		Updated  time.Time
		Feature  Feature
		Tags     []Tag
		IsActive bool
	}
	Feature struct {
		Id int
	}
	Tag struct {
		Id int
	}
)
