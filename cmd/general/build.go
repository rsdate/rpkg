/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package general

import (
	"errors"
	"fmt"
	"os"

	re "github.com/rsdate/rpkgengine/rpkgengine"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	viper_instance = viper.GetViper()
	input          string
)

func buildPackage(projectPath string) (int, error) {
	os.Chdir(projectPath + "/Package")
	fmt.Print("Building package... ")
	if code, err := re.Build(projectPath, f); err != nil {
		fmt.Println("Build failed.")
		return code, errors.New("build failed")
	} else {
		fmt.Println("Build successful.")
		return code, nil
	}
}

// buildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if dir, err := os.Getwd(); err != nil {
			fmt.Println("Could not get current directory")
			os.Exit(1)
		} else {
			fmt.Println("Your package is being built at " + dir + ". Would you like to continue? [Y/n]")
			fmt.Scan(&input)
			if input == "Y" {
				if _, err := buildPackage(dir); err != nil {
					fmt.Println("Build failed.")
					os.Exit(1)
				} else {
					fmt.Println("Build successful.")
					os.Exit(0)
				}
			} else {
				fmt.Println("Build aborted.")
				os.Exit(0)
			}
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
