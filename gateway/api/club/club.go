package club

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/api/ctxutil"
	"github.com/ubclaunchpad/pinpoint/gateway/res"
	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/models"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap"
)

const (
	keyClub   ctxutil.Key = "club_name"
	keyPeriod ctxutil.Key = "period_name"
)

// Router routes to all club endpoints
type Router struct {
	l   *zap.SugaredLogger
	c   pinpoint.CoreClient
	mux *chi.Mux
}

// NewClubRouter instantiates a new router for club functionality
func NewClubRouter(l *zap.SugaredLogger, core pinpoint.CoreClient) *Router {
	c := &Router{l.Named("clubs"), core, chi.NewRouter()}

	// club endpoints
	c.mux.Post("/create", c.createClub)

	// specific club endpoints, e.g. /my_awesome_club/overview
	c.mux.Route(fmt.Sprintf("/{%s}", keyClub), func(r chi.Router) {
		r.Get("/overview", func(w http.ResponseWriter, r *http.Request) { /* TODO */ })

		// club.period endpoints
		r.Route("/period", func(r chi.Router) {
			r.Post("/create", c.createPeriod)

			// specific period endpoints, e.g. /my_awesome_club/period/spring_2018
			r.Route(fmt.Sprintf("/{%s}", keyPeriod), func(r chi.Router) {
				// club.period.event endpoints
				r.Route("/event", func(r chi.Router) {
					r.Post("/create", c.createEvent)
				})
			})
		})
	})

	return c
}

func (c *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mux.ServeHTTP(w, r)
}

func (c *Router) createEvent(w http.ResponseWriter, r *http.Request) {
	// get associated period and club
	var club = chi.URLParam(r, string(keyClub))
	var period = chi.URLParam(r, string(keyPeriod))
	c.l.Debugw("received request to create event",
		"club", club,
		"period", period)

	// read request body
	var decoder = json.NewDecoder(r.Body)
	var data schema.CreateEvent
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}

	var fields = make([]*models.FieldProps, len(data.Fields))
	for i, f := range data.Fields {
		var pbf = &models.FieldProps{}
		switch f.Type {
		case schema.FieldTypeLongText:
			var lt = &models.FieldProps_LongText{}
			if err := json.Unmarshal(f.Properties, lt); err != nil {
				render.Render(w, r, res.ErrBadRequest("invalid data for event field properties",
					"field", f.Type, "error", err))
				return
			}
			pbf.Properties = lt
		case schema.FieldTypeShortText:
			var st = &models.FieldProps_ShortText{}
			if err := json.Unmarshal(f.Properties, st); err != nil {
				render.Render(w, r, res.ErrBadRequest("invalid data for event field properties",
					"field", f.Type, "error", err))
				return
			}
			pbf.Properties = st
		default:
			render.Render(w, r, res.ErrBadRequest("invalid type for event field properties",
				"field", f.Type))
			return
		}
		fields[i] = pbf
	}

	// TODO: create event in core
	c.c.CreateEvent(context.Background(), &request.CreateEvent{
		Event: &models.EventProps{
			Period:      period,
			EventID:     "", // TODO what the?
			Name:        data.Name,
			Club:        club,
			Description: "", //TODO
			Fields:      fields,
		},
	})

	render.Render(w, r, res.Msg("Event created successfully", http.StatusCreated))
}

func (c *Router) createClub(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var data schema.CreateClub
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}

	// TODO: create club in core

	render.Render(w, r, res.Msg("Club created successfully", http.StatusCreated))
}

func (c *Router) createPeriod(w http.ResponseWriter, r *http.Request) {
	var (
		decoder    = json.NewDecoder(r.Body)
		data       schema.CreatePeriod
		err        error
		start, end time.Time
	)
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}

	// parse form date input into standard format
	const layout = "2006-01-02"
	if start, err = time.Parse(layout, data.Start); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid start date"))
		return
	}
	if end, err = time.Parse(layout, data.End); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid end date"))
		return
	}

	// check validity
	if start.After(end) {
		render.Render(w, r, res.ErrBadRequest("start date must be before end date"))
		return
	}

	// TODO: create period in core, for now just log
	fmt.Println(start, end)

	render.Render(w, r, res.Msg("period created sucessfully", http.StatusCreated))
}
