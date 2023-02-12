package address

import (
	"strconv"

	stacktrace "github.com/palantir/stacktrace"
)

// SetHost sets the host of the address
func (a *Address) SetHost(s string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	_, err := Parse(a.proto, s, a.Port)
	if err != nil {
		return err
	}
	a.Host = s
	return nil
}

// SetPort sets the port of the address
func (a *Address) SetPort(s string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	_, err := strconv.Atoi(s)
	if err != nil {
		err = stacktrace.Propagate(err, "error parsing port %s", s)
		return err
	}
	_, err = Parse(a.proto, a.Host, s)
	if err != nil {
		return err
	}
	a.Port = s
	return nil
}
