package cmd

import (
	"github.com/ondrejsika/master/cmd/root"
	_ "github.com/ondrejsika/master/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
