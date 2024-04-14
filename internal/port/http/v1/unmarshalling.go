package v1

import (
	"banners-service/internal/app"
	"banners-service/pkg/httperr"
	"github.com/go-chi/render"
	"net/http"
)

func decode(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.Decode(r, v); err != nil {
		httperr.BadRequest("bad-request", err, w, r)

		return false
	}

	return true
}

func unmarshallBannerQuery(w http.ResponseWriter, r *http.Request, params GetBannerParams) (app.GetBannerQuery, bool) {
	var bannerQuery app.GetBannerQuery

	if ok := decode(w, r, &bannerQuery); !ok {
		return bannerQuery, false
	}

	return bannerQuery, true
}

func unmarshallUserBannerQuery(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) (app.GetBannerQuery, bool) {
	var bannerQuery app.GetBannerQuery

	if ok := render.Decode(r, &bannerQuery); ok != nil {
		return bannerQuery, false
	}

	return bannerQuery, true
}

func unmarshallCreateBannerCommand(w http.ResponseWriter, r *http.Request, params PostBannerParams) (app.CreateBannerCommand, bool) {
	var createBannerCommand app.CreateBannerCommand
	if err := render.Decode(r, &createBannerCommand); err != nil {
		httperr.BadRequest("bad-request", err, w, r)
		return app.CreateBannerCommand{}, false
	}

	return createBannerCommand, true
}
