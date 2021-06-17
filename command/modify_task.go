package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Modify an existing task.
// The client uses the modify_task command to change an existing task.
type ModifyTask struct {
	XMLName         xml.Name               `xml:"modify_task"`
	TaskId          string                 `xml:"task_id,attr" json:"task_id"`          // ID of task to modify.
	Comment         string                 `xml:"comment,omitempty" json:"comment"`     // The comment on the task.
	Alert           []*ModifyTaskAlert     `xml:"alert,omitempty" json:"alert"`         // An alert that applies to the task.
	Config          *ModifyTaskConfig      `xml:"config,omitempty" json:"config"`       // The scan configuration used by the task.
	Name            string                 `xml:"name,omitempty" json:"name"`           // The name of the task.
	Observers       string                 `xml:"observers,omitempty" json:"observers"` // Users allowed to observe this task.
	Preferences     *ModifyTaskPreferences `xml:"preferences,omitempty" json:"preferences"`
	Schedule        *ModifyTaskSchedule    `xml:"schedule,omitempty" json:"schedule"`                 // When the task will run.
	SchedulePeriods int                    `xml:"schedule_periods,omitempty" json:"schedule_periods"` // A limit to the number of times the task will be scheduled, or 0 for no limit.
	Scanner         *ModifyTaskScanner     `xml:"scanner,omitempty" json:"scanner"`                   // The scanner to use for scanning the target.
	Target          *ModifyTaskTarget      `xml:"target,omitempty" json:"target"`                     // The hosts scanned by the task.
	File            *ModifyTaskFile        `xml:"file,omitempty" json:"file"`                         // File to attach to task.

}

type ModifyTaskConfig struct {
	Id string `xml:"id,attr" json:"id"`
}

type ModifyTaskTarget struct {
	Id string `xml:"id,attr" json:"id"`
}

type ModifyTaskScanner struct {
	Id string `xml:"id,attr" json:"id"`
}

type ModifyTaskAlert struct {
	Id string `xml:"id,attr" json:"id"`
}

type ModifyTaskSchedule struct {
	Id string `xml:"id,attr" json:"id"`
}

type ModifyTaskPreferences struct {
	Preference []*ModifyTaskPreference `xml:"preference" json:"preference"`
}

type ModifyTaskPreference struct {
	ScannerName string `xml:"scanner_name" json:"scanner_name"` // Compact name of preference, from scanner.
	Value       string `xml:"value" json:"value"`
}

type ModifyTaskFile struct {
	Name   string           `xml:"name,attr" json:"name"`
	Action ModifyTaskAction `xml:"action,attr" json:"action"`
}

type ModifyTaskAction string

const (
	ActionUpdate ModifyTaskAction = "update"
	ActionRemove ModifyTaskAction = "remove"
)

type ModifyTaskResp struct {
	*CommandResp
	XMLName xml.Name `xml:"modify_task_response"`
}

func NewModifyTask(task *ModifyTask) *ModifyTask {
	return task
}

func (a *ModifyTask) Command() string {
	return "modify_task"
}

func (a *ModifyTask) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *ModifyTask) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *ModifyTask) GetRespStruct() interface{} {
	return &ModifyTaskResp{}
}

func (a *ModifyTask) Handler(data interface{}) (interface{}, error) {
	resp := data.(*ModifyTaskResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return nil, nil
}
