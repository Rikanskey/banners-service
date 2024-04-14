package command

import (
	"banners-service/internal/app"
	"context"
	"github.com/pkg/errors"
)

type RemoveBannerModel interface {
	RemoveBanner(ctx context.Context, cmd app.RemoveBannerCommand) error
}

type RemoveBannerHandler struct {
	bannerRepository bannerRepository
}

func NewRemoveBanner(db bannerRepository) RemoveBannerHandler {
	if db == nil {
		panic("db is nil")
	}

	return RemoveBannerHandler{bannerRepository: db}
}

func (r RemoveBannerHandler) Handle(ctx context.Context, cmd app.RemoveBannerCommand) error {
	return errors.Wrapf(r.bannerRepository.RemoveBanner(ctx, cmd.BannerId), "banner by id %d", cmd.BannerId)
}
