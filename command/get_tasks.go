package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Get one or many tasks.
// The client uses the get_tasks command to get task information.
// As a convenience for clients the response includes a task count and the values of the sort order, sort field and apply overrides flag that the manager applied when selecting the tasks.
// A task may be a "container" task. This means the task may not be run, but reports may be uploaded to the task with the command create_report. Container tasks are identified by having an empty id attribute in the target element.
type GetTasks struct {
	XMLName          xml.Name  `xml:"get_tasks"`
	TaskId           string    `xml:"task_id,attr,omitempty"`      // ID of single task to get.
	Filter           string    `xml:"filter,omitempty"`            // Filter term to use to filter query.
	FiltId           string    `xml:"filt_id,omitempty"`           // ID of filter to use to filter query.
	Trash            bool      `xml:"trash,omitempty"`             // Whether to get the trashcan tasks instead.
	Details          bool      `xml:"details,omitempty"`           // Whether to include full task details.
	IgnorePagination bool      `xml:"ignore_pagination,omitempty"` // Whether to ignore info used to split the report into pages like the filter terms "first" and "rows"..
	SchedulesOnly    bool      `xml:"schedules_only,omitempty"`    // Whether to only include id, name and schedule details.
	UsageType        UsageType `xml:"usage_type,omitempty"`        // Optional usage type to limit the tasks to. Affects total count unlike filter.
}

type GetTasksResp struct {
	*CommandResp
	XMLName        xml.Name               `xml:"get_tasks_response"`
	ApplyOverrides int                    `xml:"apply_overrides"`
	Task           []*GetTasksRespTask    `xml:"task" json:"task"`
	Filters        *Filters               `xml:"filters" json:"filters"`
	Sort           *Sort                  `xml:"sort" json:"sort"`
	Tasks          *GetTasksRespTasks     `xml:"tasks" json:"tasks"`
	TaskCount      *GetTasksRespTaskCount `xml:"task_count" json:"task_count"`
}

type GetTasksRespTask struct {
	Id               string                     `xml:"id,attr" json:"id"`
	Owner            *Owner                     `xml:"owner" json:"owner"`                         // Owner of the task.
	Name             string                     `xml:"name" json:"name"`                           // The name of the task.
	Comment          string                     `xml:"comment" json:"comment"`                     // The comment on the task.
	CreationTime     string                     `xml:"creation_time" json:"creation_time"`         // Creation time of the task.
	ModificationTime string                     `xml:"modification_time" json:"modification_time"` // Last time the task was modified.
	Writable         bool                       `xml:"writable" json:"writable"`                   // Whether the task is writable.
	InUse            bool                       `xml:"in_use" json:"in_use"`                       // Whether this task is currently in use.
	Permissions      *Permissions               `xml:"permissions" json:"permissions"`             // Permissions that the current user has on the task.
	UserTags         *UserTags                  `xml:"user_tags" json:"user_tags"`                 // Info on tags attached to the task.
	Status           GetTasksRespStatus         `xml:"status" json:"status"`                       // The run status of the task.
	Progress         float32                    `xml:"progress" json:"progress"`                   // The percentage of the task that is complete.
	Alterable        bool                       `xml:"alterable" json:"alterable"`                 // Whether the task is an Alterable Task.
	UsageType        UsageType                  `xml:"usage_type" json:"usage_type"`               // The usage type of the task (scan or audit).
	Config           *GetTasksRespConfig        `xml:"config" json:"config"`                       // The scan configuration used by the task.
	Target           *GetTasksRespTarget        `xml:"target" json:"target"`                       // The hosts scanned by the task.
	HostsOrdering    string                     `xml:"hosts_ordering" json:"hosts_ordering"`       // The order hosts are scanned in.
	Scanner          *GetTasksRespScanner       `xml:"scanner" json:"scanner"`                     // The scanner used to scan the target.
	Alert            *GetTasksRespAlert         `xml:"alert" json:"alert"`                         // An alert that applies to the task.
	Observers        *GetTasksRespObservers     `xml:"observers" json:"observers"`                 // Users allowed to observe this task.
	Schedule         *GetTasksRespSchedule      `xml:"schedule" json:"schedule"`                   // When the task will run.
	SchedulePeriods  int                        `xml:"schedule_periods" json:"schedule_periods"`   // A limit to the number of times the task will be scheduled, or 0 for no limit.
	ReportCount      *GetTasksRespReportCount   `xml:"report_count" json:"report_count"`           // Number of reports.
	Trend            GetTasksRespTrend          `xml:"trend" json:"trend"`
	CurrentReport    *GetTasksRespCurrentReport `xml:"current_report" json:"current_report"`
	LastReport       *GetTasksRespLastReport    `xml:"last_report" json:"last_report"`
	AverageDuration  string                     `xml:"average_duration" json:"average_duration"` // Average scan duration in seconds.
	ResultCount      string                     `xml:"result_count" json:"result_count"`         // Result count for the entire task.
	Preferences      *GetTasksRespPreferences   `xml:"preferences" json:"preferences"`
}

type GetTasksRespStatus string

const (
	TaskStatusDeleteRequested GetTasksRespStatus = "Delete Requested"
	TaskStatusDone            GetTasksRespStatus = "Done"
	TaskStatusNew             GetTasksRespStatus = "New"
	TaskStatusRequested       GetTasksRespStatus = "Requested"
	TaskStatusRunning         GetTasksRespStatus = "Running"
	TaskStatusStopRequested   GetTasksRespStatus = "Stop Requested"
	TaskStatusStopped         GetTasksRespStatus = "Stopped"
	TaskStatusInterrupted     GetTasksRespStatus = "Interrupted"
)

type GetTasksRespTrend string

const (
	TaskTrendUp   GetTasksRespTrend = "up"
	TaskTrendDown GetTasksRespTrend = "down"
	TaskTrendMore GetTasksRespTrend = "more"
	TaskTrendLess GetTasksRespTrend = "less"
	TaskTrendSame GetTasksRespTrend = "same"
)

type GetTasksRespConfig struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the config.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the config.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the config is in the trashcan.
}

type GetTasksRespTarget struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The hosts scanned by the task.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the target.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the target is in the trashcan.
}

type GetTasksRespScanner struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the scanner.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the task.
	Type        int          `xml:"type" json:"type"`               // Type of the scanner.
}

type GetTasksRespAlert struct {
	Id          string       `xml:"id,attr" json:"id"`
	Name        string       `xml:"name" json:"name"`               // The name of the alert.
	Permissions *Permissions `xml:"permissions" json:"permissions"` // Permissions the user has on the alert.
	Trash       bool         `xml:"trash" json:"trash"`             // Whether the alert is in the trashcan.
}

type GetTasksRespObservers struct {
	Group []*GetTasksRespGroup `xml:"group" json:"group"` // Group allowed to observe this task.
	Role  []*GetTasksRespRole  `xml:"role" json:"role"`   // Role allowed to observe this task.
}

type GetTasksRespGroup struct {
	Id   string `xml:"id,attr" json:"id"`
	Name string `xml:"name" json:"name"` // The name of the group.
}

type GetTasksRespRole struct {
	Id   string `xml:"id,attr" json:"id"`
	Name string `xml:"name" json:"name"` // The name of the role.
}

type GetTasksRespSchedule struct {
	Id        string `xml:"id,attr" json:"id"`
	Name      string `xml:"name" json:"name"`           // The name of the schedule.
	Trash     bool   `xml:"trash" json:"trash"`         // Whether the schedule is in the trashcan.
	Icalendar string `xml:"icalendar" json:"icalendar"` // iCalendar text containing the time data..
	Timezone  string `xml:"timezone" json:"timezone"`   // The timezone the schedule will follow..
}

type GetTasksRespReportCount struct {
	Finished int `xml:"finished" json:"finished"` // Number of reports where the scan completed.
}

type GetTasksRespCurrentReport struct {
	Report *GetTasksRespReport `xml:"report" json:"report"`
}

type GetTasksRespReport struct {
	Id        string `xml:"id,attr" json:"id"`
	Timestamp string `xml:"timestamp" json:"timestamp"`
}

type GetTasksRespLastReport struct {
	Report *GetTasksRespLastReportDetail `xml:"report" json:"report"`
}

type GetTasksRespLastReportDetail struct {
	*GetTasksRespReport
	ScanEnd     string                   `xml:"scan_end" json:"scan_end"`
	ResultCount *GetTasksRespResultCount `xml:"result_count" json:"result_count"` // Result counts for this report.
	Severity    float32                  `xml:"severity" json:"severity"`         // Result count for the entire task.
}

type GetTasksRespResultCount struct {
	FalsePositive int `xml:"false_positive" json:"false_positive"`
	Log           int `xml:"log" json:"log"`
	Info          int `xml:"info" json:"info"`
	Warning       int `xml:"warning" json:"warning"`
	Hole          int `xml:"hole" json:"hole"`
}

type GetTasksRespPreferences struct {
	Preference []*GetTasksRespPreference `xml:"preference" json:"preference"`
}

type GetTasksRespPreference struct {
	Name        string `xml:"name" json:"name"`                 // Full name of preference, suitable for end users.
	ScannerName string `xml:"scanner_name" json:"scanner_name"` // Compact name of preference, from scanner.
	Value       string `xml:"value" json:"value"`
}

type GetTasksRespTasks struct {
	Start int `xml:"start,attr" json:"start"` // First task.
	Max   int `xml:"max,attr" json:"max"`     // Maximum number of tasks.

}

type GetTasksRespTaskCount struct {
	Filtered int `xml:"filtered" json:"filtered"` // Number of tasks after filtering.
	Page     int `xml:"page" json:"page"`         // Number of tasks on current page.
}

func NewGetAllTasks() *GetTasks {
	return &GetTasks{}
}

func NewGetTaskWithId(id string) *GetTasks {
	return &GetTasks{
		TaskId: id,
	}
}

func NewGetTask(filter *GetTasks) *GetTasks {
	return filter
}

func (a *GetTasks) Command() string {
	return "get_tasks"
}

func (a *GetTasks) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *GetTasks) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *GetTasks) GetRespStruct() interface{} {
	return &GetTasksResp{}
}

func (a *GetTasks) Handler(data interface{}) (interface{}, error) {
	resp := data.(*GetTasksResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp, nil
}
