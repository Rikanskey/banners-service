package command

import (
	"banners-service/internal/app"
	"context"
)

type bannerRepository interface {
	GetUserBanner(ctx context.Context, tagId, featureId int) (app.Banner, error)
	GetAdminBanner(ctx context.Context, tagId, featureId int) (app.Banner, error)
	CreateBanner(ctx context.Context, banner app.Banner) (int, error)
	UpdateBanner(ctx context.Context, banner app.Banner) error
	RemoveBanner(ctx context.Context, bannerId int) error
}
