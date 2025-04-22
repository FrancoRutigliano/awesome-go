package model

import (
	"encoding/base64"
	"encoding/json"
)

type Cursor struct {
	LastId int `json:"last_id"`
}

func (c *Cursor) Encode() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(data), nil
}

func DecodeCursor(encoded string) (*Cursor, error) {
	data, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var cursor Cursor
	err = json.Unmarshal(data, &cursor)
	if err != nil {
		return nil, err
	}
	return &cursor, nil
}
