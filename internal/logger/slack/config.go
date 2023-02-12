package slack

import (
	"net/url"
	"os"
	"strings"
	"sync"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	stacktrace "github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
)

const (
	defaultSlackHookURL = "https://hooks.slack.com/services/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"
)

// Config is the configuration for the slack package.
type Config struct {
	mutex        sync.RWMutex
	log          *logger.WrappedLogger
	slackHookURL string
	username     string
	iconEmoji    string
	channel      string
	logLevel     string
}

// New returns a slack config with default values.
// it enables initialization of Slack logger.
func New(log *logger.WrappedLogger, opts ...Option) (*Config, error) {
	if log == nil {
		err := stacktrace.NewError("no logger was provided")
		return nil, err
	}
	result := &Config{
		log:          log,
		slackHookURL: DefaultSlackHookURL(),
		username:     DefaultUsername(),
		iconEmoji:    DefaultIconEmoji(),
		channel:      DefaultChannel(),
		logLevel:     DefaultLogLevel(),
	}
	for _, opt := range opts {
		opt(result)
	}
	return result, nil
}

//
// ──────────────────────────────────────────────────────── I ──────────
//   :::::: D E F A U L T S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────
//

// DefaultSlackHookURL returns default values for slack hook URL.
func DefaultSlackHookURL() string {
	result := os.Getenv("PODINFO_SLACK_HOOK_URL")
	if result == "" {
		result = defaultSlackHookURL
	}
	return result
}

// DefaultUsername returns default values of the username
// that authenticates against Slack API.
func DefaultUsername() string {
	result := os.Getenv("PODINFO_SLACK_USERNAME")
	if result == "" {
		result = "podinfo-bot"
	}
	return result
}

// DefaultIconEmoji returns default values of the icon emoji
// used when logs are sent to Slack.
func DefaultIconEmoji() string {
	result := os.Getenv("PODINFO_SLACK_ICON_EMOJI")
	if result == "" {
		result = ":mega:"
	}
	return result
}

// DefaultChannel returns default value of the Slack channel
// to which logs are sent.
func DefaultChannel() string {
	result := os.Getenv("PODINFO_SLACK_DEFAULT_CHANNEL")
	if result == "" {
		result = "#podinfo_notify"
	}
	return result
}

// DefaultLogLevel returns default value of the log level
// that is used when logs are sent to Slack.
func DefaultLogLevel() string {
	result := os.Getenv("PODINFO_SLACK_DEFAULT_LOG_LEVEL")
	if result == "" {
		result = "error"
	}
	return result
}

//
// ────────────────────────────────────────────────────── I ──────────
//   :::::: S E T T E R S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────
//

// SetSlackHookURL enables modification of
// slackHookURL field of the Config object after
// Constructor function has created the object.
func (c *Config) SetSlackHookURL(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.slackHookURL = arg
}

// SetUsername enables modification of
// username field of the Config object after
// Constructor function has created the object.
func (c *Config) SetUsername(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.username = arg
}

// SetIconEmoji enables modification of
// iconEmoji field of the Config object after
// Constructor function has created the object.
func (c *Config) SetIconEmoji(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.iconEmoji = arg
}

// SetChannel enables modification of
// channel field of the Config object after
// Constructor function has created the object.
func (c *Config) SetChannel(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.channel = arg
}

// SetLogLevel enables modification of
// logLevel field of the Config object after
// Constructor function has created the object.
func (c *Config) SetLogLevel(arg string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.logLevel = arg
}

//
// ────────────────────────────────────────────────────── I ──────────
//   :::::: G E T T E R S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────
//

// SlackHookURL returns the value stored in
// slackHookURL field in a concurrency safe manner.
func (c *Config) SlackHookURL() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.slackHookURL
}

// Username returns the value stored in
// username field in a concurrency safe manner.
func (c *Config) Username() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.username
}

// IconEmoji returns the value stored in
// iconEmoji field in a concurrency safe manner.
func (c *Config) IconEmoji() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.iconEmoji
}

// Channel returns the value stored in
// channel field in a concurrency safe manner.
func (c *Config) Channel() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.channel
}

// LogLevel returns the value stored in
// logLevel field in a concurrency safe manner.
func (c *Config) LogLevel() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.logLevel
}

//
// ──────────────────────────────────────────────────────────── I ──────────
//   :::::: V A L I D A T O R S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────
//

// ValidateSlackHookURL validates the value of slackHookURL field
// in a concurrency safe manner.
func (c *Config) ValidateSlackHookURL() error {
	val := c.SlackHookURL()
	val = strings.TrimSpace(val)
	c.SetSlackHookURL(val)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if val == "" {
		return stacktrace.NewError("slack webhook url is empty")
	}
	// ─── ENSURING INPUT IS VALID URI ────────────────────────────────────────────────
	_, err := url.Parse(c.SlackHookURL())
	if err != nil {
		return stacktrace.Propagate(err, "could not validate slack webhook url")
	}
	return nil
}

// ValidateUsername validates the value of username field
// in a concurrency safe manner.
func (c *Config) ValidateUsername() error {
	val := c.Username()
	val = strings.TrimSpace(val)
	c.SetUsername(val)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if val == "" {
		return stacktrace.NewError("username is empty")
	}
	return nil
}

// ValidateIconEmoji validates the value of iconEmoji field
// in a concurrency safe manner.
func (c *Config) ValidateIconEmoji() error {
	val := c.IconEmoji()
	val = strings.TrimSpace(val)
	c.SetIconEmoji(val)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if val == "" {
		return stacktrace.NewError("icon emoji is empty")
	}
	if !strings.HasPrefix(val, ":") {
		return stacktrace.NewError("icon emoji must start with a colon")
	}
	if !strings.HasSuffix(val, ":") {
		return stacktrace.NewError("icon emoji must end with a colon")
	}
	return nil
}

// ValidateChannel validates the value of channel field
// in a concurrency safe manner.
func (c *Config) ValidateChannel() error {
	val := c.Channel()
	val = strings.TrimSpace(val)
	c.SetChannel(val)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if val == "" {
		return stacktrace.NewError("slack channel name is empty")
	}
	if !strings.HasPrefix(val, "#") {
		return stacktrace.NewError("slack channel name must start with a hashtag")
	}
	return nil
}

// ValidateLogLevel validates the value of logLevel field
// in a concurrency safe manner.
func (c *Config) ValidateLogLevel() error {
	val := c.LogLevel()
	val = strings.TrimSpace(val)
	c.SetChannel(val)
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if val == "" {
		return stacktrace.NewError("log level is empty")
	}
	_, err := logrus.ParseLevel(val)
	if err != nil {
		return stacktrace.Propagate(err, "could not validate log level value")
	}
	return nil
}

// Validate function ensures that the value of all fields
// are valid. This function must be called before using the Config.
// Object to return a new Slack logger.
func (c *Config) Validate() error {
	c.log.Info("slack : validating configuration")
	var err error
	err = c.ValidateSlackHookURL()
	if err != nil {
		return stacktrace.Propagate(err, "could not validate slack configuration")
	}
	c.log.Info("slack : webhook url was successfully validated")
	err = c.ValidateUsername()
	if err != nil {
		return stacktrace.Propagate(err, "could not validate slack configuration")
	}
	c.log.Info("slack : username was successfully validated")
	err = c.ValidateIconEmoji()
	if err != nil {
		return stacktrace.Propagate(err, "could not validate slack configuration")
	}
	c.log.Info("slack : icon emoji was successfully validated")
	err = c.ValidateChannel()
	if err != nil {
		return stacktrace.Propagate(err, "could not validate slack configuration")
	}
	c.log.Info("slack : channel was successfully validated")
	err = c.ValidateLogLevel()
	if err != nil {
		return stacktrace.Propagate(err, "could not validate slack configuration")
	}
	c.log.Info("slack : log level was successfully validated")
	return nil
}
