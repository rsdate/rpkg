/*
Copyright © 2024 Rohan Date <rohan.s.date@icloud.com>
*/
package general

import (
	//	"fmt"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	mirror string = "RPKG_MIRROR"
)

func downloadFile(filepath string, url string) int {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return 1
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return 1
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return 1
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return 1
	}

	return 0
}

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a package from " + mirror,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultMirror := os.Getenv(mirror)
		if defaultMirror == "" {
			panic("fatal: default mirror not set.\nConsider setting " + mirror + " and you will not see this error again.")
		}
		fullName := "https://www." + defaultMirror + "/" + args[0] + "-" + args[1] + ".tar.gz"
		code := downloadFile("./"+args[0]+"-"+args[1]+".tar.gz", fullName)
		//		if code != 0 && err != nil {
		//			panic("fatal: internal error [ERROR 1000]: downloadFile exited with status code 1. Unable to download package.")
		//		}
		fmt.Printf("Code: %d", code)

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
