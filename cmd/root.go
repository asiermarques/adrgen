package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/asiermarques/adrgen/adr"
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetConfig(directory string) (adr.Config, error) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		return adr.Config{}, err
	}
	directory = filepath.Join(rootDirectory, directory)
	config, err := adr.GetConfig(directory)
	if err != nil {
		config.TargetDirectory = rootDirectory
	} else {
		config.TargetDirectory = filepath.Join(rootDirectory, config.TargetDirectory)
		config.TemplateFilename = filepath.Join(rootDirectory, config.TemplateFilename)
	}
	return config, err
}
