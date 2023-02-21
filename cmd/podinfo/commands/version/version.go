package version

import (
	"strings"

	buildver "github.com/da-moon/podinfo/build/go/version"
	cli "github.com/mitchellh/cli"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{
		UI: ui,
	}
	return c
}

type cmd struct {
	UI       cli.Ui
	help     string
	synopsis string
}

func (c *cmd) Run(_ []string) int {
	build := buildver.New()
	c.UI.Output(build.ToString())
	return 0
}
func (c *cmd) Synopsis() string {
	return strings.TrimSpace(c.synopsis)
}

func (c *cmd) Help() string {
	return strings.TrimSpace(c.help)
}

const synopsis = "podinfo immutable build information ."

const help = `
Usage: podinfo version

returns current release buildver.
`
