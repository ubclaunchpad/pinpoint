package api

import (
	"testing"

	"github.com/ubclaunchpad/pinpoint/utils"
)

func TestNew(t *testing.T) {
	l, err := utils.NewLogger(true)
	if err != nil {
		t.Error(err)
	}
	_, _ = New(nil, l, true)
}
