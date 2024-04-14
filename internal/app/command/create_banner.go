package command

import (
	"banners-service/internal/app"
	"context"
)

type CreateBannerModel interface {
	CreateBanner(ctx context.Context, banner app.Banner) (int, error)
}

type CreateBannerHandler struct {
	createModel CreateBannerModel
}

func NewCreateBanner(model CreateBannerModel) CreateBannerHandler {
	if model == nil {
		panic("db is nil")
	}

	return CreateBannerHandler{createModel: model}
}

func (c CreateBannerHandler) Handle(ctx context.Context, cmd app.CreateBannerCommand) (int, error) {
	c.createModel.CreateBanner(ctx)
}
