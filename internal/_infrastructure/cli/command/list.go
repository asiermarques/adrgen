package command

import (
	"fmt"
	infrastructure2 "github.com/asiermarques/adrgen/internal/_infrastructure"
	"github.com/asiermarques/adrgen/internal/adr"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// FilterQuery represents the query string for a ADR filter
//
var FilterQuery string

// CreateListCommand creates the 'list' CLI Command that shows all the ADRs
//
func CreateListCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List the ADR files",
		Long:  `List the ADR files`,
		Args:  cobra.ExactArgs(0),
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

			rawFilterQuery, metaError := cmd.LocalFlags().GetString("filter")
			if metaError != nil {
				fmt.Printf("an error occurred processing the meta parameter %s\n", metaError)
				return
			}

			var repository = infrastructure2.CreateADRDirectoryRepository(config.TargetDirectory)
			var adrFiles []adr.ADR
			if rawFilterQuery != "" {
				filterParams, err := infrastructure2.ParseFilterParams(rawFilterQuery)
				if err != nil {
					fmt.Printf("an error occurred %s", err)
					return
				}

				files, err := repository.Query(filterParams)
				if err != nil {
					fmt.Printf("error listing the adr files filtered: %s", err)
					return
				}
				adrFiles = files
			} else {
				files, err := repository.FindAll()
				if err != nil {
					fmt.Printf("error listing the adr files: %s", err)
					return
				}
				adrFiles = files
			}

			table := createTable()
			for _, adr := range adrFiles {
				addADRRow(adr, table)
			}

			fmt.Println("")
			table.Print()
			fmt.Println("")
		},
	}
	command.Flags().StringVarP(&FilterQuery, "filter", "f", "", "adrgen list -f status=accepted")
	command.Example = "adrgen list"
	return command
}

func addADRRow(adr adr.ADR, table table.Table) {
	table.AddRow(adr.Title(), adr.Status(), adr.Date(), adr.ID(), adr.Filename().Value())
}

func createTable() table.Table {
	tbl := table.New("Title", "Status", "Date", "ID", "Filename")
	return tbl
}
