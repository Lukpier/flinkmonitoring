package monitoring

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Jobs struct {
	// Jobs corresponds to the JSON schema field "jobs".
	Jobs []Job `json:"jobs,omitempty"`
}

type Job struct {
	// Id corresponds to the JSON schema field "id".
	Id *string `json:"id,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status *JobStatus `json:"status,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *JobStatus) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_JobStatus {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_JobStatus, v)
	}
	*j = JobStatus(v)
	return nil
}

type JobStatus string

const CANCELED JobStatus = "CANCELED"
const CANCELLING JobStatus = "CANCELLING"
const CREATED JobStatus = "CREATED"
const FAILED JobStatus = "FAILED"
const FAILING JobStatus = "FAILING"
const FINISHED JobStatus = "FINISHED"
const INITIALIZING JobStatus = "INITIALIZING"
const RECONCILING JobStatus = "RECONCILING"
const RESTARTING JobStatus = "RESTARTING"
const RUNNING JobStatus = "RUNNING"
const SUSPENDED JobStatus = "SUSPENDED"

var enumValues_JobStatus = []interface{}{
	"INITIALIZING",
	"CREATED",
	"RUNNING",
	"FAILING",
	"FAILED",
	"CANCELLING",
	"CANCELED",
	"FINISHED",
	"RESTARTING",
	"SUSPENDED",
	"RECONCILING",
}

type FlinkExceptions struct {
	// AllExceptions corresponds to the JSON schema field "all-exceptions".
	AllExceptions []Exception `json:"all-exceptions,omitempty"`

	// ExceptionHistory corresponds to the JSON schema field "exceptionHistory".
	ExceptionHistory *ExceptionHistory `json:"exceptionHistory,omitempty"`

	// RootException corresponds to the JSON schema field "root-exception".
	RootException *string `json:"root-exception,omitempty"`

	// Timestamp corresponds to the JSON schema field "timestamp".
	Timestamp *int `json:"timestamp,omitempty"`

	// Truncated corresponds to the JSON schema field "truncated".
	Truncated *bool `json:"truncated,omitempty"`
}

type Exception struct {
	// Exception corresponds to the JSON schema field "exception".
	Exception *string `json:"exception,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *string `json:"location,omitempty"`

	// Task corresponds to the JSON schema field "task".
	Task *string `json:"task,omitempty"`

	// Timestamp corresponds to the JSON schema field "timestamp".
	Timestamp *int `json:"timestamp,omitempty"`
}

type ExceptionHistory struct {
	// Entries corresponds to the JSON schema field "entries".
	Entries []Entry `json:"entries,omitempty"`

	// Truncated corresponds to the JSON schema field "truncated".
	Truncated *bool `json:"truncated,omitempty"`
}

type Entry struct {
	// ConcurrentExceptions corresponds to the JSON schema field
	// "concurrentExceptions".
	ConcurrentExceptions []ConcurrentException `json:"concurrentExceptions,omitempty"`

	// ExceptionName corresponds to the JSON schema field "exceptionName".
	ExceptionName *string `json:"exceptionName,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *string `json:"location,omitempty"`

	// Stacktrace corresponds to the JSON schema field "stacktrace".
	Stacktrace *string `json:"stacktrace,omitempty"`

	// TaskName corresponds to the JSON schema field "taskName".
	TaskName *string `json:"taskName,omitempty"`

	// Timestamp corresponds to the JSON schema field "timestamp".
	Timestamp *int `json:"timestamp,omitempty"`
}

type ConcurrentException struct {
	// ExceptionName corresponds to the JSON schema field "exceptionName".
	ExceptionName *string `json:"exceptionName,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *string `json:"location,omitempty"`

	// Stacktrace corresponds to the JSON schema field "stacktrace".
	Stacktrace *string `json:"stacktrace,omitempty"`

	// TaskName corresponds to the JSON schema field "taskName".
	TaskName *string `json:"taskName,omitempty"`

	// Timestamp corresponds to the JSON schema field "timestamp".
	Timestamp *int `json:"timestamp,omitempty"`
}

type IMailClient interface {
	SendMail(jobId string, body string) error
}

type IFlinkClient interface {
	GetJobs() (Jobs, error)
	LookupExceptions(jobId string) (FlinkExceptions, error)
}
