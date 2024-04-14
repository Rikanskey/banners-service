package cache

import (
	"banners-service/internal/app"
	lru "github.com/hashicorp/golang-lru/v2/expirable"
	"time"
)

type FeatureTag struct {
	FeatureId int
	TagId     int
}

type BannerCache struct {
	cache *lru.LRU[FeatureTag, *app.Banner]
}

func InitCache() *BannerCache {
	return &BannerCache{cache: lru.NewLRU[FeatureTag, *app.Banner](100, nil, time.Millisecond*5)}
}

func (bc *BannerCache) GetUserBanner(tagId, featureId int) (app.Banner, error) {
	banner, ok := bc.cache.Get(FeatureTag{FeatureId: featureId, TagId: tagId})
	if !ok {
		return app.Banner{}, app.ErrBannerDoesNotExist
	}
	return *banner, nil
}

func (bc *BannerCache) SetUserBanner(banner app.Banner) {
	for _, tag := range banner.Tags {
		bc.cache.Remove(FeatureTag{FeatureId: banner.Feature.Id, TagId: tag.Id})
		bc.cache.Add(FeatureTag{FeatureId: banner.Feature.Id, TagId: tag.Id}, &banner)
	}
}
