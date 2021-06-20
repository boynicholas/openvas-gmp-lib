package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Modify an existing config.

// This command can perform four types of actions: 1) modifying preferences, 2) changing the family selection, 3) changing the NVT selection of families, or 4) modifying basic fields like the name or comment.
// Since version 21.04 it is possible to perform multiple of these actions at once in the order they appear in the command. Note that the changes can influence each other, so this can change the overall result.
// The first type modifies the preferences on the config. If a preference includes an NVT, then the preference is an NVT preference, otherwise the preference is a scanner preference. If the preference includes a value then the manager updates the value of the preference, otherwise the manager removes the preference. The value must be base64 encoded.
// The second type, with a family selection, lets the client modify the NVTs selected by the config at a family level. The idea is that the client modifies the entire family selection at once, so a client will often need to include many families in a family selection.
// The family_selection may include a growing element to indicate whether new families should be added to the selection. It may also include any number of family elements.
// A family element must always include a name and may include a growing element and an all element. The all element indicates whether all NVTs in the family should be selected. The growing element indicates whether new NVTs in the family should be added to the selection as they arrive. Leaving a family out of the family_selection is equivalent to including the family with all 0 and growing 0.
// The effect of the all 0 and growing 0 case is subtle: if all NVTs were selected then all are removed (effectively removing the family from the config). However if some NVTs were selected then they remain selected. As a result the client must include in the family_selection all families that must have all NVTs selected.
// The third option, NVT selections, must include a family and may include any number of NVTs. The manager updates each given family in the config to include only the given NVTs. If the family selection is also changed, this should be done first as it can also change the NVT selection of families.
// If there was no error with the command sent by the client, the manager will apply the changes to the config and will reply with a response code indicating success.

type ModifyConfig struct {
	XMLName  xml.Name `xml:"modify_config"`
	ConfigId string   `xml:"config_id,attr" json:"config_id"`  // ID of config to modify.
	Name     string   `xml:"name,omitempty" json:"name"`       // New name for the config.
	Comment  string   `xml:"comment,omitempty" json:"comment"` // New comment for the config.
	Scanner  string   `xml:"scanner,omitempty" json:"scanner"` // New scanner's UUID for the config.

}

type ModifyConfigPreference struct {
	Name  string           `xml:"name" json:"name"`             // The name of the preference to modify.
	Nvt   *ModifyConfigNvt `xml:"nvt,omitempty" json:"nvt"`     // NVT associated with preference to modify.
	Value string           `xml:"value,omitempty" json:"value"` // New value for preference.
}

type ModifyConfigNvt struct {
	Oid string `xml:"oid,omitempty" json:"oid"`
}

type ModifyConfigResp struct {
	*CommandResp
}

func NewModifyConfig(config *ModifyConfig) *ModifyConfig {
	return config
}

func (a *ModifyConfig) Command() string {
	return "modify_config"
}

func (a *ModifyConfig) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *ModifyConfig) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *ModifyConfig) GetRespStruct() interface{} {
	return &ModifyConfigResp{}
}

func (a *ModifyConfig) Handler(data interface{}) (interface{}, error) {
	resp := data.(*ModifyConfigResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return nil, nil
}
