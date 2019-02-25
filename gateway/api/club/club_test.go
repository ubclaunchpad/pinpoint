package club

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	"github.com/ubclaunchpad/pinpoint/protobuf/fakes"
	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestClubRouter_createClub(t *testing.T) {
	l, err := utils.NewLogger(true, "")
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		club *schema.CreateClub
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"bad input", args{nil}, http.StatusBadRequest},
		{"successfully create club", args{&schema.CreateClub{
			Name: "UBC Launch Pad",
			Desc: "The best software engineering club",
		}}, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fake := &fakes.FakeCoreClient{}

			// create club router
			u := NewClubRouter(l, fake)

			// create request
			var b []byte
			var err error
			if tt.args.club != nil {
				if b, err = json.Marshal(tt.args.club); err != nil {
					t.Error(err)
					return
				}
			}
			reader := bytes.NewReader(b)
			req, err := http.NewRequest("POST", "/create", reader)
			if err != nil {
				t.Error(err)
				return
			}

			// Record responses
			recorder := httptest.NewRecorder()
			u.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantCode {
				t.Errorf("expected %d, got %d", tt.wantCode, recorder.Code)
			}

			// TODO: test behaviour with fake core
		})
	}
}

func TestClubRouter_createEvent(t *testing.T) {
	// TODO once event stuff is more finalized
	l, err := utils.NewLogger(true, "")
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		path   string
		period *schema.CreateEvent
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"bad param", args{"/my_club/period/my_period/event/create", nil}, http.StatusBadRequest},
		{"invalid fields", args{
			"/my_club/period/my_period/event/create",
			&schema.CreateEvent{
				Name: "Winter Semester",
				Fields: []schema.FieldProps{
					{"", []byte(`{"julia": "has failed"}`)},
				},
			}}, http.StatusBadRequest},
		{"successfully created event", args{
			"/my_club/period/my_period/event/create",
			&schema.CreateEvent{
				Name: "Winter Semester",
				Fields: []schema.FieldProps{
					{"long_text", []byte(`{"1": "2"}`)},
				},
			}}, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create club router
			fake := &fakes.FakeCoreClient{}
			u := NewClubRouter(l, fake)

			// create request
			var b []byte
			if tt.args.period != nil {
				if b, err = json.Marshal(tt.args.period); err != nil {
					t.Error(err)
					return
				}
			}
			reader := bytes.NewReader(b)
			req, err := http.NewRequest("POST", tt.args.path, reader)
			if err != nil {
				t.Error(err)
				return
			}

			// Record responses
			recorder := httptest.NewRecorder()
			u.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantCode {
				t.Errorf("expected %d, got %d", tt.wantCode, recorder.Code)
			}

			// TODO: test behaviour with fake core
		})
	}
}

func TestClubRouter_createPeriod(t *testing.T) {
	l, err := utils.NewLogger(true, "")
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		path   string
		period *schema.CreatePeriod
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"bad input", args{"/my_club/period/create", nil}, http.StatusBadRequest},
		{"invalid start", args{
			"/my_club/period/create",
			&schema.CreatePeriod{
				Name:  "Winter Semester",
				Start: "2018asdasdfawkjefe-09",
				End:   "2018-08-09",
			}}, http.StatusBadRequest},
		{"invalid end", args{
			"/my_club/period/create",
			&schema.CreatePeriod{
				Name:  "Winter Semester",
				Start: "2018-08-09",
				End:   "2018-08asdfasdfasdf-12",
			}}, http.StatusBadRequest},
		{"end before start", args{
			"/my_club/period/create",
			&schema.CreatePeriod{
				Name:  "Winter Semester",
				Start: "2018-08-15",
				End:   "2018-08-10",
			}}, http.StatusBadRequest},
		{"successfully create period", args{
			"/my_club/period/create",
			&schema.CreatePeriod{
				Name:  "Winter Semester",
				Start: "2018-08-09",
				End:   "2018-08-10",
			}}, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create club router
			fake := &fakes.FakeCoreClient{}
			u := NewClubRouter(l, fake)

			// create request
			var b []byte
			if tt.args.period != nil {
				if b, err = json.Marshal(tt.args.period); err != nil {
					t.Error(err)
					return
				}
			}
			reader := bytes.NewReader(b)
			req, err := http.NewRequest("POST", tt.args.path, reader)
			if err != nil {
				t.Error(err)
				return
			}

			// Record responses
			recorder := httptest.NewRecorder()
			u.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantCode {
				t.Errorf("expected %d, got %d", tt.wantCode, recorder.Code)
			}

			// TODO: test behaviour with fake core
		})
	}
}
