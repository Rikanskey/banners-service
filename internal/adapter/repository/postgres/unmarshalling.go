package postgres

import "banners-service/internal/app"

func unmarshallBanner(b banner) app.Banner {
	return app.Banner{
		Id:      b.Id,
		Feature: unmarshallFeature(b.Feature),
		Content: b.Content,
		Created: b.Created,
		Updated: b.Updated,
		Tags:    unmarshallTags(b.Tags),
	}
}

func unmarshallFeature(f feature) app.Feature {
	return app.Feature{
		Id: f.Id,
	}
}

func unmarshallTags(t []tag) []app.Tag {
	tags := make([]app.Tag, len(t))
	for i, t := range t {
		tags[i] = unmarshallTag(t)
	}
	return tags
}

func unmarshallTag(tag tag) app.Tag {
	return app.Tag{
		Id: tag.Id,
	}
}
