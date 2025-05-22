package root

import (
	"github.com/ondrejsika/master/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "master",
	Short: "master, " + version.Version,
}
