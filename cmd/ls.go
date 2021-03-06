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
	"encoding/json"
	"fmt"
	"os"

	"github.com/jamillosantos/omg/config"
	"github.com/jamillosantos/omg/internal"
	"github.com/spf13/cobra"
)

var outputTypeFlag string

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all .proto files of the project.",
	Long:  `List all .proto files according to the 'src' definition at the 'omg.yaml'.`,
	Run: func(cmd *cobra.Command, args []string) {
		ls, err := internal.List(&config.Config)
		if err != nil {
			panic(err)
		}

		switch outputTypeFlag {
		case "text":
			for _, l := range ls {
				fmt.Println(l)
			}
		case "json":
			json.NewEncoder(os.Stdout).Encode(ls)
		default:
			fmt.Fprintln(os.Stderr, "invalid output-type")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().StringVar(&outputTypeFlag, "output-type", "text", "Defines the output type for the files (text, json) (DEFAULT text)")
}
