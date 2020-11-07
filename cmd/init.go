package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/application"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewInitCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "init",
		Short: "Initialize the ADRs working directory",
		Long: `Initialize the ADRs working directory`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			directory, err := os.Getwd()
			if err!=nil {
				fmt.Printf("an error ocurred listing the current directory %s\n", err)
				return
			}

			targetDirectory := filepath.Join(directory, args[0])
			if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
				if err := os.Mkdir(targetDirectory, 0644); err != nil {
					fmt.Printf("an error ocurred creating the directory %s\n", err)
					return
				}
			}

			meta, metaError := cmd.LocalFlags().GetStringSlice("meta")
			if metaError != nil {
				fmt.Printf("an error ocurred processing the meta parameter %s\n", metaError)
				return
			}
			for i, value := range meta {
				meta[i] = strings.TrimSpace(value)
			}

			if err := application.InitProject(targetDirectory, "adr_template.md", meta); err != nil {
				fmt.Printf("an error ocurred initializing the project in the target directory %s", targetDirectory)
				return
			}

		},
	}
	command.LocalFlags().StringSliceP("meta", "m", nil, "")
	return command
}

func init() {
	rootCmd.AddCommand(NewCreateCmd())
}
