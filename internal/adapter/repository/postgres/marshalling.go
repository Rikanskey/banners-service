package postgres

import (
	"banners-service/internal/app"
	"time"
)

type banner struct {
	Id       int            `sql:"id"`
	Content  map[string]any `sql:"content"`
	Created  time.Time      `sql:"created_at"`
	Updated  time.Time      `sql:"updated_at"`
	isActive bool           `sql:"is_active"`
	Feature  feature        `sql:"feature"`
	Tags     []tag          `sql:"tags"`
}

type tag struct {
	Id int `sql:"id"`
}

type feature struct {
	Id int `sql:"id"`
}

func marshallBanner(b app.Banner) banner {
	return banner{
		Id:      b.Id,
		Content: b.Content,
		Created: b.Created,
		Updated: b.Updated,
		Feature: marshallFeature(b.Feature),
		Tags:    marshallTags(b.Tags),
	}
}

func marshallFeature(f app.Feature) feature {
	return feature{Id: f.Id}
}

func marshallTags(t []app.Tag) []tag {
	a := make([]tag, len(t))
	for i, el := range t {
		a[i] = marshallTag(el)
	}

	return a
}

func marshallTag(t app.Tag) tag {
	return tag{Id: t.Id}
}
