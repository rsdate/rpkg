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
	download_dir   string       = "DOWNLOAD_DIR"
	viper_instance *viper.Viper = viper.GetViper()
	conf           string
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

func DownloadPackage(filepath string, url string, panicMode bool) (int, error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, []any{"Could not create file"}...)
		if panicMode {
			panic(fmt.Errorf("error: could not create the file. error message: %v", []any{err}...))
		} else {
			return 1, fmt.Errorf("error: could not create the file. error message: %v", []any{err}...)
		}
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, []any{"Could not get data from server"}...)
		if panicMode {
			panic(fmt.Errorf("error: could not get data from the server. error message: %v", []any{err}...))
		} else {
			return 1, fmt.Errorf("error: could not get data from the server. error message: %v", []any{err}...)
		}
	}
	defer resp.Body.Close()

	// Check server response
	switch code := resp.StatusCode; code {
	case http.StatusNotFound:
		fmt.Fprintln(os.Stderr, []any{"Server did not find file"}...)
		if panicMode {
			panic(fmt.Errorf("error: server did not find file"))
		} else {
			return 1, fmt.Errorf("error: server did not find file")
		}
	case http.StatusForbidden:
		fmt.Fprintln(os.Stderr, []any{"Server did not allow permission to access the resource"}...)
		if panicMode {
			panic(fmt.Errorf("error: server did not allow permission to access the resource"))
		} else {
			return 1, fmt.Errorf("error: server did not allow permission to access the resource")
		}
	case http.StatusUnauthorized:
		fmt.Fprintln(os.Stderr, []any{"User is not authorized to access the resource"}...)
		if panicMode {
			panic(fmt.Errorf("error: user is not authorized to access the resource"))
		} else {
			return 1, fmt.Errorf("error: user is not authorized to access the resource")
		}
	case http.StatusInternalServerError:
		fmt.Fprintln(os.Stderr, []any{"Server encountered an internal error"}...)
		if panicMode {
			panic(fmt.Errorf("error: server encountered an internal error"))
		} else {
			return 1, fmt.Errorf("error: server encountered an internal error")
		}
	case http.StatusServiceUnavailable:
		fmt.Fprintln(os.Stderr, []any{"Server is currently unavailable"}...)
		if panicMode {
			panic(fmt.Errorf("error: server is currently unavailable"))
		} else {
			return 1, fmt.Errorf("error: server is currently unavailable")
		}
	case http.StatusOK:
		return 0, nil
	}
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, []any{"Could not write to file"}...)
		if panicMode {
			panic(fmt.Errorf("error: could not write to the file. error message: %v", []any{err}...))
		} else {
			return 1, fmt.Errorf("error: could not write to the file. error message: %v", []any{err}...)
		}
	}
	return 0, nil
}

func BuildPackage(projectPath string, panicMode bool) (int, error) {
	if viper_instance == nil {
		fmt.Fprintln(os.Stderr, []any{"No configuration file found. If your build file is located in a different directory, please specify the path using the --buildfile flag"}...)
		if panicMode {
			panic(fmt.Errorf("error: no configuartion file found"))
		} else {
			return 1, fmt.Errorf("error: no configuartion file found")
		}
	}
	f := initVars(viper_instance)
	os.Chdir(projectPath + "/Package")
	fmt.Fprint(os.Stdout, []any{"Building package... "}...)
	if _, err := re.Build(projectPath, f, false); err != nil {
		fmt.Fprintln(os.Stderr, []any{"Build failed"}...)
		if panicMode {
			panic(fmt.Errorf("error: build failed. error message: %v", []any{err}...))
		} else {
			return 1, fmt.Errorf("error: build failed. error message: %v", []any{err}...)
		}
	} else {
		fmt.Fprintln(os.Stdout, []any{"Build successful."}...)
		return 0, nil
	}
}

func InstallPackage(downloadPath string, projectPath string, dirName string, panicMode bool) (int, error) {
	fullName := "https://" + os.Getenv(mirror) + "/projects/" + projectPath
	fmt.Fprintf(os.Stdout, "The package path on the mirror is %s and it will download to %s.\nWould you like to proceed with the installation? [Y or n]", []any{projectPath, downloadPath}...)
	fmt.Scan(&conf)
	if conf == "Y" {
		fmt.Print("Downloading package... ")
		_, err := DownloadPackage(downloadPath, fullName, true)
		if err != nil {
			fmt.Fprintln(os.Stderr, []any{"Package could not be downloaded"})
			if panicMode {
				panic(fmt.Errorf("error: package could not be downloaded. error message: %v", []any{err}...))
			} else {
				return 1, fmt.Errorf("error: package could not be downloaded. error message: %v", []any{err}...)
			}
		}
		fmt.Fprintln(os.Stdout, []any{"Package downloaded successfully."}...)
		fmt.Fprint(os.Stdout, []any{"Unziping package... "}...)
		cmd := exec.Command("tar", "-xzf", projectPath)
		cmd.Stdout = nil
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, []any{"Could not unzip package"}...)
			if panicMode {
				panic(fmt.Errorf("error: package could not be unzipped. error message: %v", []any{err}...))
			} else {
				return 1, fmt.Errorf("error: package could not be unzipped. error message: %v", []any{err}...)
			}
		}
		fmt.Fprintln(os.Stdout, []any{"Package unziped successfully."}...)
		fmt.Fprint(os.Stdout, []any{"Building package... "}...)
		os.Chdir(dirName)
		if _, err := BuildPackage(".", true); err != nil {
			fmt.Fprintln(os.Stdout, []any{"Build failed."}...)
			if panicMode {
				panic(fmt.Errorf("error: package could not be built. error message: %v", []any{err}...))
			} else {
				return 1, fmt.Errorf("error: package could not be built. error message: %v", []any{err}...)
			}
		}
		fmt.Fprintln(os.Stdout, []any{"Installation completed! 🎉"}...)
	} else if conf == "n" {
		fmt.Fprintln(os.Stdout, []any{"Installation aborted."}...)
		os.Exit(0)
	}
	return 0, nil
}
