/*
Copyright © 2024 Rohan Date <rohan.s.date@icloud.com>
*/
package general

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rsdate/rpkgengine/rpkgengine"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	mirror       string = "RPKG_MIRROR"
	conf         string
	download_dir string = "DOWNLOAD_DIR"
)

func initVars(viper_instance *viper.Viper) rpkgengine.RpkgBuildFile {
	name := viper_instance.Get("name").(string)
	version := viper_instance.Get("version").(string)
	revision := viper_instance.Get("revision").(int)
	authors := viper_instance.Get("authors").([]interface{})
	deps := viper_instance.Get("deps").([]interface{})
	buildDeps := viper_instance.Get("build_deps").([]interface{})
	buildWith := viper_instance.Get("build_with").(string)
	buildCommands := viper_instance.Get("build_commands").([]interface{})
	f := rpkgengine.RpkgBuildFile{
		Name:          name,
		Version:       version,
		Revision:      revision,
		Authors:       authors,
		Deps:          deps,
		BuildDeps:     buildDeps,
		BuildWith:     buildWith,
		BuildCommands: buildCommands,
	}
	return f
}

func DownloadPackage(filepath string, url string) (int, error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, []any{"error: could not create file"}...)
		return 1, err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, []any{"error: could not get data from server"}...)
		return 1, err
	}
	defer resp.Body.Close()

	// Check server response
	switch code := resp.StatusCode; code {
	case http.StatusNotFound:
		fmt.Fprintln(os.Stderr, []any{"error: server did not find file"}...)
		return 1, err
	case http.StatusForbidden:
		fmt.Fprintln(os.Stderr, []any{"error: server did not allow permission to access the resource"}...)
		return 1, err
	case http.StatusUnauthorized:
		fmt.Fprintln(os.Stderr, []any{"error: server did not allow permission to access the resource"}...)
		return 1, err
	case http.StatusInternalServerError:
		fmt.Fprintln(os.Stderr, []any{"error: server encountered an internal error"}...)
		return 1, err
	case http.StatusServiceUnavailable:
		fmt.Fprintln(os.Stderr, []any{"error: server is currently unavailable"}...)
		return 1, err
	case http.StatusFound:
		return 0, nil
	default:
		fmt.Fprintln(os.Stderr, []any{"error: server returned an unexpected status code"}...)

	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, []any{"error: could not write to file"}...)
		return 1, err
	}

	return 0, nil
}

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
				code, err := DownloadPackage(downloadPath, fullName)
				if code != 0 && err != nil {
					panic(fmt.Errorf("fatal: Unable to download package. Please check to see whether your package actually exists. Error Message: %s", err))
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
