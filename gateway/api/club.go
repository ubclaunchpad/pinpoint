package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"go.uber.org/zap"
)

// ClubRouter routes to all club endpoints
type ClubRouter struct {
	l   *zap.SugaredLogger
	c   pinpoint.CoreClient
	mux *chi.Mux
}

func newClubRouter(l *zap.SugaredLogger, c pinpoint.CoreClient) *ClubRouter {
	router := chi.NewRouter()
	club := &ClubRouter{l, c, router}

	// club-related endpoints
	router.Post("create", club.createClub)

	// club-event-related endpoints
	router.Mount("event", router.Group(func(r chi.Router) {
		r.Post("create", club.createEvent)
	}))

	// club-period-related endpoints
	router.Mount("period", router.Group(func(r chi.Router) {
		r.Post("create", club.createPeriod)
	}))

	return &ClubRouter{l.Named("clubs"), c, router}
}

func (club *ClubRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	club.mux.ServeHTTP(w, r)
}

func (club *ClubRouter) createEvent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// parse request data
	var eData schema.CreateEvent

	if eData.Fields == nil {
		render.Render(w, r, res.ErrBadRequest(r, errors.New("Missing fields"), "Missing fields"))
	}

	if err := decoder.Decode(&eData); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}

	// TODO: create event in core

	render.Render(w, r, res.Message(r, "Event created successfully", http.StatusCreated))
}

func (club *ClubRouter) createClub(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// parse request data
	var eData schema.CreateClub

	if err := decoder.Decode(&eData); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid request"))
		return
	}

	// TODO: create club in core

	render.Render(w, r, res.Message(r, "Club created successfully", http.StatusCreated))
}

func (club *ClubRouter) createPeriod(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	// parse request data
	var eData schema.CreatePeriod

	if err := decoder.Decode(&eData); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}
	// parse form date input into standard format
	layout := "2006-01-02"
	start, err := time.Parse(layout, eData.Start)
	end, err := time.Parse(layout, eData.End)
	if err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}

	// TODO: create period in core, for now just log
	fmt.Println(start, end)

	render.Render(w, r, res.Message(r, "Period created sucessfully", http.StatusCreated))
}
