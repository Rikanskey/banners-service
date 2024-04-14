package query

import (
	"banners-service/internal/adapter/cache"
	"banners-service/internal/app"
	"context"
	"github.com/pkg/errors"
)

type getUserBannerModel interface {
	GetUserBanner(ctx context.Context, tagId, featureId int) (app.Banner, error)
}

type GetUserBannerHandler struct {
	getModel getUserBannerModel
	cache    *cache.BannerCache
}

func NewGetUserBannerHandler(model getUserBannerModel, cache *cache.BannerCache) GetUserBannerHandler {
	if model == nil {
		panic("nil getBannerByTagAndFeatureHandler")
	}

	return GetUserBannerHandler{getModel: model, cache: cache}
}

func (h GetUserBannerHandler) Handle(ctx context.Context, query app.GetBannerQuery, actual bool) (app.Banner, error) {
	var banner app.Banner
	var err error

	if actual {
		banner, err = h.getModel.GetUserBanner(ctx, query.TagId, query.FeatureId)
	} else {
		banner, err = h.cache.GetUserBanner(query.TagId, query.FeatureId)
	}

	return banner, errors.Wrapf(err, "banner by tag id and feature id %d", query.FeatureId)
}
