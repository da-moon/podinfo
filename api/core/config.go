package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	logger "github.com/da-moon/northern-labs-interview/internal/logger"
	metrics "github.com/da-moon/northern-labs-interview/sdk/api/metrics"
	port "github.com/da-moon/northern-labs-interview/sdk/api/port"
	mapstructure "github.com/mitchellh/mapstructure"
	stacktrace "github.com/palantir/stacktrace"
)

// config is the configuration for the agent.
//
//go:generate gomodifytags -override -file $GOFILE -struct config -add-tags json,yaml,mapstructure -w -transform snakecase
type config struct {
	mutex           sync.RWMutex `json:"lock" yaml:"lock" mapstructure:"lock"`
	DevelopmentMode bool         `mapstructure:"development_mode" json:"development_mode,omitempty" yaml:"development_mode"`
	NodeName        string       `mapstructure:"node_name" json:"node_name,omitempty" yaml:"node_name"`
	APIAddr         string       `mapstructure:"api_addr" json:"api_addr,omitempty" yaml:"api_addr"`
	// ─── METRICS ────────────────────────────────────────────────────────────────────
	MetricsPrefix           string                `mapstructure:"metrics_prefix" json:"metrics_prefix" yaml:"metrics_prefix"`
	StatsiteAddr            string                `mapstructure:"statsite_addr" json:"statsite_addr,omitempty" yaml:"statsite_addr"`
	StatsdAddr              string                `mapstructure:"statsd_addr" json:"statsd_addr,omitempty" yaml:"statsd_addr"`
	PrometheusRetentionTime time.Duration         `mapstructure:"prometheus_retention_time" json:"prometheus_retention_time" yaml:"prometheus_retention_time"`
	log                     *logger.WrappedLogger `json:"log" yaml:"log" mapstructure:"log"`
}

// Mergeconfig takes in two config objects
// and merges them into one, giving precedence to the second config.
// nolint:gocognit,gocyclo // this function is well tested and won't be used often
func Mergeconfig(a, b *config) *config { // revive:disable:unexported-return
	result := *a
	result.DevelopmentMode = b.DevelopmentMode
	if b.NodeName != "" {
		result.NodeName = b.NodeName
	}
	if b.APIAddr != "" {
		result.APIAddr = b.APIAddr
	}
	// ────────────────────────────────────────────────────────────────────────────────
	if b.MetricsPrefix != "" {
		result.MetricsPrefix = b.MetricsPrefix
	}
	if b.StatsiteAddr != "" {
		result.StatsiteAddr = b.StatsiteAddr
	}
	if b.StatsdAddr != "" {
		result.StatsdAddr = b.StatsdAddr
	}
	if b.PrometheusRetentionTime.Nanoseconds() >= 1 {
		result.PrometheusRetentionTime = b.PrometheusRetentionTime
	}
	return &result
	// revive:enable:unexported-return
}

//
// ──────────────────────────────────────────────────────────────────── I ──────────
//   :::::: D E F A U L T   V A L U E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────────────────────
//

// DefaultConfig returns a new config struct
func DefaultConfig(log *logger.WrappedLogger) (*config, error) { // revive:disable:unexported-return
	if log == nil {
		err := stacktrace.NewError("no logger was provided")
		return nil, err
	}
	// ─── DEVELOPMENT MODE DEFAULT ───────────────────────────────────────────────────
	developmentMode := DefaultDevelopmentMode()
	// ─── NODE NAME DEFAULT ──────────────────────────────────────────────────────────
	nodeName, err := DefaultNodeName()
	if err != nil {
		err = stacktrace.Propagate(err, "cannot prepare default api config struct")
		return nil, err
	}
	// ─── API ADDRESS DEFAULT ─────────────────────────────────────────────────────────
	apiAddr, err := DefaultAPIAddr()
	if err != nil {
		err = stacktrace.Propagate(err, "cannot prepare default api config struct")
		return nil, err
	}
	// ────────────────────────────────────────────────────────────────────────────────
	result := &config{
		log:             log,
		DevelopmentMode: developmentMode,
		NodeName:        nodeName,
		APIAddr:         apiAddr,
		// ─── METRICS ─────────────────────────────────────────────────────
		MetricsPrefix:           metrics.DefaultMetricsPrefix(),
		StatsiteAddr:            metrics.DefaultStatsiteAddr(),
		StatsdAddr:              metrics.DefaultStatsdAddr(),
		PrometheusRetentionTime: metrics.DefaultPrometheusRetentionTime(),
	}
	return result, nil
	// revive:enable:unexported-return
}

func DefaultDevelopmentMode() bool {
	var result bool
	podinfoDevelString := os.Getenv("PODINFO_DEVEL")
	if podinfoDevelString != "" {
		var err error
		result, err = strconv.ParseBool(podinfoDevelString)
		if err != nil {
			result = false
		}
	}
	return result
}
func DefaultNodeName() (string, error) {
	var err error
	result := os.Getenv("PODINFO_NODE_NAME")
	if result == "" {
		result, err = os.Hostname()
		if err != nil {
			err = stacktrace.Propagate(err, "cannot get default node name")
			return "", err
		}
	}
	return result, nil
}

// TODO: add validation to ensure address is valid when the user provided the address
func DefaultAPIAddr() (string, error) {
	apiAddr := os.Getenv("PODINFO_API_ADDR")
	if apiAddr == "" {
		tcpAddr, err := port.New().TCP()
		if err != nil {
			err = stacktrace.Propagate(err, "cannot find a random port to bind to")
			return "", err
		}
		apiAddr = fmt.Sprintf("0.0.0.0:%s", tcpAddr.Port)
	}
	return apiAddr, nil
}

//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: G E T T E R : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//

// Decodeconfig takes an io.reader and decodes
// underlying byte stream into a config struct
func Decodeconfig(r io.Reader) (*config, error) { // revive:disable:unexported-return
	var raw interface{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}
	var md mapstructure.Metadata
	var result config
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:    &md,
		Result:      &result,
		ErrorUnused: true,
	})
	if err != nil {
		return nil, err
	}
	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}
	return &result, nil
	// revive:enable:unexported-return
}

// Log returns the underlying log writer
func (c *config) Log() *logger.WrappedLogger {
	c.mutex.RLock()
	return c.log
}

// ──────────────────────────────────────────────────── I ──────────
//
//	:::::: S E T T E R : :  :   :    :     :        :          :
//
// ──────────────────────────────────────────────────────────────
func (c *config) SetDevelopmentMode() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.DevelopmentMode = true
}

func (c *config) SetNodeName(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.NodeName = value
	}
}
func (c *config) SetAPIAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.APIAddr = value
	}
}
func (c *config) SetMetricsPrefix(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.MetricsPrefix = value
	}
}
func (c *config) SetStatsiteAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.StatsiteAddr = value
	}
}
func (c *config) SetStatsdAddr(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value != "" {
		c.StatsdAddr = value
	}
}
func (c *config) SetPrometheusRetentionTime(value time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if value > 0 {
		c.PrometheusRetentionTime = value
	}
}

//
// ────────────────────────────────────────────────────────────── I ──────────
//   :::::: F I L E   L O A D E R : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────────
//

// ReadconfigPaths takes in an array of file (json)
// paths and returns a config struct.
// TODO: read configs in parallel with help of a waitgroup as calling defer in loop might lead to bugs
// nolint:gocognit,gocyclo // I can't think of any ways to lower complexity
func ReadconfigPaths(paths []string) (*config, error) { // revive:disable:unexported-return,defer
	result := &config{
		log: logger.DefaultWrappedLogger(string(logger.ErrorLevel)),
	}
	for _, p := range paths {
		path := filepath.Clean(p)
		f, err := os.Open(path)
		if err != nil {
			return nil, stacktrace.Propagate(err, "error reading '%s'", path)
		}
		// note(damoon) this may cause the application to open many
		// file descriptors which might cause the kernel to panic
		if f != nil {
			defer func() {
				err = f.Close()
				if err != nil {
					result.log.Error("cannot close file:%v", err)
				}
			}()
		}
		fi, err := f.Stat()
		if err != nil {
			return nil, stacktrace.Propagate(err, "error reading '%s'", path)
		}
		if !fi.IsDir() {
			dec := new(config) // nolint:staticcheck // SA4006 dec is used right away
			dec, err = Decodeconfig(f)
			if err != nil {
				return nil, stacktrace.Propagate(err, "error decoding '%s'", path)
			}
			result = Mergeconfig(result, dec)
			continue
		}
		contents, err := f.Readdir(-1)
		if err != nil {
			return nil, stacktrace.Propagate(err, "error reading '%s'", path)
		}
		sort.Sort(dirEnts(contents))
		for _, fi := range contents {
			if fi.IsDir() {
				continue
			}
			if !strings.HasSuffix(fi.Name(), ".json") {
				continue
			}
			subpath := filepath.Clean(
				filepath.Join(path, fi.Name()),
			)
			ff, err := os.Open(subpath)
			if err != nil {
				return nil, stacktrace.Propagate(err, "error reading '%s'", subpath)
			}
			if ff != nil {
				defer func() {
					err = ff.Close()
					if err != nil {
						result.log.Error("cannot close file:%v", err)
					}
				}()
			}
			config, err := Decodeconfig(ff)
			if err != nil {
				return nil, stacktrace.Propagate(err, "error decoding '%s'", subpath)
			}
			result = Mergeconfig(result, config)
		}
	}
	return result, nil
	// revive:enable:unexported-return,defer
}

// ─── SORT INTERFACE ─────────────────────────────────────────────────────────────
type dirEnts []os.FileInfo

// Len implements sort.Interface.
func (d dirEnts) Len() int { return len(d) }

// Less implements sort.Interface
func (d dirEnts) Less(i, j int) bool { return d[i].Name() < d[j].Name() }

// Swap implement sort.Interface
func (d dirEnts) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

// ────────────────────────────────────────────────────────────────────────────────
