package club

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ubclaunchpad/pinpoint/gateway/schema"
	"github.com/ubclaunchpad/pinpoint/protobuf/fakes"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"go.uber.org/zap/zaptest"
)

func TestClubRouter_createClub(t *testing.T) {
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
			u := NewClubRouter(zaptest.NewLogger(t).Sugar(), fake)

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
					{Type: "", Properties: []byte(`{"julia": "has failed"}`)},
				},
			}}, http.StatusBadRequest},
		{"successfully created event", args{
			"/my_club/period/my_period/event/create",
			&schema.CreateEvent{
				Name: "Winter Semester",
				Fields: []schema.FieldProps{
					{Type: "long_text", Properties: []byte(`{"1": "2"}`)},
				},
			}}, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create club router
			fake := &fakes.FakeCoreClient{}
			u := NewClubRouter(zaptest.NewLogger(t).Sugar(), fake)

			// create request
			var b []byte
			var err error
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
	type args struct {
		path   string
		period *request.CreatePeriod
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{"bad input", args{"/my_club/period/create", nil}, http.StatusBadRequest},
		{"successfully create period", args{
			"/my_club/period/create",
			&request.CreatePeriod{
				Period: "Winter Semester",
				ClubID: "UBC Launch Pad",
			}}, http.StatusCreated},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create club router
			fake := &fakes.FakeCoreClient{}
			u := NewClubRouter(zaptest.NewLogger(t).Sugar(), fake)

			// create request
			var b []byte
			var err error
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
