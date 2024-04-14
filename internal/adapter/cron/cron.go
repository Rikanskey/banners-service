package cron

import (
	"banners-service/internal/adapter/cache"
	"banners-service/internal/adapter/repository/postgres"
	"github.com/robfig/cron/v3"
)

type DbDownloader struct {
	bannerRepository *postgres.BannerRepository
	cache            *cache.BannerCache
}

func InitCron(bannerRepository *postgres.BannerRepository, cache *cache.BannerCache) {
	a := cron.Cron{}
	b := DbDownloader{bannerRepository: bannerRepository, cache: cache}
	a.AddJob("*/5 * * * *", &b)

	a.Start()
}

func (db *DbDownloader) fromDbToCache() {
	banners, _ := db.bannerRepository.GetAll()

	for _, banner := range banners {
		db.cache.SetUserBanner(banner)
	}
}

func (db *DbDownloader) Run() {
	db.fromDbToCache()
}
