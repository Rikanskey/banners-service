package app

import "banners-service/internal/domain/banner"

type (
	CreateBannerCommand struct {
		Content map[string]any
		Feature banner.Feature
		Tags    []banner.Tag
	}
	UpdateBannerCommand struct {
		Content map[string]any
		Feature banner.Feature
		Tags    []banner.Tag
	}
	RemoveBannerCommand struct {
		BannerId int
	}
)
