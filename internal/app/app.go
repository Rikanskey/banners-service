package app

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type (
	Commands struct {
		CreateBanner createBanner
		UpdateBanner updateBanner
		RemoveBanner removeBanner
	}
	createBanner interface {
		Handle(ctx context.Context, cmd CreateBannerCommand) (int, error)
	}
	updateBanner interface {
		Handle(ctx context.Context, cmd UpdateBannerCommand) error
	}
	removeBanner interface {
		Handle(ctx context.Context, cmd RemoveBannerCommand) error
	}
)

type (
	Queries struct {
		GetUserBannerQuery  getUserBanner
		GetAdminBannerQuery getAdminBanner
	}
	getUserBanner interface {
		Handle(ctx context.Context, query GetBannerQuery, actual bool) (banner Banner, err error)
	}
	getAdminBanner interface {
		Handle(ctx context.Context, query GetBannerQuery) (banner Banner, err error)
	}
)
