package goutils_test

import (
	"errors"
	"testing"

	goutils "github.com/onichandame/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	e := errors.New("test")
	t.Run("run try func", func(t *testing.T) {
		var i int
		fn := func() {
			i++
		}
		goutils.Try(fn)
		assert.Equal(t, 1, i)
	})
	t.Run("panic", func(t *testing.T) {
		fn := func() {
			panic(e)
		}
		assert.ErrorIs(t, goutils.Try(fn), e)
	})
	t.Run("not panic", func(t *testing.T) {
		fn := func() {}
		assert.Nil(t, goutils.Try(fn))
	})
}
