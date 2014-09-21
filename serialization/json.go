// json.go
package serialization

import (
	"encoding/json"
)

type JSON struct {
}

func (j *JSON) Encode(value []interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func (j *JSON) Decode(data []byte) ([]interface{}, error) {
	var value []interface{}
	err := json.Unmarshal(data, value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
