package command

import (
	"fmt"
	infrastructure2 "github.com/asiermarques/adrgen/internal/_infrastructure"
	"github.com/asiermarques/adrgen/internal/adr"
	"github.com/asiermarques/adrgen/internal/template"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// SupersededADRId represents the ID for a superseded ADR File
//
var SupersededADRId int

// AmendedADRId represents the ID for a amended ADR File
//
var AmendedADRId int

// CreateCreateCommand creates the 'create' CLI Command related to the ADR file creation
//
func CreateCreateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "create [the ADR title]",
		Short: "Create a new ADR File in the current directory",
		Long:  `Create a new ADR File in the current directory, you can add meta parameters for decisions tracing`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			config, err := GetConfig()
			if err != nil {
				if config.TargetDirectory != "" {
					fmt.Printf("error creating file: %s", err)
				}

				fmt.Printf(
					"config file not found, working in the %s directory\n",
					config.TargetDirectory,
				)
			} else {
				fmt.Printf("config file found, working in the %s directory\n", config.TargetDirectory)
			}

			meta, metaError := cmd.LocalFlags().GetStringSlice("meta")
			if metaError != nil {
				fmt.Printf("an error occurred processing the meta parameter %s\n", metaError)
				return
			}
			for i, value := range meta {
				meta[i] = strings.TrimSpace(value)
			}
			config.MetaParams = append(config.MetaParams, meta...)

			supersedesADRId, supersedesError := cmd.LocalFlags().GetInt("supersedes")
			if supersedesError != nil {
				fmt.Printf(
					"an error occurred processing the supersedes parameter %s\n",
					supersedesError,
				)
				return
			}
			amendsADRId, amendsErr := cmd.LocalFlags().GetInt("amends")
			if amendsErr != nil {
				fmt.Printf("an error occurred processing the amends parameter %s\n", amendsErr)
				return
			}

			currentTime := time.Now()
			date := currentTime.Format("2006-01-02")

			adrWriter := infrastructure2.CreateFileADRWriter(config.TargetDirectory)
			templateService := template.CreateService(
				infrastructure2.CreateCustomTemplateContentFileReader(config),
			)

			filename, creationError := adr.CreateFile(
				date,
				args[0],
				meta,
				supersedesADRId,
				amendsADRId,
				config,
				infrastructure2.CreateADRDirectoryRepository(config.TargetDirectory),
				adrWriter,
				templateService,
				adr.CreateRelationsManager(
					templateService,
					adr.CreateStatusManager(config),
				),
			)
			if creationError != nil {
				fmt.Println(creationError)
				return
			}
			fmt.Println(fmt.Sprintf("%s created\n", filename))
		},
	}

	command.Flags().IntVarP(&SupersededADRId, "supersedes", "s", 0, "")
	command.Flags().IntVarP(&AmendedADRId, "amends", "a", 0, "")
	command.Flags().StringSliceVarP(&MetaFlag, "meta", "m", []string{}, "")

	command.Example = "adrgen create \"Using ADR to record and maintain decisions records\""

	return command
}
