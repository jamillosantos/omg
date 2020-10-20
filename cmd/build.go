/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/jamillosantos/omg/config"
	"github.com/jamillosantos/omg/external/buf"
	"github.com/jamillosantos/omg/external/protoc"
	"github.com/jamillosantos/omg/internal"
	"github.com/spf13/cobra"
)

var colorFileName = color.Style{color.Bold, color.White}

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the source code according to the configuration",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			files, err := internal.List(&config.Config)
			if err != nil {
				internal.Fatal(1, "error listing files: ", err)
			}
			args = files
		}

		bufCmd := buf.BuildCmd(&buf.BuildRequest{
			Source: config.Config.BufSources(),
		})
		bufCmd.Stderr = os.Stderr
		bufStdoutBuffer := bytes.NewBuffer(nil)
		bufCmd.Stdout = bufStdoutBuffer
		if err := bufCmd.Start(); err != nil {
			internal.Fatal(30, "error starting buf: ", err)
		}
		if err := bufCmd.Wait(); err != nil {
			internal.Fatal(50, "buf exec error: ", err)
		}

		for _, file := range args {
			fmt.Printf("Processing %s ...\n", colorFileName.Render(file))

			protocCmd := protoc.Run(&config.Config, file)
			protocCmd.Stderr = os.Stderr
			protocCmd.Stdout = os.Stdout
			protocStdin, err := protocCmd.StdinPipe()
			if err != nil {
				internal.Fatal(20, "error setting stdin for protoc: ", err)
			}

			if err := protocCmd.Start(); err != nil {
				internal.Fatal(40, "error starting protoc: ", err)
			}

			bufStdoutBuffer.WriteTo(protocStdin)

			protocStdin.Close()

			if err := protocCmd.Wait(); err != nil {
				internal.Fatal(60, "protoc error: ", err)
			}
			fmt.Printf("  %s\n", color.Green.Render("ok"))
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
