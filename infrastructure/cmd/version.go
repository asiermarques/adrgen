package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NewVersionCmd creates the 'version' CLI Command to get the version of ADRgen
//
func NewVersionCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Shows the ADRgen version",
		Long:  `Shows the ADRgen version`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("version: %s", VERSION))
		},
	}
	command.Example = "adrgen version"
	return command
}

func init() {
	rootCmd.AddCommand(NewVersionCmd())
}
