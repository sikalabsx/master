package main

import (
	_ "github.com/sikalabs/slr/cmd"
	slr_root "github.com/sikalabs/slr/cmd/root"
)

func main() {
	slr_root.Cmd.Execute()
}
