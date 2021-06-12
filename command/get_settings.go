package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Get one or many settings
// The client uses the get_settings command to get the settings.

type GetSettings struct {
	XMLName   xml.Name `xml:"get_settings"`
	SettingId string   `xml:"setting_id,attr,omitempty" json:"setting_id"` // ID of single setting to get.
	Filter    string   `xml:"filter,attr,omitempty" json:"filter"`         // Filter term.
	First     int      `xml:"first,attr,omitempty" json:"first"`           // First result.
	Max       int      `xml:"max,attr,omitempty" json:"max"`               // Maximum number of results in response.
	SortOrder string   `xml:"sort_order,attr,omitempty" json:"sort_order"`
	SortField string   `xml:"sort_field,attr,omitempty" json:"sort_field"`
}

type GetSettingsResp struct {
	*CommandResp
	XMLName      xml.Name                     `xml:"get_settings_response"`
	Term         *GetSettingsRespFilters      `xml:"term" json:"term"`
	Settings     *GetSettingsRespSettings     `xml:"settings" json:"settings"`
	SettingCount *GetSettingsRespSettingCount `xml:"setting_count" json:"setting_count"`
}

type GetSettingsRespFilters struct {
	Term string `xml:"term" json:"term"` // Filter term.
}

type GetSettingsRespSettings struct {
	Start   int                       `xml:"start,attr" json:"start"` // First setting.
	Max     int                       `xml:"max,attr" json:"max"`     // Maximum number of settings.
	Setting []*GetSettingsRespSetting `xml:"setting" json:"setting"`
}

type GetSettingsRespSetting struct {
	Name    string `xml:"name" json:"name"`
	Comment string `xml:"comment" json:"comment"`
	Value   string `xml:"value" json:"value"`
}

type GetSettingsRespSettingCount struct {
	Filtered int `xml:"filtered" json:"filtered"` // Number of settings after filtering.
	Page     int `xml:"page" json:"page"`         // Number of settings on current page.
}

func NewGetSettingsWithSingle(settingId string) *GetSettings {
	return &GetSettings{
		SettingId: settingId,
	}
}

func NewGetSettings(first int, max int, filter string) *GetSettings {
	return &GetSettings{
		First:  first,
		Max:    max,
		Filter: filter,
	}
}

func (a *GetSettings) Command() string {
	return "get_settings"
}

func (a *GetSettings) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *GetSettings) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *GetSettings) GetRespStruct() interface{} {
	return &GetSettingsResp{}
}

func (a *GetSettings) Handler(data interface{}) (interface{}, error) {
	resp := data.(*GetSettingsResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp, nil
}
