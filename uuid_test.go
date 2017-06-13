package main

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
)

func TestItReturnsTheCorrectVersion(t *testing.T) {
	cases := []int{1, 4}
	for _, v := range cases {
		app := NewApp()
		s, err := app.Generate(v)
		if err != nil {
			t.Errorf("\nhave err: %s\nwant: uuid version: %d", err.Error(), v)
		}
		if u, err := uuid.FromString(s); err != nil || u.Version() != uint(v) {
			t.Errorf("generated uuid: %s\nfailed test for version: %d", s, v)
		}
	}
}

func TestItReturnsAnUnsupportedError(t *testing.T) {
	msgFmt := "unsupported uuid version: %d"
	cases := []int{2, 3, 5}
	for _, v := range cases {
		app := NewApp()
		_, err := app.Generate(v)
		if err == nil {
			t.Errorf("didn't receive error on version: %d", v)
		} else if have, want := err.Error(), fmt.Sprintf(msgFmt, v); have != want {
			t.Errorf("\nhave: %s\nwant: %s", have, want)
		}
	}
}

func TestItReturnsAnUnknownVersionError(t *testing.T) {
	msgFmt := "unknown uuid version: %d"
	cases := []int{-100, -10, -1, 0, 6, 10, 100}
	for _, v := range cases {
		app := NewApp()
		_, err := app.Generate(v)
		if err == nil {
			t.Errorf("didn't receive error on version: %d", v)
		} else if have, want := err.Error(), fmt.Sprintf(msgFmt, v); have != want {
			t.Errorf("\nhave: %s\nwant: %s", have, want)
		}
	}
}
