package club

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
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

	var fields = make([]*models.EventProps_FieldProps, len(data.Fields))
	for i, f := range data.Fields {
		var pbf = &models.EventProps_FieldProps{}
		switch f.Type {
		case schema.FieldTypeLongText:
			var lt = &models.EventProps_FieldProps_LongText_{}
			if err := json.Unmarshal(f.Properties, lt); err != nil {
				render.Render(w, r, res.ErrBadRequest("invalid data for event field properties",
					"field", f.Type, "error", err))
				return
			}
			pbf.Properties = lt
		case schema.FieldTypeShortText:
			var st = &models.EventProps_FieldProps_ShortText_{}
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
			EventID:     data.EventID,
			Name:        data.Name,
			Club:        club,
			Description: data.Description,
			Fields:      fields,
		},
	})

	render.Render(w, r, res.Msg("Event created successfully", http.StatusCreated))
}

func (c *Router) createClub(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var data request.CreateClub
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}
	_, claims, _ := jwtauth.FromContext(r.Context())
	data.Email = fmt.Sprintf("%v", claims["email"])

	resp, err := c.c.CreateClub(r.Context(), &data)
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(err.Error(), err))
		return
	}

	render.Render(w, r, res.Msg(resp.GetMessage(), http.StatusCreated, "clubID", data.GetClubID()))
}

func (c *Router) createPeriod(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var data request.CreatePeriod
	if err := decoder.Decode(&data); err != nil {
		render.Render(w, r, res.ErrBadRequest("invalid request"))
		return
	}
	data.ClubID = chi.URLParam(r, string(keyClub))
	resp, err := c.c.CreatePeriod(r.Context(), &data)
	if err != nil {
		render.Render(w, r, res.ErrInternalServer(err.Error(), err))
		return
	}

	render.Render(w, r, res.Msg(resp.GetMessage(), http.StatusCreated, "period", data.GetPeriod()))
}
