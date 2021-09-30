package goutils_test

import (
	"errors"
	"testing"

	goutils "github.com/onichandame/go-utils"
)

func TestRecoverToErr(t *testing.T) {
	throwFn := func(e interface{}) {
		panic(e)
	}
	errFn := func(e *error, err error) *error {
		defer goutils.RecoverToErr(e)
		throwFn(err)
		return e
	}
	var err error
	thrownErr := errors.New("test error")
	errFn(&err, thrownErr)
	if !errors.Is(err, thrownErr) {
		t.Error("cannot parse panics to error")
	}
}
