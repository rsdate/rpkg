/*
Copyright © 2024 Rohan Date <rohan.s.date@icloud.com>
*/
package general

import (
	//	"fmt"
	// "errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	mirror string = "RPKG_MIRROR"
	conf   string
)

func DownloadFile(filepath string, url string) (int, error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return 1, err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return 1, err
	}
	defer resp.Body.Close()

	// Check server response
	switch code := resp.StatusCode; code {
	case http.StatusNotFound:
		fmt.Fprintln(os.Stdout, []any{"error: server did not find file"}...)
		return 1, err
	case http.StatusForbidden:
		fmt.Fprintln(os.Stdout, []any{"error: server did not allow permission to access the resource"}...)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return 1, err
	}

	return 0, nil
}

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a package from the environment variable RPKG_MIRROR",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		downloadPath := "./" + args[0] + "-" + args[1] + ".tar.gz"
		projectPath := args[0] + "-" + args[1] + ".tar.gz"
		defaultMirror := os.Getenv(mirror)
		if defaultMirror == "" {
			fmt.Fprintln(os.Stdout, []any{"warning: environment variable RPKG_MIRROR not set.\nReverting to default mirror..."}...)
			code, err := DownloadFile(downloadPath, "https://rsdate.github.io/projects/"+projectPath)
			if code != 0 && err != nil {
				panic(fmt.Errorf("fatal: Unable to download package. Please check to see whether your package actually exists. Error Message: %s", err))
			}
		}
		fullName := "https://" + defaultMirror + "/projects/" + projectPath
		fmt.Fprintf(os.Stdout, "The package path on the mirror is %s and it will download to %s.\nWould you like to proceed with the installation? [Y or n]", []any{projectPath, downloadPath}...)
		fmt.Scan(&conf)
		if conf == "Y" {
			code, err := DownloadFile(downloadPath, fullName)
			if code != 0 && err != nil {
				panic(fmt.Errorf("fatal: Unable to download package. Please check to see whether your package actually exists. Error Message: %s", err))
			}
		} else if conf == "n" {
			fmt.Fprintln(os.Stdout, []any{"Installation aborted."}...)
			os.Exit(0)
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
