package command

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// Manually start an existing task.
// The client uses the start_task command to manually start an existing task.
type StartTask struct {
	XMLName xml.Name `xml:"start_task"`
	TaskId  string   `xml:"task_id,attr" json:"task_id"`
}

type StartTaskResp struct {
	*CommandResp
	XMLName  xml.Name `xml:"start_task_response"`
	ReportId string   `xml:"report_id" json:"report_id"`
}

func NewStartTask(taskId string) *StartTask {
	return &StartTask{
		TaskId: taskId,
	}
}

func (a *StartTask) Command() string {
	return "start_task"
}

func (a *StartTask) String() string {
	data, _ := xml.Marshal(a)
	return string(data)
}

func (a *StartTask) StringWithJson() string {
	data, _ := json.Marshal(a)
	return string(data)
}

func (a *StartTask) GetRespStruct() interface{} {
	return &StartTaskResp{}
}

func (a *StartTask) Handler(data interface{}) (interface{}, error) {
	resp := data.(*StartTaskResp)

	if !HasSuccess(resp.Status) {
		return nil, errors.New(resp.StatusText)
	}

	return resp.ReportId, nil
}
