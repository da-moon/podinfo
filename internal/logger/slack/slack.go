package slack

import (
	slackrus "github.com/johntdyer/slackrus"
	stacktrace "github.com/palantir/stacktrace"
	logrus "github.com/sirupsen/logrus"
)

// Logger is the main function clients would call after
// initializing configuration. This function is responsible
// for fulfilling core responsibility of this library ,
// which is returning a Slack logger.
func (c *Config) Logger() (*logrus.Logger, error) {
	err := c.Validate()
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a Slack logger")
		return nil, err
	}
	result := logrus.StandardLogger()
	result.SetOutput(c.log.Writer())
	c.log.Debug("slack : initialized Standard Logrus logger")
	result.SetFormatter(&logrus.JSONFormatter{})
	c.log.Debug("slack : successfully set Logrus logger's formatter to default JSON formatter")
	logLevel, err := logrus.ParseLevel(c.LogLevel())
	// [ NOTE ] this code path must be unreachable as Validate() functions will
	// fail and return an error earlier.
	if err != nil {
		err = stacktrace.Propagate(err, "could not create a Slack logger")
		return nil, err
	}
	result.SetLevel(logLevel)
	result.AddHook(&slackrus.SlackrusHook{
		HookURL:        c.SlackHookURL(),
		AcceptedLevels: slackrus.LevelThreshold(logLevel),
		Channel:        c.Channel(),
		IconEmoji:      c.IconEmoji(),
		Username:       c.Username(),
	})
	return result, nil
}
