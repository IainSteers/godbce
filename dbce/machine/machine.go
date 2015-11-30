package dbcego

import (
	"encoding/json"
	"fmt"
)

const machinePath = "v0/machines"

type MachinesService interface {
	List(*ListOptions) ([]Machine, *Response, error)
	Get(string) (*Machine, *Response, error)
	Create(*MachineCreateRequest) (*Machine, *Response, error)
	Delete(string) (*Response, error)
}

type MachinesServiceOp struct {
	client *Client
}

var _ MachinesService = &MachinesServiceOp{}

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

type MachineCreateRequest struct {
	Machine *Machine
}

func (s *MachinesServiceOp) List(opt *ListOptions) ([]Machine, *Response, error) {
	path := machinePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	machines := new(machinesRoot)
	resp, err := s.client.Do(req, machines)
	if err != nil {
		return nil, nil, err
	}
	return machines, resp, err
}
