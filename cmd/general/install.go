/*
Copyright © 2024 Rohan Date <rohan.s.date@icloud.com>
*/
package general

import (
	//	"fmt"
	"errors"
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
	if resp.StatusCode != http.StatusOK {
		var err error = errors.New("server returned bad status code")
		return 1, err
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
	Short: "Installs a package from " + mirror,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultMirror := os.Getenv(mirror)
		if defaultMirror == "" {
			panic("fatal: default mirror not set.\nConsider setting " + mirror + " and you will not see this error again.")
		}
		fullName := "https://" + defaultMirror + "/projects/" + args[0] + "-" + args[1] + ".tar.gz"
		fmt.Printf("The package path is: %s. Would you like to proceed with the installation? [Y or n]", fullName)
		fmt.Scan(conf)
		if conf == "Y" {
			code, err := DownloadFile("./"+args[0]+"-"+args[1]+".tar.gz", fullName)
			if code != 0 && err != nil {
				panic(fmt.Errorf("fatal: Unable to download package. Please check to see whether your package actually exists. Error Message: %s", err))
			}
		} else if conf == "n" {
			fmt.Println("Installation aborted.")
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
