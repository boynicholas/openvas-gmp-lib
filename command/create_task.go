package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Create a task.
// The client uses the create_task command to create a new task.
// When given a target with an id of 0, the command creates a "container" task. This kind of task can not be run, but it is possible to import reports into the task. Importing is done using the create_report command. The report being imported must be in the XML format.
// When creating a container task, the elements config, hosts_ordering, scanner, alert, schedule, schedule_periods, observers and preferences are ignored.
type CreateTask struct {
	XMLName         xml.Name               `xml:"create_task"`
	Name            string                 `xml:"name" json:"name"`                                   // A name for the task.
	Comment         string                 `xml:"comment,omitempty" json:"comment"`                   // A comment on the task.
	Copy            string                 `xml:"copy,omitempty" json:"copy"`                         // The UUID of an existing task.
	Alterable       bool                   `xml:"alterable,omitempty" json:"alterable"`               // Whether the task is alterable.
	UsageType       UsageType              `xml:"usage_type" json:"usage_type"`                       // Usage type for the task (scan or audit), defaulting to scan.
	Config          *CreateTaskConfig      `xml:"config" json:"config"`                               // The scan configuration used by the task.
	Target          *CreateTaskTarget      `xml:"target" json:"target"`                               // The hosts scanned by the task.
	HostsOrdering   string                 `xml:"hosts_ordering,omitempty" json:"hosts_ordering"`     // The order hosts are scanned in.
	Scanner         *CreateTaskScanner     `xml:"scanner" json:"scanner"`                             // The scanner to use for scanning the target.
	Alert           *CreateTaskAlert       `xml:"alert,omitempty" json:"alert"`                       // An alert that applies to the task.
	Schedule        *CreateTaskSchedule    `xml:"schedule,omitempty" json:"schedule"`                 // When the task will run.
	SchedulePeriods int                    `xml:"schedule_periods,omitempty" json:"schedule_periods"` // A limit to the number of times the task will be scheduled, or 0 for no limit.
	Observers       string                 `xml:"observers,omitempty" json:"observers"`               // Users allowed to observe this task.
	Preferences     *CreateTaskPreferences `xml:"preferences,omitempty" json:"preferences"`
}

type CreateTaskConfig struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTaskTarget struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTaskScanner struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTaskAlert struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTaskSchedule struct {
	Id string `xml:"id,attr" json:"id"`
}

type CreateTaskPreferences struct {
	Preference []*CreateTaskPreference `xml:"preference" json:"preference"`
}

type CreateTaskPreference struct {
	ScannerName string `xml:"scanner_name" json:"scanner_name"` // Compact name of preference, from scanner.
	Value       string `xml:"value" json:"value"`
}

type CreateTaskResp struct {
	*CommandResp
	XMLName xml.Name `xml:"create_task_response"`
	Id      string   `xml:"id,attr" json:"id"`
}

func NewCreateTask(task *CreateTask) *CreateTask {
	return task
}

func (a *CreateTask) Command() string {
	return "create_task"
}

func (a *CreateTask) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *CreateTask) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *CreateTask) GetRespStruct() interface{} {
	return &CreateTaskResp{}
}

func (a *CreateTask) Handler(data interface{}) (interface{}, error) {
	resp := data.(*CreateTaskResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp.Id, nil
}
