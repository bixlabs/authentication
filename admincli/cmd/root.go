package cmd

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/tools"
	"github.com/spf13/cobra"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
// TODO: use a constructor and rename
var rootCmd = &AuthCommand{
	Command: cobra.Command{
		Use:          "admincli",
		Short:        "Command line utility to manage users.",
		Long:         `Command line utility to create/update/delete/find users and reset their passwords.`,
		SilenceUsage: true,
	},
}

type AuthCommand struct {
	Command       cobra.Command
	Authenticator interactors.Authenticator
}

func (ac *AuthCommand) setAuth(auth interactors.Authenticator) {
	ac.Authenticator = auth
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(authenticator interactors.Authenticator) {
	// TODO: use constructor
	rootCmd.Authenticator = authenticator

	if err := rootCmd.Command.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	rootCmd.Command.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.admincli.yaml)")
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
			tools.Log().Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".admincli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".admincli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		tools.Log().Info("Using config file:", viper.ConfigFileUsed())
	}
}
