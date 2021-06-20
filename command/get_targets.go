package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Get one or many targets.
// The client uses the get_targets command to get target information.
type GetTargets struct {
	XMLName  xml.Name `xml:"get_targets"`
	TargetId string   `xml:"target_id,attr,omitempty"` // ID of single target to get.
	Filter   string   `xml:"filter,attr,omitempty"`    // Filter term to use to filter query.
	FiltId   string   `xml:"filt_id,attr,omitempty"`   // ID of filter to use to filter query.
	Trash    bool     `xml:"trash,attr,omitempty"`     // Whether to get the trashcan tasks instead.
	Tasks    bool     `xml:"tasks,attr,omitempty"`     // Whether to include list of tasks that use the target.
}

type GetTargetsResp struct {
	*CommandResp
	XMLName     xml.Name                   `xml:"get_targets_response"`
	Target      []*GetTargetsRespTarget    `xml:"target" json:"target"`
	Filters     *Filters                   `xml:"filters" json:"filters"`
	Sort        *Sort                      `xml:"sort" json:"sort"`
	Targets     *GetTargetsRespTargets     `xml:"targets" json:"targets"`
	TargetCount *GetTargetsRespTargetCount `xml:"target_count" json:"target_count"`
}

type GetTargetsRespTarget struct {
	Id                   string                        `xml:"id,attr" json:"id"`
	Owner                *Owner                        `xml:"owner" json:"owner"`                         // Owner of the target.
	Name                 string                        `xml:"name" json:"name"`                           // The name of the target.
	Comment              string                        `xml:"comment" json:"comment"`                     // The comment on the target.
	CreationTime         string                        `xml:"creation_time" json:"creation_time"`         // Date and time the target was created.
	ModificationTime     string                        `xml:"modification_time" json:"modification_time"` // Date and time the target was last modified.
	Writable             bool                          `xml:"writable" json:"writable"`                   // Whether the target is writable.
	InUse                bool                          `xml:"in_use" json:"in_use"`                       // Whether any tasks are using the target.
	Permissions          *Permissions                  `xml:"permissions" json:"permissions"`             // Permissions that the current user has on the target.
	UserTags             *UserTags                     `xml:"user_tags" json:"user_tags"`                 // Info on tags attached to the target.
	Hosts                string                        `xml:"hosts" json:"hosts"`                         // A list of hosts.
	ExcludeHosts         string                        `xml:"exclude_hosts" json:"exclude_hosts"`         // A list of hosts to exclude.
	MaxHosts             int                           `xml:"max_hosts" json:"max_hosts"`                 // The maximum number of hosts defined by the target.
	SshCredential        *GetTargetsRespSshCredential  `xml:"ssh_credential,omitempty" json:"ssh_credential"`
	SmbCredential        *GetTargetsRespSmbCredential  `xml:"smb_credential,omitempty" json:"smb_credential"`
	EsxiCredential       *GetTargetsRespEsxiCredential `xml:"esxi_credential,omitempty" json:"esxi_credential"`
	SnmpCredential       *GetTargetsRespSnmpCredential `xml:"snmp_credential,omitempty" json:"snmp_credential"`
	PortRange            string                        `xml:"port_range,omitempty" json:"port_range"` // Port range of the target.
	PortList             *GetTargetsRespPortList       `xml:"port_list" json:"port_list"`
	AliveTests           string                        `xml:"alive_tests" json:"alive_tests"`                       // Which alive tests to use.
	AllowSimultaneousIps bool                          `xml:"allow_simultaneous_ips" json:"allow_simultaneous_ips"` // Whether to scan multiple IPs of the same host simultaneously.
	ReverseLookupOnly    bool                          `xml:"reverse_lookup_only" json:"reverse_lookup_only"`       // Whether to scan only hosts that have names.
	ReverseLookupUnify   bool                          `xml:"reverse_lookup_unify" json:"reverse_lookup_unify"`     // Whether to scan only one IP when multiple IPs have the same name.
	Tasks                *GetTargetsRespTasks          `xml:"tasks" json:"tasks"`                                   // All tasks using the target.
}

type GetTargetsRespSshCredential struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the SSH LSC credential.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the credential.
	Port        string       `xml:"port" json:"port"`               // The port the LSCs will use.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the LSC credential is in the trashcan.
}

type GetTargetsRespSmbCredential struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the SMB LSC credential.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the credential.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the LSC credential is in the trashcan.
}

type GetTargetsRespEsxiCredential struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the ESXi LSC credential.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the credential.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the LSC credential is in the trashcan.
}

type GetTargetsRespSnmpCredential struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the SNMP credential.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the credential.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the credential is in the trashcan.
}

type GetTargetsRespPortList struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the port_list.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the credential.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the port_list is in the trashcan.
}

type GetTargetsRespTasks struct {
	Task []*GetTargetsRespTask `xml:"task" json:"task"`
}

type GetTargetsRespTask struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the task.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the task.
}

type GetTargetsRespTargets struct {
	Start int `xml:"start,attr" json:"start"` // First target.
	Max   int `xml:"max,attr" json:"max"`     // Maximum number of targets.

}

type GetTargetsRespTargetCount struct {
	Filtered int `xml:"filtered" json:"filtered"` // Number of targets after filtering.
	Page     int `xml:"page" json:"page"`         // Number of targets on current page.
}

func NewGetAllTargets() *GetTargets {
	return &GetTargets{}
}

func NewGetTargetsWithId(id string) *GetTargets {
	return &GetTargets{
		TargetId: id,
	}
}

func NewGetTargets(filter *GetTargets) *GetTargets {
	return filter
}

func (a *GetTargets) Command() string {
	return "get_targets"
}

func (a *GetTargets) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *GetTargets) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *GetTargets) GetRespStruct() interface{} {
	return &GetTargetsResp{}
}

func (a *GetTargets) Handler(data interface{}) (interface{}, error) {
	resp := data.(*GetTargetsResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp, nil
}
