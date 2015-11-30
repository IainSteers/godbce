package dbcego

import (
	"encoding/json"
	"fmt"
)

type Machine struct {
	Id            string                `json:"id,omitempty"`
	Name          string                `json:"name,omitempty"`
	Image         *MachineImage         `json:"image,omitempty"`
	Configuration *MachineConfiguration `json:"configuration.omitempty"`
	Network       *Network              `json:"network,omitempty"`
	PublicIp      bool                  `json:"publicip,omitempty"`
	PublicKey     string                `json:"publickey,omitempty"`
	CloudConfig   string                `json:"cloudconfig,omitempty"`
	State         *MachineState         `json:"state,omitempty"`
	Platform      *Platform             `json:"platform,omitempty"`
	Compute       float32               `json:"compute,omitempty"`
	Memory        float32               `json:"memory,omitempty"`
	Storage       float32               `json:"storage,omitempty"`
	Addresses     []*MachineAddress     `json:"addresses,omitempty"`
}

type MachineImage struct {
	Id     string     `json:"id,omitempty"`
	Name   string     `json:"name,omitempty"`
	Type   *ImageType `json:"type,omitempty"`
	OsType *OsType    `json:"ostype,omitempty"`
}

type MachineConfiguration struct {
	Id      string  `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Compute float32 `json:"compute,omitempty"`
	Memory  float32 `json:"memory,omitempty"`
	Storage float32 `json:"storage,omitempty"`
}

type MachineState struct {
	State string
}

type MachineAddress struct {
	Address            string `json:"address,omitempty"`
	MachineAddressType string `json:"machineaddresstype,omitempty"`
}
