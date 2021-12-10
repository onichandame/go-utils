package goutils_test

import (
	"errors"
	"testing"
	"time"

	goutils "github.com/onichandame/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		t.Run("passes on success", func(t *testing.T) {
			var i int
			fn := func() { i++ }
			assert.NotPanics(t, func() { goutils.Retry(fn, nil) })
			assert.EqualValues(t, 1, i)
		})
		t.Run("panics on failure", func(t *testing.T) {
			var i int
			fn := func() { i++; panic(errors.New(``)) }
			assert.Panics(t, func() { goutils.Retry(fn, nil) })
			assert.EqualValues(t, 1, i)
		})
	})
	t.Run("finite attempts + no interval", func(t *testing.T) {
		t.Run("passes on success", func(t *testing.T) {
			var i int
			fn := func() { i++ }
			assert.NotPanics(t, func() { goutils.Retry(fn, &goutils.RetryConfig{Attempts: 2}) })
			assert.Equal(t, 1, i)
		})
		t.Run("panics on failure", func(t *testing.T) {
			var i int
			fn := func() { i++; panic(errors.New(``)) }
			assert.Panics(t, func() { goutils.Retry(fn, &goutils.RetryConfig{Attempts: 2}) })
			assert.Equal(t, 2, i)
		})
	})
	t.Run("finite attempts + positive interval", func(t *testing.T) {
		t.Run("passes on success", func(t *testing.T) {
			var i int
			fn := func() { i++ }
			assert.NotPanics(t, func() { goutils.Retry(fn, &goutils.RetryConfig{Attempts: 2, Interval: time.Millisecond}) })
			assert.Equal(t, 1, i)
		})
		t.Run("panics on failure", func(t *testing.T) {
			var i int
			fn := func() { i++; panic(errors.New(``)) }
			assert.Panics(t, func() { goutils.Retry(fn, &goutils.RetryConfig{Attempts: 2, Interval: time.Millisecond}) })
			assert.Equal(t, 2, i)
		})
	})
}
