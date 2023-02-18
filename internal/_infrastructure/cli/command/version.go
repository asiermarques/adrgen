package command

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal"

	"github.com/spf13/cobra"
)

// CreateVersionCommand creates the 'version' CLI Command to get the version of ADRgen
//
func CreateVersionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Shows the ADRgen version",
		Long:  `Shows the ADRgen version`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("version: %s", internal.VERSION))
		},
	}
	command.Example = "adrgen version"
	return command
}
