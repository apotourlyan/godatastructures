package test

import (
	"fmt"
	"testing"

	"github.com/apotourlyan/godatastructures/internal/utilities/panics"
)

const gotWantInt = "got %d, want %d\n"
const gotWantFloat = "got %f, want %f\n"
const gotWantBool = "got %t, want %t\n"
const gotWantString = "got %q, want %q\n"
const gotWantGeneric = "got %#v, want %#v\n"

func GotWant[T comparable](t *testing.T, got T, want T) {
	t.Helper()
	if got != want {
		text := getErrorText(got, want)
		t.Error(text)
	}
}

func GotWantSlice[T comparable](t *testing.T, got []T, want []T) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("got length %d, want length %d", len(got), len(want))
	} else {
		for i := range got {
			if got[i] != want[i] {
				text := getErrorText(got[i], want[i])
				text += fmt.Sprintf("at position %d", i)
				t.Error(text)
			}
		}
	}
}

func GotWantError(t *testing.T, err error, want string) {
	t.Helper()
	if want == "" {
		return
	}

	if err == nil {
		t.Errorf("got error 'nil', want error %q", want)
	} else if got := err.Error(); got != want {
		t.Errorf("got error %q, want error %q", got, want)
	}
}

func getErrorText[T any](got T, want T) string {
	g := any(got)
	w := any(want)
	switch g.(type) {
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return fmt.Sprintf(gotWantInt, g, w)
	case float64, float32:
		return fmt.Sprintf(gotWantFloat, g, w)
	case bool:
		return fmt.Sprintf(gotWantBool, g, w)
	case string:
		return fmt.Sprintf(gotWantString, g, w)
	default:
		return fmt.Sprintf(gotWantGeneric, g, w)
	}
}

func GotWantPanic(t *testing.T, f func(), want string) {
	t.Helper()
	panicked, got := panics.CatchPanic(f)
	if !panicked {
		t.Errorf("got panic 'nil', want panic %q", want)
	} else if got != want {
		t.Errorf("got panic %q, want panic %q", got, want)
	}
}
