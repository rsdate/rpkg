package general

// This file contains all the excess functions and variables from the general package

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	re "github.com/rsdate/rpkgengine/rpkgengine"
	"github.com/spf13/viper"
)

var (
	mirror         string = "RPKG_MIRROR"
	conf           string
	download_dir   string = "DOWNLOAD_DIR"
	viper_instance        = viper.GetViper()
	input          string
)

func initVars(viper_instance *viper.Viper) re.RpkgBuildFile {
	name := viper_instance.Get("name").(string)
	version := viper_instance.Get("version").(string)
	revision := viper_instance.Get("revision").(int)
	authors := viper_instance.Get("authors").([]interface{})
	deps := viper_instance.Get("deps").([]interface{})
	buildDeps := viper_instance.Get("build_deps").([]interface{})
	buildWith := viper_instance.Get("build_with").(string)
	buildCommands := viper_instance.Get("build_commands").([]interface{})
	f := re.RpkgBuildFile{
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
	case http.StatusOK:
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

func BuildPackage(projectPath string) (int, error) {
	if viper_instance == nil {
		fmt.Println("No configuration file found. If your build file is located in a different directory, please specify the path using the --buildfile flag.")
		return 1, errors.New("no configuration file found")
	}
	f := initVars(viper_instance)
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
