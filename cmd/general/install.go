/*
Copyright © 2024 Rohan Date <rohan.s.date@icloud.com>
*/
package general

import (
	// "errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a package from the environment variable RPKG_MIRROR",
	Long: `Install installs a package from the environment variable RPKG_MIRROR.
For example: rpkg install mypackage 1.0.0 will install mypackage version 1.0.0 from
https://RPKG_MIRROR/projects/mypackage-1.0.0.tar.gz`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := args[0] + "-" + args[1]
		projectPath := dirName + ".tar.gz"
		downloadPath := os.Getenv(download_dir) + "/" + projectPath
		defaultMirror := os.Getenv(mirror)
		if defaultMirror == "" {
			fmt.Fprintln(os.Stdout, []any{"warning: environment variable RPKG_MIRROR not set.\nReverting to default mirror... "}...)
			checkErr(os.Getenv(panic_mode), "package installation failed", func() (int, error) {
				_, err := InstallPackage(downloadPath, projectPath, dirName, os.Getenv(panic_mode))
				return returnErr(err)
			})
		} else {
			checkErr(os.Getenv(panic_mode), "package installation failed", func() (int, error) {
				_, err := InstallPackage(downloadPath, projectPath, dirName, os.Getenv(panic_mode))
				return returnErr(err)
			})
		}

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
