package goutils_test

import (
	"errors"
	"testing"

	goutils "github.com/onichandame/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestAssert(t *testing.T) {
	t.Run("should not panics if nil", func(t *testing.T) {
		assert.NotPanics(t, func() { goutils.Assert(nil) })
	})
	t.Run("should panics if not nil", func(t *testing.T) {
		assert.Panics(t, func() { goutils.Assert(errors.New("")) })
	})
}
