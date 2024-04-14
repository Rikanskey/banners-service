package v1

import (
	"banners-service/internal/app"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

func marshalBanner(w http.ResponseWriter, r *http.Request, banner app.Banner) {
	response, _ := marshalBannerToBannerResponse(banner)

	render.Respond(w, r, response)
}

func marshalBannerToBannerResponse(banner app.Banner) ([]byte, error) {
	return json.Marshal(banner)
}
