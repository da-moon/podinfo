package readiness

// Status is an enum representing accepted status values for
// this request handler
type Status int

const (
	// OK represents the OK (ready) response
	OK Status = iota
	// Unavailable represents the Unavailable (not ready) response
	Unavailable
	// Unknown is the 'catch all' option
	Unknown
)

// String function returns a string representation of
// Status enum
func (s Status) String() string {
	Status := [...]string{"OK", "SERVICE UNAVAILABLE", "unknown"}
	if len(Status) < int(s) {
		return ""
	}
	return Status[s]
}
