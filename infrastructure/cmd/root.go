package cmd

import (
	"fmt"
	"os"

	"github.com/asiermarques/adrgen/domain"
	"github.com/asiermarques/adrgen/infrastructure"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "adrgen",
	Short: "A cli utility to create and manage Architecture Decision Records",
	Long: `A cli utility to create and manage Architecture Decision Records
     ___     ___     ___     ___                   
    /   \   |   \   | _ \   / __|    ___    _ _    
    | - |   | |) |  |   /  | (_ |   / -_)  | ' \   
    |_|_|   |___/   |_|_\   \___|   \___|  |_||_|

  O       o O       o O       o O       o O       o
  | O   o | | O   o | | O   o | | O   o | | O   o |  
  | | O | | | | O | | | | O | | | | O | | | | O | |
  | o   O | | o   O | | o   O | | o   O | | o   O | 
  o       O o       O o       O o       O o       O 

`,
}

// Execute executes the root Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetConfig is used by the CLI Commands that need the project's configuration values
//
func GetConfig() (domain.Config, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return domain.Config{}, err
	}

	configManager := infrastructure.CreateConfigFileManager(rootDirectory)
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
