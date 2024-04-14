package command

import (
	"banners-service/internal/app"
	"context"
)

type UpdateBannerModel interface {
	Handle(ctx context.Context, cmd app.UpdateBannerCommand) error
}

type UpdateBannerHandler struct {
	bannerRepository bannerRepository
}

func NewUpdateBanner(db bannerRepository) UpdateBannerHandler {
	if db == nil {
		panic("db is nil")
	}

	return UpdateBannerHandler{bannerRepository: db}
}

func (b UpdateBannerHandler) Handle(ctx context.Context, cmd app.UpdateBannerCommand) error {
	return nil
}
