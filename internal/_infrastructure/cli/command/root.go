package command

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal"
	"github.com/asiermarques/adrgen/internal/_infrastructure"
	"github.com/asiermarques/adrgen/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adrgen",
	Short: "A cli utility to create and manage Architecture Decision Records",
	Long: fmt.Sprintf(`
   ___     ___     ___     ___                   
  /   \   |   \   | _ \   / __|    ___    _ _    
  | - |   | |) |  |   /  | (_ |   / -_)  | ' \   
  |_|_|   |___/   |_|_\   \___|   \___|  |_||_|

O       o O       o O       o O       o O       o
| O   o | | O   o | | O   o | | O   o | | O   o |  
| | O | | | | O | | | | O | | | | O | | | | O | |
| o   O | | o   O | | o   O | | o   O | | o   O | 
o       O o       O o       O o       O o       O 

A cli utility to create and manage Architecture Decision Records
version: %s

`, internal.VERSION),
}

// Execute executes the root Command
func ExecuteRootCommand() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetConfig is used by the CLI Commands that need the project's configuration values
//
func GetConfig() (config.Config, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return config.Config{}, err
	}

	configManager := _infrastructure.CreateConfigFileManager(rootDirectory)
	config, err := configManager.Read()
	if err != nil {
		config := configManager.GetDefault()
		config.TargetDirectory = rootDirectory
		return config, fmt.Errorf("config file not found")
	}
	return config, nil
}

// MetaFlag slice for the meta param that could be used by the cli commands
//
var MetaFlag []string

func init() {
	rootCmd.AddCommand(CreateVersionCommand())
	rootCmd.AddCommand(CreateStatusChangeCommand())
	rootCmd.AddCommand(CreateListCommand())
	rootCmd.AddCommand(CreateInitCommand())
	rootCmd.AddCommand(CreateCreateCommand())
}
