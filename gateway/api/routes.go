package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
)

func (a *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := a.c.GetStatus(r.Context(), &request.Status{})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)
}
