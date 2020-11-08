package cmd

import (
	"fmt"
	"github.com/asiermarques/adrgen/adr"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

//var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adrgen",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetConfig(directory string) (adr.Config, error) {
	rootDirectory, err := os.Getwd()
	if err!=nil {
		return adr.Config{}, err
	}
	directory = filepath.Join(rootDirectory, directory)
	config, err := adr.GetConfig(directory)
	if err != nil {
		config.TargetDirectory = rootDirectory
	}else{
		config.TargetDirectory = filepath.Join(rootDirectory, config.TargetDirectory)
		config.TemplateFilename = filepath.Join(rootDirectory, config.TemplateFilename)
	}
	return config, err
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.adrgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	/*if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".adrgen" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".adrgen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}*/
}
