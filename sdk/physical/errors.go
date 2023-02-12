package physical

import (
	"fmt"
	"strings"
)

// Errors
const (
	ValueTooLargeErrorMessage = "put failed due to value being too large"
	RelativePathErrorMessage  = "relative paths not supported"
	NonUTF8ErrorMessage       = "key contains invalid UTF-8 characters"
	NonPrintableErrorMessage  = "key contains non-printable characters"
)

// BackendInitializedFailedErrorMessage returns a formatted
// error message for a backend initialization failure
func BackendInitializedFailedErrorMessage(backend string, msg string) string {
	backend = strings.TrimSpace(backend)
	msg = strings.TrimSpace(msg)
	result := fmt.Sprintf("<%s> Backend - initialization failure", backend)
	if msg != "" {
		result = fmt.Sprintf("%s : %s", result, msg)
	}
	return result
}
func operationErrorMessage(backend, op, msg string) string {
	backend = strings.TrimSpace(backend)
	msg = strings.TrimSpace(msg)
	result := fmt.Sprintf("<%s> Backend - %s operation failed", backend, op)
	if msg != "" {
		result = fmt.Sprintf("%s : %s", result, msg)
	}
	return result
}

// PutErrorMessage returns a formatted error message for a put operation
func PutErrorMessage(backend, msg string) string {
	return operationErrorMessage(backend, "Put", msg)
}

// GetErrorMessage returns a formatted error message for a get operation
func GetErrorMessage(backend, msg string) string {
	return operationErrorMessage(backend, "Get", msg)
}

// ListErrorMessage returns a formatted error message for a list operation
func ListErrorMessage(backend, msg string) string {
	return operationErrorMessage(backend, "List", msg)
}

// DeleteErrorMessage returns a formatted error message for a delete operation
func DeleteErrorMessage(backend, msg string) string {
	return operationErrorMessage(backend, "Delete", msg)
}

type errNotExist string

// ErrNotExist is the error a physical backend returns in case path does not exist
func ErrNotExist(path string) errNotExist { //revive:disable:unexported-return
	return errNotExist(path)
	//revive:enable:unexported-return
}

// Error satisfies the error interface
func (e errNotExist) Error() string {
	s := string(e)
	return fmt.Sprintf("path does not exist : %s", s)
}

// IsErrNotExist returns true if the error is of type errNotExist
func IsErrNotExist(err error) bool {
	_, ok := err.(errNotExist)
	if !ok {
		ok = strings.Contains(err.Error(), "path does not exist")
	}
	return ok
}

// ─── TRANSACTIONAL ERRORS ───────────────────────────────────────────────────────
var _ error = ErrTransactionalOperations{}

// ErrTransactionalOperations represent an error that may occur in a transaction
//
//go:generate gomodifytags -add-options json=omitempty -file $GOFILE -struct ErrTransactionalOperations -add-tags json -w -transform snakecase
type ErrTransactionalOperations struct {
	Key       string `json:"key,omitempty"`
	Operation string `json:"operation,omitempty"`
	Message   string `json:"message,omitempty"`
}

// ErrTransactionalOperation is the constructor of errTransactionalOperation
func ErrTransactionalOperation(txn TxnEntry, msg string) ErrTransactionalOperations {
	return ErrTransactionalOperations{
		Key:       txn.Entry.Key,
		Operation: string(txn.Operation),
		Message:   msg,
	}
}

// Error satisfies the error interface
func (e ErrTransactionalOperations) Error() string {
	result := fmt.Sprintf("transactional [ %s ] operation on key < %s > failed", e.Operation, e.Key)
	if e.Message != "" {
		result = fmt.Sprintf("%s. Cause : %s", result, e.Message)
	}
	return result
}
