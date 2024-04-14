package postgres

import (
	"banners-service/internal/app"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

type BannerRepository struct {
	banners *sql.DB
}

func NewBannerRepository(db *sql.DB) *BannerRepository {
	return &BannerRepository{banners: db}
}

func (b *BannerRepository) GetAll() ([]app.Banner, error) {
	rows, err := b.banners.Prepare("SELECT * FROM banner")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	var banners []app.Banner

	result, err := rows.Query()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer result.Close()
	result.Next()

	for result.Next() {
		var dbBan banner
		err = result.Scan(&dbBan)
		banners = append(banners, unmarshallBanner(dbBan))

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

	}

	return banners, nil
}

func (b *BannerRepository) GetUserBanner(ctx context.Context, tagId, featureId int) (app.Banner, error) {
	stmt, err := b.banners.Prepare("SELECT * FROM banner JOIN public.banner_tag bt on banner.id = bt.banner_id WHERE tag_id=$1 AND feature_id=$2 AND is_active=true")
	defer stmt.Close()

	if err != nil {
		logrus.WithError(err).Println("Error preparing statement")
		return app.Banner{}, err
	}

	result, err := stmt.QueryContext(ctx, tagId, featureId)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return app.Banner{}, err
	}
	defer result.Close()
	result.Next()

	dbBanner := banner{}
	err = result.Scan(&dbBanner)
	if err != nil {
		logrus.WithError(err).Println("Error scanning row")
		return app.Banner{}, err
	}

	return unmarshallBanner(dbBanner), err
}

func (b *BannerRepository) GetAdminBanner(ctx context.Context, tagId, featureId int) (app.Banner, error) {
	stmt, err := b.banners.Prepare("SELECT * FROM banner JOIN public.banner_tag bt on banner.id = bt.banner_id WHERE tag_id=$1 AND feature_id=$2")
	defer stmt.Close()

	if err != nil {
		logrus.WithError(err).Println("Error preparing statement")
		return app.Banner{}, err
	}

	result, err := stmt.QueryContext(ctx, tagId, featureId)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return app.Banner{}, err
	}
	defer result.Close()
	result.Next()

	dbBanner := banner{}
	err = result.Scan(&dbBanner)
	if err != nil {
		logrus.WithError(err).Println("Error scanning row")
		return app.Banner{}, err
	}

	return unmarshallBanner(dbBanner), nil
}

func (b *BannerRepository) CreateBanner(ctx context.Context, bannerApp app.Banner) (int, error) {
	stmt, err := b.banners.Prepare("INSERT INTO banner (feature_id, content, created, updated, is_active) VALUES ($1, $2, $3, $4, $5) returning id")
	if err != nil {
		logrus.WithError(err).Println("Error preparing statement")
		return 0, err
	}
	defer stmt.Close()

	dbBanner := marshallBanner(bannerApp)

	result, err := stmt.QueryContext(ctx, dbBanner.Feature.Id, dbBanner.Content, time.Now(), time.Now(), true)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return 0, err
	}
	defer result.Close()
	result.Next()

	var id int
	err = result.Scan(&id)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return 0, err
	}

	return id, nil
}

func (b *BannerRepository) RemoveBanner(ctx context.Context, bannerId int) error {
	stmt, err := b.banners.Prepare("DELETE FROM banner WHERE id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err).Println("Error preparing statement")
		return err
	}

	_, err = stmt.ExecContext(ctx, bannerId)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return err
	}

	return nil
}

func (b *BannerRepository) UpdateBanner(ctx context.Context, banner app.Banner) error {
	stmt, err := b.banners.Prepare("UPDATE banner SET content=$1, feature_id=$2, updated=$3, is_active=$4 WHERE id=$5")
	if err != nil {
		logrus.WithError(err).Println("Error preparing statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, banner.Content, banner.Feature, time.Now(), banner.IsActive, banner.Id)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return err
	}
	err = b.updateBannerTags(ctx, banner.Id, banner.Tags)
	if err != nil {
		logrus.WithError(err).Println("Error executing statement")
		return err
	}

	return nil
}

func (b *BannerRepository) updateBannerTags(ctx context.Context, bannerId int, tags []app.Tag) error {
	stmt, err := b.banners.Prepare("INSERT INTO banner_tag (banner_id, tag_id)  VALUES ($1, $2) ")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, tag := range tags {
		_, err = stmt.ExecContext(ctx, bannerId, tag.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
