package cmd

import (
	"github.com/ondrejsika/master/version"
	_ "github.com/sikalabs/slr/cmd"
	slr_root "github.com/sikalabs/slr/cmd/root"
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/tergum/cmd"
	tergum_root "github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "master",
	Short: "master, " + version.Version,
	Long:  "master (" + version.Version + "): contains master branch of slu, slr, and tergum",
}

func init() {
	Cmd.AddCommand(slu_root.RootCmd)
	Cmd.AddCommand(slr_root.Cmd)
	Cmd.AddCommand(tergum_root.Cmd)
}

func Execute() {
	Cmd.Execute()
}
