package dbcego

import (
	"encoding/json"
	"fmt"
)

type Platform struct {
	Id      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	OsTypes []string `json:"ostypes,omitempty"`
}
