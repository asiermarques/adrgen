package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/application"
	"strings"

	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create a new ADR File in the current directory",
		Long: `Create a new ADR File in the current directory, you can add meta parameters for decisions tracing`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			config, err := GetConfig("")
			if err != nil {
				fmt.Printf("config file not found, working in the %s directory\n", config.TargetDirectory)
			} else {
				fmt.Printf("config file found, working in the %s directory\n", config.TargetDirectory)
			}


			meta, metaError := cmd.LocalFlags().GetStringSlice("meta")
			if metaError != nil {
				fmt.Printf("an error ocurred processing the meta parameter %s\n", metaError)
				return
			}
			for i, value := range meta {
				meta[i] = strings.TrimSpace(value)
			}
			config.MetaParams = append(config.MetaParams, meta...)

			filename, creationError := application.CreateADRFile(args[0], config)
			if creationError!=nil {
				fmt.Println(creationError)
				return
			}
			fmt.Println(fmt.Sprintf("%s created\n", filename))
		},
	}
	command.LocalFlags().StringSliceP("meta", "m", nil, "")
	return command
}

func init() {
	rootCmd.AddCommand(NewCreateCmd())
}
