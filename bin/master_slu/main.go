package main

import (
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
)

func main() {
	slu_root.RootCmd.Execute()
}
