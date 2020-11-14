package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/asiermarques/adrgen/application"
	"github.com/asiermarques/adrgen/infrastructure"
	"github.com/spf13/cobra"
)

// NewInitCmd creates the 'init' CLI Command related to the project initialization
//
func NewInitCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "init [directory]",
		Short: "Initialize the ADRs working directory",
		Long:  `Initialize the ADRs working directory`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			targetDirectory := args[0]
			if _, err := os.Stat(targetDirectory); os.IsNotExist(err) {
				if err := os.MkdirAll(targetDirectory, os.ModePerm); err != nil {
					fmt.Printf("an error occurred creating the target directory %s\n", err)
					return
				}
			}

			meta, metaError := cmd.LocalFlags().GetStringSlice("meta")
			if metaError != nil {
				fmt.Printf("an error occurred processing the meta parameter %s\n", metaError)
				return
			}
			for i, value := range meta {
				meta[i] = strings.TrimSpace(value)
			}

			configManager := infrastructure.CreateConfigFileManager(".")
			config := configManager.GetDefault()
			config.TargetDirectory = targetDirectory
			config.TemplateFilename = filepath.Join(targetDirectory, "adr_template.md")
			config.MetaParams = meta

			if err := application.InitProject(
				config,
				configManager,
				infrastructure.CreateTemplateFileWriter(config)); err != nil {
				fmt.Printf(
					"an error occurred initializing the project in the target directory %s, %s",
					targetDirectory,
					err,
				)
				return
			}
		},
	}
	command.Flags().StringSliceVarP(&MetaFlag, "meta", "m", []string{}, "")
	command.Example = "adrgen init \"docs/adrs\""
	return command
}

func init() {
	rootCmd.AddCommand(NewInitCmd())
}
