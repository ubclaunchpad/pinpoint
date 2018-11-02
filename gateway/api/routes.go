package api

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
)

func (a *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := a.c.GetStatus(context.Background(), &request.Status{Callback: "hello world"})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)
}

func (a *API) createAccountHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	resp, err := a.c.CreateAccount(context.Background(), &request.CreateAccount{Email: email})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)
}

func (a *API) verifyHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")
	resp, err := a.c.Verify(context.Background(), &request.Verify{Hash: hash})
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(r, err))
		return
	}
	render.JSON(w, r, resp)
}
