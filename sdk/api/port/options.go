package port

// Option is a function that configures the Port struct
type Option func(*Port)

// WithMinPort sets the lower bound of the port range
func WithMinPort(arg int) Option {
	return func(b *Port) {
		b.minPort = arg
	}
}

// WithMaxPort sets the upper bound of the port range
func WithMaxPort(arg int) Option {
	return func(b *Port) {
		b.maxPort = arg
	}
}
