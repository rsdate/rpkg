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
		downloadPath := os.Getenv(download_dir) + "/" + args[0] + "-" + args[1] + ".tar.gz"
		projectPath := args[0] + "-" + args[1] + ".tar.gz"
		defaultMirror := os.Getenv(mirror)
		if defaultMirror == "" {
			fmt.Fprintln(os.Stdout, []any{"warning: environment variable RPKG_MIRROR not set.\nReverting to default mirror..."}...)
			code, err := DownloadPackage(downloadPath, "https://rsdate.github.io/projects/"+projectPath)
			if code != 0 && err != nil {
				panic(fmt.Errorf("fatal: Unable to download package. Please check to see whether your package actually exists. Error Message: %s", err))
			}
		} else {
			InstallPackage(downloadPath, projectPath, args[0]+"-"+args[1])
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
