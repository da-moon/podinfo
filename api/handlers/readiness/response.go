package readiness

type Status int

const (
	OK Status = iota
	Unavailable
	Unknown
)
