package goutils_test

import (
	"reflect"
	"testing"

	goutils "github.com/onichandame/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestUnwrapType(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert.Nil(t, goutils.UnwrapType(nil))
	})
	t.Run("struct", func(t *testing.T) {
		assert.Equal(t, reflect.Struct, goutils.UnwrapType(reflect.TypeOf(struct{}{})).Kind())
	})
	t.Run("pointer", func(t *testing.T) {
		assert.Equal(t, reflect.Struct, goutils.UnwrapType(reflect.TypeOf(&struct{}{})).Kind())
	})
}

func TestUnwrapValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert.False(t, goutils.UnwrapValue(reflect.ValueOf(nil)).IsValid())
	})
	t.Run("struct", func(t *testing.T) {
		assert.Equal(t, reflect.Struct, goutils.UnwrapValue(reflect.ValueOf(struct{}{})).Kind())
	})
	t.Run("pointer", func(t *testing.T) {
		assert.Equal(t, reflect.Struct, goutils.UnwrapValue(reflect.ValueOf(&struct{}{})).Kind())
	})
}
