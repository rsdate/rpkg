package general

// This file contains all the excess functions and variables from the general package

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	re "github.com/rsdate/rpkgengine/rpkgengine"
	"github.com/spf13/viper"
)

var (
	mirror         string       = "RPKG_MIRROR"
	download_dir   string       = "RPKG_DOWNLOAD_DIR"
	panic_mode     string       = "RPKG_PANICMODE"
	viper_instance *viper.Viper = viper.GetViper()
	conf           string
	input          string
	eM             []string = []string{
		// DownloadPackage error messages
		"could not create the file",
		"could not get data from server",
		"server did not find file",
		"server did not allow permission to access the resource",
		"user is not authorized to access the resource",
		"server encountered an internal error",
		"server is currently unavailable",
		"could not write to the file",
		// BuildPackage error messages
		"no configuartion file found",
		"build failed",
		// InstallPackage error messages
		"package could not be downloaded",
		"package could not be unzipped",
		"package could not be built",
	}
)

func createErr(message string) error {
	return fmt.Errorf("%s", []any{message}...)
}

func returnErr(err error) (int, error) {
	if err != nil {
		return 1, err
	}
	return 0, nil
}

func checkErr(panicMode string, errMessage string, y func() (int, error)) (int, error) {
	code, err := y()
	if code != 0 && err != nil {
		fmt.Fprint(os.Stderr, []any{errMessage}...)
		if panicMode == "true" {
			panic(fmt.Errorf("error: %s. error message: %v", []any{errMessage, err}...))
		} else if panicMode == "false" {
			return 1, fmt.Errorf("error: %s. error message: %v", []any{errMessage, err}...)
		}
	}
	return 0, nil
}

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

func DownloadPackage(filepath string, url string, panicMode string) (int, error) {
	// Create the file
	out, err := os.Create(filepath)
	checkErr(panicMode, eM[0], func() (int, error) {
		return returnErr(err)
	})
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	checkErr(panicMode, eM[1], func() (int, error) {
		return returnErr(err)
	})
	defer resp.Body.Close()

	// Check server response
	switch code := resp.StatusCode; code {
	case http.StatusNotFound:
		checkErr(panicMode, eM[2], func() (int, error) {
			return returnErr(createErr(eM[2]))
		})
	case http.StatusForbidden:
		checkErr(panicMode, eM[3], func() (int, error) {
			return returnErr(createErr(eM[3]))
		})
	case http.StatusUnauthorized:
		checkErr(panicMode, eM[4], func() (int, error) {
			return returnErr(createErr(eM[4]))
		})
	case http.StatusInternalServerError:
		checkErr(panicMode, eM[5], func() (int, error) {
			return returnErr(createErr(eM[5]))
		})
	case http.StatusServiceUnavailable:
		checkErr(panicMode, eM[6], func() (int, error) {
			return returnErr(createErr(eM[6]))
		})
	case http.StatusOK:
		return 0, nil
	}
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	checkErr(panicMode, eM[7], func() (int, error) {
		return returnErr(createErr(eM[7]))
	})
	return 0, nil
}

func BuildPackage(projectPath string, panicMode string) (int, error) {
	if viper_instance == nil {
		checkErr(panicMode, eM[8], func() (int, error) {
			return returnErr(createErr(eM[8]))
		})
	}
	fmt.Fprint(os.Stdout, []any{"Building package... "}...)
	checkErr(panicMode, eM[9], func() (int, error) {
		f := initVars(viper_instance)
		os.Chdir(projectPath + "/Package")
		_, err := re.Build(projectPath, f, false)
		return returnErr(err)
	})
	fmt.Fprintln(os.Stdout, []any{"Build successful."}...)
	return 0, nil
}

func InstallPackage(downloadPath string, projectPath string, dirName string, panicMode string) (int, error) {
	fullName := "https://" + os.Getenv(mirror) + "/projects/" + projectPath
	fmt.Fprintf(os.Stdout, "The package path on the mirror is %s and it will download to %s.\nWould you like to proceed with the installation? [Y or n]", []any{projectPath, downloadPath}...)
	fmt.Scan(&conf)
	if conf == "Y" {
		fmt.Fprint(os.Stdout, []any{"Downloading package... "}...)
		checkErr(panicMode, eM[10], func() (int, error) {
			_, err := DownloadPackage(downloadPath, fullName, panicMode)
			return returnErr(err)
		})
		fmt.Fprintln(os.Stdout, []any{"Package downloaded successfully."}...)
		fmt.Fprint(os.Stdout, []any{"Unziping package... "}...)
		checkErr(panicMode, eM[11], func() (int, error) {
			cmd := exec.Command("tar", "-xzf", projectPath)
			cmd.Stdout = nil
			err := cmd.Run()
			return returnErr(err)
		})
		fmt.Fprintln(os.Stdout, []any{"Package unziped successfully."}...)
		fmt.Fprint(os.Stdout, []any{"Building package... "}...)
		checkErr(panicMode, eM[12], func() (int, error) {
			os.Chdir(dirName)
			_, err := BuildPackage(".", panicMode)
			return returnErr(err)
		})
		fmt.Fprintln(os.Stdout, []any{"Installation completed! 🎉"}...)
	} else if conf == "n" {
		fmt.Fprintln(os.Stdout, []any{"Installation aborted."}...)
		os.Exit(0)
	}
	return 0, nil
}
