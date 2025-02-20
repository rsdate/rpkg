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
	if code, err := re.Build(projectPath, f, false); err != nil {
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
	Short: "Builds the package in the project path specified",
	Long: `Builds the package in the project path specified.
For example: if you have a project in /home/user/MyProject, you can build it 
by running rpkg build /home/user/MyProject`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Your package is being built at " + args[0] + ". Would you like to continue? [Y/n]")
		fmt.Scan(&input)
		if input == "Y" {
			if _, err := buildPackage(args[0]); err != nil {
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
