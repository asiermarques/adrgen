package cmd

import (
	"../application"
	"fmt"
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

			targetDirectory := args[0]
			currentDirectory, err := os.Getwd()
			if err!=nil {
				fmt.Printf("an error ocurred listing the current directory %s\n", err)
				return
			}

			targetDirectoryAbs := filepath.Join(currentDirectory, targetDirectory)
			if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
				if err := os.MkdirAll(targetDirectoryAbs, os.ModePerm); err != nil {
					fmt.Printf("an error ocurred creating the target directory %s\n", err)
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
				fmt.Printf("an error ocurred initializing the project in the target directory %s, %s", targetDirectory, err)
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
