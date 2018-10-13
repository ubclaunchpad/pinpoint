package api

import (
	"testing"
	"time"

	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestAPI_New(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
	}
	_, _ = New(l)
}

func TestAPI_Run(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	api, err := New(l)
	if err != nil {
		t.Error(err)
		return
	}

	// stub run
	go api.Run("localhost", "8080", RunOpts{})
	time.Sleep(time.Millisecond)
}
