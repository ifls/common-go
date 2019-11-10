package util

import (
	"testing"
	"time"
)

func TestRunDelay(t *testing.T) {
	RunDelay(time.Timer{
		C: nil,
	}, 10, func() {

	})
}
