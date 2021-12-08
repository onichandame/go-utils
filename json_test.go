package goutils_test

import (
	"testing"

	goutils "github.com/onichandame/go-utils"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	type Obj struct {
		ID       int    `json:"id"`
		Children []*Obj `json:"children"`
	}
	t.Run("can unmarshal from map", func(t *testing.T) {
		var obj Obj
		assert.NotPanics(t, func() {
			goutils.UnmarshalJSONFromMap(map[string]interface{}{"id": 1, "children": []interface{}{map[string]interface{}{"id": 2}}}, &obj)
		})
		assert.Equal(t, 1, obj.ID)
		assert.Len(t, obj.Children, 1)
		assert.Equal(t, 2, obj.Children[0].ID)
	})
	t.Run("can marshal to map", func(t *testing.T) {
		var obj Obj
		obj.ID = 1
		obj.Children = make([]*Obj, 0)
		obj.Children = append(obj.Children, &Obj{ID: 2})
		m := make(map[string]interface{})
		goutils.MarshalJSONToMap(&obj, m)
		assert.EqualValues(t, 1, m["id"])
		assert.Len(t, m["children"], 1)
		assert.IsType(t, make([]interface{}, 0), m["children"])
		child := m["children"].([]interface{})[0]
		assert.IsType(t, make(map[string]interface{}), child)
		assert.EqualValues(t, 2, child.(map[string]interface{})["id"])
	})
}
