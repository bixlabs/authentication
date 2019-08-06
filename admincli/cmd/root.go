package cmd

import (
	"fmt"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/spf13/cobra"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &AuthCommand{
	Command: cobra.Command{
		Use:   "admincli",
		Short: "Command line utility to manage users.",
		Long:  `Command line utility to create/update/delete/find users and reset their passwords.`,
	},
}

type AuthCommand struct {
	cobra.Command
	authenticator interactors.Authenticator
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(authenticator interactors.Authenticator) {
	// TODO: use constructor
	rootCmd.authenticator = authenticator

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.admincli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".admincli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".admincli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
