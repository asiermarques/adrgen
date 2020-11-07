package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/application"
	"os"
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
			directory, err := os.Getwd()
			if err!=nil {
				fmt.Printf("an error ocurred listing the current directory %s\n", err)
				return
			}

			meta, metaError := cmd.LocalFlags().GetStringSlice("meta")
			if metaError != nil {
				fmt.Printf("an error ocurred processing the meta parameter %s\n", metaError)
				return
			}
			for i, value := range meta {
				meta[i] = strings.TrimSpace(value)
			}

			_, creationError := application.CreateADRFile(args[0], directory, os.Getenv("ADRGEN_TEMPLATE"), meta)
			if creationError!=nil {
				fmt.Println(err)
			}
		},
	}
	command.LocalFlags().StringSliceP("meta", "m", nil, "")
	return command
}

func init() {
	rootCmd.AddCommand(NewCreateCmd())
}
