package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/application"
	"strconv"

	"github.com/spf13/cobra"
)

func NewStatusChangeCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "status [new status] [ADR ID]",
		Short: "Update the status in a ADR File",
		Long: `Update the status in a ADR File`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			config, err := GetConfig("")
			if err != nil {
				fmt.Printf("config file not found, working in the %s directory\n", config.TargetDirectory)
			} else {
				fmt.Printf("config file found, working in the %s directory\n", config.TargetDirectory)
			}


			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			filename, updateError := application.ChangeADRStatus(id, args[1], config)
			if updateError!=nil {
				fmt.Println(updateError)
				return
			}
			fmt.Println(fmt.Sprintf("%s updated with status %s\n", filename, args[1]))
		},
	}
	command.Example = "adrgen status 9 accepted"
	return command
}

func init() {
	rootCmd.AddCommand(NewStatusChangeCmd())
}
