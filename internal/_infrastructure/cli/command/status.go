package command

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/_infrastructure"
	"github.com/asiermarques/adrgen/internal/adr"
	"strconv"

	"github.com/spf13/cobra"
)

// CreateStatusChangeCommand creates the 'status' CLI Command related to the ADR file status change use case
//
func CreateStatusChangeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "status [ADR ID] [new status]",
		Short: "Update the status in a ADR File",
		Long:  `Update the status in a ADR File`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := GetConfig()
			if err != nil {
				fmt.Printf(
					"config file not found, working in the %s directory\n",
					config.TargetDirectory,
				)
			} else {
				fmt.Printf("config file found, working in the %s directory\n", config.TargetDirectory)
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			filename, updateError := adr.ChangeStatus(
				id,
				args[1],
				_infrastructure.CreateADRDirectoryRepository(config.TargetDirectory),
				adr.CreateStatusManager(config),
				_infrastructure.CreateFileADRWriter(config.TargetDirectory),
			)
			if updateError != nil {
				fmt.Println(updateError)
				return
			}
			fmt.Println(fmt.Sprintf("%s updated with status %s\n", filename, args[1]))
		},
	}
	command.Example = "adrgen status 9 accepted"
	return command
}
