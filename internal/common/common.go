package common

import "encoding/json"

// JSON alias type
type JSON = map[string]interface{}

func ToJSON(v interface{}) JSON {
	var jsonObj JSON
	respStr, _ := json.Marshal(v)
	json.Unmarshal(respStr, &jsonObj)

	return jsonObj
}
