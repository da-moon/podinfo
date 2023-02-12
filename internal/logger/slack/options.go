package slack

// Option is a variadic function, used to modify state of
// Config struct in constructor.
type Option func(*Config)

// WithSlackHookURL enables modification of
// slackHookUrl field of Config struct in Constructor function.
func WithSlackHookURL(arg string) Option {
	return func(c *Config) {
		c.slackHookURL = arg
	}
}

// WithUsername enables modification of
// username field of Config struct in Constructor function.
func WithUsername(arg string) Option {
	return func(c *Config) {
		c.username = arg
	}
}

// WithIconEmoji enables modification of
// iconEmoji field of Config struct in Constructor function.
func WithIconEmoji(arg string) Option {
	return func(c *Config) {
		c.iconEmoji = arg
	}
}

// WithChannel enables modification of
// iconEmoji field of Config struct in Constructor function.
func WithChannel(arg string) Option {
	return func(c *Config) {
		c.channel = arg
	}
}
