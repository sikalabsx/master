package main

import (
	_ "github.com/sikalabs/mon/cmd"
	mon_root "github.com/sikalabs/mon/cmd/root"
)

func main() {
	mon_root.Cmd.Execute()
}
