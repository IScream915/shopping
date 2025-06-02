package main

import (
	"shopping/pkg/common/cmd"
	"shopping/pkg/program"
)

func main() {
	if err := cmd.NewApiCmd().Exec(); err != nil {
		program.ExitWithError(err)
	}
}
