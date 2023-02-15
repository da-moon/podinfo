package version

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

// Build information. Populated at build-time.
var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	Toolchain = runtime.Version()
)

// BuildInformation helps with working with immutable build information.
//
//go:generate gomodifytags -override -file $GOFILE -struct BuildInformation -add-tags json,mapstructure -w -transform snakecase
type BuildInformation struct {
	Version   string `json:"version" mapstructure:"version"`
	Revision  string `json:"revision" mapstructure:"revision"`
	Branch    string `json:"branch" mapstructure:"branch"`
	BuildDate string `json:"build_date" mapstructure:"build_date"`
	Toolchain string `json:"toolchain" mapstructure:"toolchain"`
	BuildUser string `json:"build_user" mapstructure:"build_user"`
}

// New returns a new BuildInformation.
func New() *BuildInformation {
	result := &BuildInformation{
		Version:   Version,
		Revision:  Revision,
		Branch:    Branch,
		BuildDate: BuildDate,
		Toolchain: Toolchain,
		BuildUser: BuildUser,
	}
	return result
}

// ToString returns version information.
func (b *BuildInformation) ToString() string {
	var result bytes.Buffer
	if b.Version != "" {
		fmt.Fprintf(&result, "\nVersion      :  %s", strings.TrimPrefix(b.Version, "v"))
	}
	if b.Revision != "" {
		fmt.Fprintf(&result, "\nRevision     :  %s", b.Revision)
	}
	if b.Branch != "" {
		fmt.Fprintf(&result, "\nBranch       :  %s", b.Branch)
	}
	if b.BuildUser != "" {
		fmt.Fprintf(&result, "\nBuild User   :  %s", b.BuildUser)
	}
	if b.BuildDate != "" {
		fmt.Fprintf(&result, "\nBuild Date   :  %s", b.BuildDate)
	}
	if b.Toolchain != "" {
		fmt.Fprintf(&result, "\nToolchain    :  %s", strings.TrimPrefix(b.Toolchain, "go"))
	}
	return result.String()
}

// Info returns version, branch and revision information.
func (b *BuildInformation) Info() string {
	var result bytes.Buffer
	if b.Version != "" {
		fmt.Fprintf(&result, "version=%s, ", strings.TrimPrefix(b.Version, "v"))
	}
	if b.Branch != "" {
		fmt.Fprintf(&result, "branch=%s, ", b.Branch)
	}
	if b.Revision != "" {
		fmt.Fprintf(&result, "revision=%s", b.Revision)
	}
	return fmt.Sprintf("(%s)", result.String())
}

// BuildContext returns toolchain, buildUser and buildDate information.
func (b *BuildInformation) BuildContext() string {
	var result bytes.Buffer
	if b.Toolchain != "" {
		fmt.Fprintf(&result, "go=%s, ", strings.TrimPrefix(b.Toolchain, "go"))
	}
	if b.BuildUser != "" {
		fmt.Fprintf(&result, "user=%s, ", b.BuildUser)
	}
	if b.BuildDate != "" {
		fmt.Fprintf(&result, "date=%s", b.BuildDate)
	}
	return fmt.Sprintf("(%s)", result.String())
}
