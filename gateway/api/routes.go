package api

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/grpc/request"
)

func (a *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := a.c.GetStatus(context.Background(), &request.Status{Callback: "hello world"})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)
}
