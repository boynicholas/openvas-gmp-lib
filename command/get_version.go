package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

type GetVersion struct {
	XMLName xml.Name `xml:"get_version"`
}

type GetVersionResp struct {
	XMLName xml.Name `xml:"get_version_response"`
	*CommandResp

	Version string `xml:"version"`
}

func NewGetVersion() *GetVersion {
	return &GetVersion{}
}

func (a *GetVersion) Command() string {
	return "authenticate"
}

func (a *GetVersion) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *GetVersion) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *GetVersion) GetRespStruct() interface{} {
	return &GetVersionResp{}
}

func (a *GetVersion) Handler(data interface{}) (interface{}, error) {
	resp := data.(*GetVersionResp)

	if resp.Status != "200" {
		return nil, errors.New(resp.StatusText)
	}

	return resp.Version, nil
}
