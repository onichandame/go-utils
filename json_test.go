package goutils_test

import (
	"testing"

	goutils "github.com/onichandame/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("can unmarshal from map", func(t *testing.T) {
		type Obj struct {
			ID       int `json:"id"`
			Children []*struct {
				ID int `json:"id"`
			} `json:"children"`
		}
		var obj Obj
		assert.NotPanics(t, func() {
			goutils.UnmarshalJSONFromMap(map[string]interface{}{"id": 1, "children": []interface{}{map[string]interface{}{"id": 2}}}, &obj)
		})
		assert.Equal(t, 1, obj.ID)
		assert.Len(t, obj.Children, 1)
		assert.Equal(t, 2, obj.Children[0].ID)
	})
}
