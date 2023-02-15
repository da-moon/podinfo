package server

import (
	"os"
	"strconv"
	"strings"

	core "github.com/da-moon/northern-labs-interview/api/core"
	flagset "github.com/da-moon/northern-labs-interview/internal/cli/flagset"
	value "github.com/da-moon/northern-labs-interview/internal/cli/value"
	logger "github.com/da-moon/northern-labs-interview/internal/logger"
)

// ServerFlags is a struct that
// contains the flags for used with configuring
// the Podinfo-API server daemon.
// it enables easy extraction of values.
type ServerFlags struct {
	*flagset.FlagSet
	logLevel value.String
	dev      value.Bool
	cache    value.Bool
	nodeName value.String
	apiAddr  value.String
}

// init function must be called after Creating
// a new ConsulFlags object or else
// the flag values would not be accepted
func (f *ServerFlags) init() {
	f.FlagSet = flagset.New("Server", help)
	f.Var(&f.logLevel, "log-level",
		[]string{
			"flag used to set stdlogger level.",
			"This can also be specified via the 'PODINFO_LOG_LEVEL' env variable.",
		})
	f.Var(&f.dev, "dev",
		[]string{
			"flag used to start the server in development mode",
			"This can also be specified via the 'PODINFO_DEVEL' env variable (true|false)",
		})
	f.Var(&f.nodeName, "node-name",
		[]string{
			"flag used to set podinfo node name.",
			"This can also be specified via the 'PODINFO_NODE_NAME' env variable.",
		})
	f.Var(&f.apiAddr, "api-addr",
		[]string{
			"flag used to set the address podinfo is listening on.",
			"This can also be specified via the 'PODINFO_API_ADDR' env variable.",
		})
	f.Var(&f.cache, "disable-cache",
		[]string{
			"flag used to disable redis cache and related endpoints",
			"This can also be specified via the 'PODINFO_DISABLE_CACHE' env variable (true|false)",
		})

}

// LogLevel parses '-log-level' flag
func (f *ServerFlags) LogLevel() string {
	str := f.logLevel.Get()
	if str == "" {
		str = os.Getenv("PODINFO_LOG_LEVEL")
	}
	result := logger.LogLevel(strings.ToUpper(str))
	ok := false
	for _, val := range logger.DefaultLogLevels {
		if strings.EqualFold(str, string(val)) {
			ok = true
			break
		}
	}
	if !ok {
		result = logger.DebugLevel
	}
	return string(result)
}

// Dev parses '-dev' flag
func (f *ServerFlags) Dev() bool {
	result := f.dev.Get()
	if !result {
		result = core.DefaultDevelopmentMode()
	}
	return result
}

// DisableCache parses '-disable-cache' flag
func (f *ServerFlags) DisableCache() bool {
	result := f.cache.Get()
	envstr := strings.TrimSpace(os.Getenv("PODINFO_DISABLE_CACHE"))
	if envstr != "" {
		result, _ = strconv.ParseBool(envstr)
	}
	return result
}

// NodeName parses '-node-name' flag
func (f *ServerFlags) NodeName() string {
	result := f.nodeName.Get()
	if result == "" {
		var err error
		result, err = core.DefaultNodeName()
		if err != nil {
			panic(err)
		}
	}
	return result
}

// APIAddr parses '-apr-addr' flag
func (f *ServerFlags) APIAddr() string {
	result := f.apiAddr.Get()
	if result == "" {
		var err error
		result, err = core.DefaultAPIAddr()
		if err != nil {
			panic(err)
		}
	}
	return result
}
