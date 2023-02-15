package core

import (
	"context"
	"net"

	address "github.com/da-moon/northern-labs-interview/sdk/api/address"
	stacktrace "github.com/palantir/stacktrace"
)

// Listener parsers listener address in
// which is supposed to be a string in
// form of "host:port" and returns a net.Listener object.
// In case address is empty, it would try to find a random port
// and return a listener on that port.
func (c *config) Listener() (*address.Address, error) {
	addr, err := address.TCPAddress(c.APIAddr)
	if err != nil {
		newAddrStr, ierr := DefaultAPIAddr()
		if ierr != nil {
			ierr = stacktrace.Propagate(ierr, "cannot extract listener from server configuration.Please provide a valid address")
			return nil, ierr
		}
		addr, err = address.TCPAddress(newAddrStr)
	}
	if err != nil {
		err = stacktrace.Propagate(err, "cannot extract listener from server configuration.Please provide a valid address")
		return nil, err
	}
	if addr == nil {
		err = stacktrace.NewError("cannot extract listener from server configuration.Please provide a valid address")
		return nil, err
	}
	return addr, nil
}

// initListener creates a new L4 tcp listener based on the give
// configuration
// NOTE(damoon) for now, this function does not
// need any arguments

func (s *RestfulServer) initListener(_ context.Context) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	//nolint:gocritic // this is a debug function
	// s.log.Info("initializing server TCP listener")
	listener, err := s.config.Listener()
	if err != nil {
		err = stacktrace.Propagate(
			err,
			"cannot initialize server on configured listener",
		)
		return err
	}
	if listener == nil {
		err = stacktrace.NewError(
			"cannot initialize server on configured listener",
		)
		return err
	}
	parsed, err := listener.Parse()
	if err != nil {
		err = stacktrace.Propagate(
			err,
			"cannot initialize server on configured listener",
		)
		return err
	}
	tcpListener, err := net.Listen("tcp", parsed.String())
	if err != nil {
		if tcpListener != nil {
			defer func() {
				derr := tcpListener.Close()
				if derr != nil {
					s.log.Error("failed to close listener: %v", derr)
				}
				return
			}()
		}
		err = stacktrace.Propagate(
			err,
			"cannot initialize server on configured listener: %s",
			parsed.String(),
		)
		return err
	}
	s.listener = tcpListener
	return nil
}

// Port function returns actual port server is listening on
func (s *RestfulServer) Port() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.listener.Addr().(*net.TCPAddr).Port
}
