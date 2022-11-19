package dto

import "encoding/json"

type MessageJSON struct {
	Message string `json:"message"`
}

func (m MessageJSON) Marshal() []byte {
	b, _ := json.Marshal(m) //nolint:errchkjson

	return b
}
