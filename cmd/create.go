package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/application"
	"os"

	"github.com/spf13/cobra"
)

func NewCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new ADR File in the current directory",
		Long: `Create a new ADR File in the current directory`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			directory, err := os.Getwd()
			if err!=nil {
				fmt.Printf("an error ocurred listing the current directory %s\n", err)
				return
			}
			_, creationError := application.CreateADRFile(args[0], directory, os.Getenv("ADRGEN_TEMPLATE"))
			if creationError!=nil {
				fmt.Println(err)
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(NewCreateCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
