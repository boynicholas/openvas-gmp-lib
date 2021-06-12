package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Create a target
// The client uses the create_target command to create a new target.
// If the list of hosts is empty, the command must also include a target locator.
type CreateTarget struct {
	XMLName              xml.Name                    `xml:"create_target"`
	Name                 string                      `xml:"name" json:"name"`                                               // A name for the target.
	Comment              string                      `xml:"comment,omitempty" json:"comment"`                               // A comment on the target.
	Copy                 string                      `xml:"copy,omitempty" json:"copy"`                                     // The UUID of an existing target.
	AssetHosts           *CreateTargetAssetHosts     `xml:"asset_hosts,omitempty" json:"asset_hosts"`                       // This is mutually exclusive with `hosts`, Hosts from which to create the target,
	Hosts                string                      `xml:"hosts,omitempty" json:"hosts"`                                   // This is mutually exclusive with `asset_hosts`, A comma-separated list of hosts, which may be empty.
	ExcludeHosts         string                      `xml:"exclude_hosts,omitempty" json:"exclude_hosts"`                   // A list of hosts to exclude.
	SshCredential        *CreateTargetSshCredential  `xml:"ssh_credential,omitempty" json:"ssh_credential"`                 // SSH login credentials for target.
	SmbCredential        *CreateTargetSmbCredential  `xml:"smb_credential,omitempty" json:"smb_credential"`                 // SMB login credentials for target.
	EsxiCredential       *CreateTargetEsxiCredential `xml:"esxi_credential,omitempty" json:"esxi_credential"`               // ESXi credential to use on target.
	SnmpCredential       *CreateTargetSnmpCredential `xml:"snmp_credential,omitempty" json:"snmp_credential"`               // SNMP credentials to use on target.
	AliveTests           string                      `xml:"alive_tests,omitempty" json:"alive_tests"`                       // Which alive tests to use.
	AllowSimultaneousIps bool                        `xml:"allow_simultaneous_ips,omitempty" json:"allow_simultaneous_ips"` // Whether to scan multiple IPs of the same host simultaneously.
	ReverseLookupOnly    bool                        `xml:"reverse_lookup_only,omitempty" json:"reverse_lookup_only"`       // Whether to scan only hosts that have names.
	ReverseLookupUnify   bool                        `xml:"reverse_lookup_unify,omitempty" json:"reverse_lookup_unify"`     // Whether to scan only one IP when multiple IPs have the same name.
	PortRange            string                      `xml:"port_range,omitempty" json:"port_range"`                         // This is mutually exclusive with `port_list`, Comma separated list of port ranges for the target (allowing whitespace).
	PortList             *CreateTargetPortList       `xml:"port_list,omitempty" json:"port_list"`                           // Port list for the target.
}

type CreateTargetAssetHosts struct {
	Filter string `xml:"filter,attr" json:"filter"`
}

type CreateTargetSshCredential struct {
	Id   string `xml:"id,attr" json:"id"`
	Port string `xml:"port" json:"port"`
}

type CreateTargetSmbCredential struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTargetEsxiCredential struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTargetSnmpCredential struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTargetPortList struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTargetResp struct {
	*CommandResp
	XMLName xml.Name `xml:"create_target_response"`
	Id      string   `xml:"id,attr" json:"id"`
}

func NewCreateTarget(target *CreateTarget) *CreateTarget {
	return target
}

func (a *CreateTarget) Command() string {
	return "create_target"
}

func (a *CreateTarget) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *CreateTarget) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *CreateTarget) GetRespStruct() interface{} {
	return &CreateTargetResp{}
}

func (a *CreateTarget) Handler(data interface{}) (interface{}, error) {
	resp := data.(*CreateTargetResp)

	if resp.Status != "201" {
		return nil, errors.New(resp.StatusText)
	}

	return resp.Id, nil
}
