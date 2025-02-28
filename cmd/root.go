/*
Copyright © 2025 Rohan Date rohan.s.date@icloud.com
*/
package cmd

// Lines in this file: 87

import (
	"fmt"
	"os"

	g "github.com/rsdate/rpkg/cmd/general"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	buildFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rpkg",
	Short: "A brief description of your application",
	Long: `Rpkg is a CLI tool for building and installing packages from a remote repository.
For example, you can build a package by running rpkg build /path/to/project
You can also install a package by running rpkg install mypackage 1.0.0`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&buildFile, "buildfile", "b", "", "set the rpkg.build.yaml path (default is PROJECTDIR/rpkg.build.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}

func addSubCommands() {
	rootCmd.AddCommand(g.InstallCmd)
	rootCmd.AddCommand(g.BuildCmd)
	rootCmd.AddCommand(g.RemoveCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if buildFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(buildFile)
	} else {
		// Find home directory.
		curdir, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".rpkg" (without extension).
		viper.AddConfigPath(curdir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("rpkg.build")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
