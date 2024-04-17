package v1

import (
	"banners-service/internal/app"
	"banners-service/internal/port/http/v1/mock"
	"banners-service/pkg/httperr"
	"github.com/pkg/errors"
	"net/http"
)

func (h handler) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {
	if params.Token == nil {
		err := errors.New("token required")
		httperr.Unauthorized("user-unauthorized", err, w, r)
		return
	} else if !mock.CheckUserToken(*params.Token) {
		err := errors.New("invalid token")
		httperr.Forbidden("user-unauthorized", err, w, r)
		return
	}

	_, ok := unmarshallBannerQuery(w, r, params)
	if !ok {
		return
	}

	var banner app.Banner
	var err error

	if err == nil {
		marshalBanner(w, r, banner)
	}

	if errors.Is(err, app.ErrBannerDoesNotExist) {
		httperr.NotFound("banner-not-found", err, w, r)

		return
	}

	httperr.InternalServerError("unexpected-error", err, w, r)
}

func (h handler) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {
	if params.Token == nil {
		err := errors.New("token required")
		httperr.Unauthorized("user-unauthorized", err, w, r)
		return
	} else if !mock.CheckAdminToken(*params.Token) {
		err := errors.New("invalid token")
		httperr.Forbidden("user-unauthorized", err, w, r)
		return
	}

	cmd, ok := unmarshallCreateBannerCommand(w, r, params)
	if !ok {
		return
	}

	_, err := h.app.Commands.CreateBanner.Handle(r.Context(), cmd)
	if err != nil {
		httperr.BadRequest("command-not-found", err, w, r)
		return
	}

	//marshalBanner(app.Banner{Id: id})

	httperr.InternalServerError("unexpected-error", err, w, r)
}

func (h handler) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {
	//TODO implement me
	panic("implement me")
}

func (h handler) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	if params.Token == nil {
		err := errors.New("token required")
		httperr.Unauthorized("user-unauthorized", err, w, r)
		return
	} else if !mock.CheckUserToken(*params.Token) {
		err := errors.New("invalid token")
		httperr.Forbidden("user-unauthorized", err, w, r)
		return
	}

	qry, ok := unmarshallUserBannerQuery(w, r, params)
	if !ok {
		return
	}

	useLastRevision := true
	if params.UseLastRevision != nil {
		useLastRevision = *params.UseLastRevision
	}

	banner, err := h.app.Queries.GetUserBannerQuery.Handle(r.Context(), qry, useLastRevision)
	if err == nil {
		marshalBanner(w, r, banner)
	}

	if errors.Is(err, app.ErrBannerDoesNotExist) {
		httperr.NotFound("banner-not-found", err, w, r)

		return
	}

	httperr.InternalServerError("unexpected-error", err, w, r)
}
