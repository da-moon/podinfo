package errors

import (
	"net/http"
	"sync"

	response "github.com/da-moon/podinfo/sdk/api/response"
	stacktrace "github.com/palantir/stacktrace"
)

// ─── MALFORMED UPDATE ENTRY ─────────────────────────────────────────────────────
var _ response.HTTPError = &errMalformedUpdateEntries{}

//go:generate gomodifytags -add-options json=omitempty -override -file $GOFILE -struct errMalformedUpdateEntries -add-tags json -w -transform snakecase
type errMalformedUpdateEntries struct {
	mutex         sync.RWMutex `json:"lock,omitempty"`
	MissingKeys   []int        `json:"missing_keys,omitempty"`
	MissingValues []struct {
		Key   string `json:"key,omitempty"`
		Index int    `json:"index,omitempty"`
	} `json:"missing_values,omitempty"`
	MalformedValues []struct {
		Key   string `json:"key,omitempty"`
		Index int    `json:"index,omitempty"`
	} `json:"malformed_values,omitempty"`
}

// ErrMalformedUpdateEntries is constructor of errMalformedUpdateEntries
func ErrMalformedUpdateEntries() *errMalformedUpdateEntries {
	return &errMalformedUpdateEntries{
		MissingKeys: []int{},
		MissingValues: []struct {
			Key   string `json:"key,omitempty"`
			Index int    `json:"index,omitempty"`
		}{},
		MalformedValues: []struct {
			Key   string `json:"key,omitempty"`
			Index int    `json:"index,omitempty"`
		}{},
	}
}

// Error satisfies the error interface
// it is the root error cause (no stacktrace)
func (e *errMalformedUpdateEntries) Error() string {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	msg := "Malformed request."
	if len(e.MissingKeys) > 0 {
		msg += "update entires had some elements with empty paths."
	}
	if len(e.MissingValues) > 0 {
		msg += "update entires had some elements with empty values."
	}
	if len(e.MalformedValues) > 0 {
		msg += "update entires had some elements with malformed values."
	}
	return stacktrace.NewMessageWithCode(e.StatusCode(), msg).Error()
}
func (e *errMalformedUpdateEntries) AppendMissingKey(idx int) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.MissingKeys == nil || len(e.MissingKeys) == 0 {
		e.MissingKeys = make([]int, 0)
	}
	e.MissingKeys = append(e.MissingKeys, idx)
}

func (e *errMalformedUpdateEntries) AppendMissingEntry(idx int, key string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.MissingValues == nil || len(e.MissingValues) == 0 {
		e.MissingValues = make([]struct {
			Key   string `json:"key,omitempty"`
			Index int    `json:"index,omitempty"`
		}, 0)
	}
	e.MissingValues = append(e.MissingValues, struct {
		Key   string `json:"key,omitempty"`
		Index int    `json:"index,omitempty"`
	}{Key: key, Index: idx})
}
func (e *errMalformedUpdateEntries) AppendMalformedEntry(idx int, key string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if e.MalformedValues == nil || len(e.MalformedValues) == 0 {
		e.MalformedValues = make([]struct {
			Key   string `json:"key,omitempty"`
			Index int    `json:"index,omitempty"`
		}, 0)
	}
	e.MalformedValues = append(e.MalformedValues, struct {
		Key   string `json:"key,omitempty"`
		Index int    `json:"index,omitempty"`
	}{Key: key, Index: idx})
}

// StatusCode returns 400 status code
func (e *errMalformedUpdateEntries) StatusCode() stacktrace.ErrorCode {
	return stacktrace.ErrorCode(http.StatusBadRequest)
}

// EncodeJSON encodes the error as JSON
func (e *errMalformedUpdateEntries) EncodeJSON() ([]byte, error) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	resp := &response.Response{
		Success: false,
		Body:    e,
	}
	return resp.EncodeJSON()
}
func (e *errMalformedUpdateEntries) IsError() bool {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return len(e.MissingKeys) > 0 && len(e.MissingValues) > 0 && len(e.MalformedValues) > 0
}
