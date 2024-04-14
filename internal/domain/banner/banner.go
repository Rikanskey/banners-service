package banner

import "time"

type Banner struct {
	id        int
	content   map[string]any
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
	feature   Feature
	tags      []Tag
}
