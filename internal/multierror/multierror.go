package multierror

import (
	"errors"
	"strings"
	"sync"
)

type MultiError struct {
	sync.Mutex
	Errs []error
}

func (m *MultiError) Add(err string) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	err = strings.TrimSpace(err)
	if err != "" {
		m.Errs = append(m.Errs, errors.New(err))
	}
}

func (m *MultiError) HasError() error {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	if len(m.Errs) == 0 {
		return nil
	}
	return m
}

func (m *MultiError) Error() string {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	formattedError := make([]string, len(m.Errs))
	for i, e := range m.Errs {
		formattedError[i] = e.Error()
	}

	return strings.Join(formattedError, ", ")
}
func (m *MultiError) Unwrap() error {
	err := m.Error()
	err = strings.TrimSpace(err)
	if err == "" {
		return nil
	}
	return errors.New(m.Error())
}
