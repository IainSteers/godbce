package dbcego

import (
	"encoding/json"
	"fmt"
)

type Capacity struct {
	data *Quantities `json:"data,omitempty"`
}
