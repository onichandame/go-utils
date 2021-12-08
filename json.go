package goutils

import "encoding/json"

func UnmarshalJSONFromMap(m map[string]interface{}, ent interface{}) {
	bytes, err := json.Marshal(m)
	Assert(err)
	Assert(json.Unmarshal(bytes, ent))
}

func MarshalJSONToMap(ent interface{}, m map[string]interface{}) {
	bytes, err := json.Marshal(ent)
	Assert(err)
	json.Unmarshal(bytes, &m)
}
