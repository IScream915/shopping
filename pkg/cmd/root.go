package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	FlagConfPath = "config"
)

type RootCmd struct {
	Command     cobra.Command
	processName string
}

func NewRootCmd(processName string) *RootCmd {
	rootCmd := &RootCmd{
		processName: processName,
	}
	cmd := cobra.Command{
		Use:           fmt.Sprintf("Start %s application", processName),
		Long:          fmt.Sprintf(`Start %s `, processName),
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	cmd.Flags().StringP(FlagConfPath, "c", "", "path of config file")

	rootCmd.Command = cmd
	return rootCmd
}
