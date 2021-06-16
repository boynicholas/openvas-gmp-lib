package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Create a config.
// The client uses the create_config command to create a new config.

// This command can be called in three ways:

// With a copy element. The new config is a copy of the existing config, with the given name.
// With an embedded get_configs response element. The config is created as defined by the get_configs response element. Name is actually optional in this case. The config is given the name of the config in the get_configs response. If there is already a config with this name, then a number is attached to the name to make it unique.
// With a scanner element referencing an OSP scanner. The config is then created by retrieving the expected preferences from the given scanner via OSP.

type CreateConfig struct {
	XMLName            xml.Name  `xml:"create_config"`
	Comment            string    `xml:"comment" json:"comment"`                           // A comment on the config.
	Copy               string    `xml:"copy" json:"copy"`                                 // The UUID of an existing config.
	GetConfigsResponse string    `xml:"get_configs_response" json:"get_configs_response"` // Response to get_configs command
	Scanner            string    `xml:"scanner" json:"scanner"`                           // The UUID of an OSP scanner to get config data from.
	Name               string    `xml:"name" json:"name"`                                 // A name for the config.
	UsageType          UsageType `xml:"usage_type" json:"usage_type"`                     // Usage type (scan or policy) for the config. Can overwrite the one in get_configs_response.
}

type CreateConfigResp struct {
	*CommandResp
	XMLName xml.Name `xml:"create_config_response"`
	Id      string   `xml:"id,attr" json:"id"`
}

func NewCreateConfig(cfg *CreateConfig) *CreateConfig {
	return cfg
}

func (a *CreateConfig) Command() string {
	return "create_config"
}

func (a *CreateConfig) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *CreateConfig) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *CreateConfig) GetRespStruct() interface{} {
	return &CreateConfigResp{}
}

func (a *CreateConfig) Handler(data interface{}) (interface{}, error) {
	resp := data.(*CreateConfigResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp.Id, nil
}
