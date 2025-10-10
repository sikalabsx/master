package main

import (
	_ "github.com/sikalabs/tergum/cmd"
	tergum_root "github.com/sikalabs/tergum/cmd/root"
)

func main() {
	tergum_root.Cmd.Execute()
}
