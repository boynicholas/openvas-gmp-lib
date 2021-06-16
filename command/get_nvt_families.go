package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Get a list of all NVT families.
// The client uses the get_nvt_families command to get NVT family information. If the command sent by the client was valid, the manager will reply with a list of NVT families to the client.

type GetNvtFamilies struct {
	XMLName   xml.Name `xml:"get_nvt_families"`
	SortOrder string   `xml:"sort_order,attr,omitempty" json:"sort_order"`
}

type GetNvtFamiliesResp struct {
	*CommandResp
	XMLName  xml.Name                    `xml:"get_nvt_families_response"`
	Families *GetNvtFamiliesRespFamilies `xm;:"families" json:"families"`
}

type GetNvtFamiliesRespFamilies struct {
	Family []*GetNvtFamiliesRespFamily `xml:"family" json:"family"`
}

type GetNvtFamiliesRespFamily struct {
	Name        string `xml:"name" json:"name"`
	MaxNvtCount int    `xml:"max_nvt_count" json:"max_nvt_count"`
}

func NewGetNvtFamilies(sortOrder string) *GetNvtFamilies {
	return &GetNvtFamilies{
		SortOrder: sortOrder,
	}
}

func (a *GetNvtFamilies) Command() string {
	return "get_nvt_families"
}

func (a *GetNvtFamilies) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *GetNvtFamilies) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *GetNvtFamilies) GetRespStruct() interface{} {
	return &GetNvtFamiliesResp{}
}

func (a *GetNvtFamilies) Handler(data interface{}) (interface{}, error) {
	resp := data.(*GetNvtFamiliesResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp, nil
}
