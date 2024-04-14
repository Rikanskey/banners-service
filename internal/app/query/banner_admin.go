package query

import (
	"banners-service/internal/app"
	"context"
	"github.com/pkg/errors"
)

type getAdminBannerModel interface {
	GetAdminBanner(ctx context.Context, tagId, featureId int) (app.Banner, error)
}

type GetAdminBannerHandler struct {
	getModel getAdminBannerModel
}

func NewGetAdminBannerHandler(model getAdminBannerModel) GetAdminBannerHandler {
	if model == nil {
		panic("nil getBannerByTagAndFeatureHandler")
	}

	return GetAdminBannerHandler{getModel: model}
}

func (h GetAdminBannerHandler) Handle(ctx context.Context, query app.GetBannerQuery) (app.Banner, error) {
	banner, err := h.getModel.GetAdminBanner(ctx, query.TagId, query.FeatureId)

	return banner, errors.Wrapf(err, "banner by tag id and feature id %d", query.FeatureId)
}
