/*
Copyright © 2024 Rohan Date <rohan.s.date@icloud.com>
*/
package general

import (
	// "errors"
	"fmt"
	"os"
	"os/exec"

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
			fullName := "https://" + defaultMirror + "/projects/" + projectPath
			fmt.Fprintf(os.Stdout, "The package path on the mirror is %s and it will download to %s.\nWould you like to proceed with the installation? [Y or n]", []any{projectPath, downloadPath}...)
			fmt.Scan(&conf)
			if conf == "Y" {
				fmt.Print("Downloading package... ")
				code, err := DownloadPackage(downloadPath, fullName)
				if code != 0 && err != nil {
					panic(fmt.Errorf("fatal: Unable to download package. Please check to see whether your package actually exists. Error Message: %s", err))
				}
				fmt.Println("Package downloaded successfully.")
				fmt.Print("Unziping package... ")
				cmd := exec.Command("tar", "-xzf", projectPath)
				cmd.Stdout = nil
				err = cmd.Run()
				if err != nil {
					fmt.Fprint(os.Stderr, []any{"error: could not unzip package"}...)
					os.Exit(1)
				}
				fmt.Println("Package unziped successfully.")
				fmt.Print("Building package... ")
				os.Chdir(args[0] + "-" + args[1])
				if _, err := BuildPackage("."); err != nil {

				}
				fmt.Println("Installation completed! 🎉")
			} else if conf == "n" {
				fmt.Fprintln(os.Stdout, []any{"Installation aborted."}...)
				os.Exit(0)
			}
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
