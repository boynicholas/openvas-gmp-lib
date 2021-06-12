package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Synchronize a config with a scanner.
// The client uses the "sync_config" command to request an OSP config synchronization with a scanner, adding new parameters and removing old ones.
type SyncConfig struct {
	XMLName  xml.Name `xml:"sync_config"`
	ConfigId string   `xml:"config_id,attr" json:"configId"`
}

type SyncConfigResp struct {
	*CommandResp
	XMLName xml.Name `xml:"sync_config_response"`
}

func NewSyncConfig(configId string) *SyncConfig {
	return &SyncConfig{
		ConfigId: configId,
	}
}

func (a *SyncConfig) Command() string {
	return "sync_config"
}

func (a *SyncConfig) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *SyncConfig) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *SyncConfig) GetRespStruct() interface{} {
	return &SyncConfigResp{}
}

func (a *SyncConfig) Handler(data interface{}) (interface{}, error) {
	resp := data.(*SyncConfigResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return nil, nil
}
