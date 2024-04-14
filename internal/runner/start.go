package runner

import (
	"banners-service/internal/adapter/cache"
	"banners-service/internal/adapter/cron"
	"banners-service/internal/adapter/repository/postgres"
	"banners-service/internal/app"
	"banners-service/internal/app/command"
	"banners-service/internal/app/query"
	"banners-service/internal/config"
	"banners-service/internal/port/http"
	"banners-service/internal/server"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

func Start(configDir string) {
	cfg := newConfig(configDir)
	db := initDB(cfg)
	bannerCache := cache.InitCache()
	application := newApplication(db, bannerCache)
	startServer(cfg, application)
}

func newConfig(configDir string) *config.Config {
	cfg, err := config.New(configDir)
	if err != nil {
		log.Panicln(err)
	}

	return cfg
}

func initDB(cfg *config.Config) *sql.DB {
	dbInfo := fmt.Sprintf("postgresql://%s:%s@postgres/%s?sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Name)
	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		fmt.Println("DB err")
	}
	err = db.Ping()
	if err != nil {
		logrus.WithError(err).Println("No connection")
	}
	defer db.Close()
	return db
}

func newApplication(db *sql.DB, cache *cache.BannerCache) app.Application {
	bannersRepository := postgres.NewBannerRepository(db)

	cron.InitCron(bannersRepository, cache)

	return app.Application{
		Commands: app.Commands{
			CreateBanner: command.NewCreateBanner(bannersRepository),
			RemoveBanner: command.NewRemoveBanner(bannersRepository),
			UpdateBanner: command.NewUpdateBanner(bannersRepository),
		},
		Queries: app.Queries{
			GetUserBannerQuery:  query.NewGetUserBannerHandler(bannersRepository, cache),
			GetAdminBannerQuery: query.NewGetAdminBannerHandler(bannersRepository),
		},
	}
}

func startServer(cfg *config.Config, application app.Application) {
	logrus.Info(fmt.Sprintf("Starting HTTP server on address: %s", cfg.HTTP.Port))

	httpServer := server.New(cfg, http.NewHandler(application))

	err := httpServer.Run()

	logrus.WithError(err).Fatal("HTTP server stopped")
}
