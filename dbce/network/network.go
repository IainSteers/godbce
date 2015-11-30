package dbcego

import (
	"encoding/json"
	"fmt"
)

type Network struct {
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Platform *Platform `json:"platform,omitempty"`
	Subnet   string    `json:"subnet,omitempty"`
	Gateway  string    `json:"gateway,omitempty"`
}
