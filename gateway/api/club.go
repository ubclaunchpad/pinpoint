package api

import (
	"encoding/json"
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

func newClubRouter(l *zap.SugaredLogger, core pinpoint.CoreClient) *ClubRouter {
	c := &ClubRouter{l.Named("clubs"), core, chi.NewRouter()}

	// club-related endpoints
	c.mux.Post("/create", c.createClub)

	// club-event-related endpoints
	c.mux.Route("/event", func(r chi.Router) {
		r.Post("/create", c.createEvent)
	})

	// club-period-related endpoints
	c.mux.Route("/period", func(r chi.Router) {
		r.Post("/create", c.createPeriod)
	})

	return c
}

func (club *ClubRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	club.mux.ServeHTTP(w, r)
}

func (club *ClubRouter) createEvent(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var data schema.CreateEvent
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}

	// TODO: create event in core

	render.Render(w, r, res.Message(r, "Event created successfully", http.StatusCreated))
}

func (club *ClubRouter) createClub(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var data schema.CreateClub
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid request"))
		return
	}

	// TODO: create club in core

	render.Render(w, r, res.Message(r, "Club created successfully", http.StatusCreated))
}

func (club *ClubRouter) createPeriod(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var data schema.CreatePeriod
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}

	// parse form date input into standard format
	layout := "2006-01-02"
	start, err := time.Parse(layout, data.Start)
	end, err := time.Parse(layout, data.End)
	if err != nil {
		render.Render(w, r, res.ErrBadRequest(r, err, "Invalid input"))
		return
	}

	// TODO: create period in core, for now just log
	fmt.Println(start, end)

	render.Render(w, r, res.Message(r, "Period created sucessfully", http.StatusCreated))
}
