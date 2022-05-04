package types

import (
	"bytes"
	"encoding/json"
)


func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
    encoder.SetEscapeHTML(false)
    encoder.Encode(data)
	return buff.Bytes()
}