package main

import (
	_ "github.com/sikalabs/slr/cmd"
	slr_root "github.com/sikalabs/slr/cmd/root"
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
	slu_tools "github.com/sikalabs/slu/cmd/tools"
	_ "github.com/sikalabs/tergum/cmd"
	tergum_root "github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "master",
	Short: "master: contains master branch of slu, slr, and tergum",
}

func init() {
	Cmd.AddCommand(slu_root.RootCmd)
	Cmd.AddCommand(slr_root.Cmd)
	Cmd.AddCommand(tergum_root.Cmd)
	Cmd.AddCommand(slu_tools.Cmd)
}

func Execute() {
	Cmd.Execute()
}

func main() {
	Execute()
}
