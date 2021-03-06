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
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list xtelnet sessions",
	Long:  `list xtelnet sessions`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		homedir, err := socketHomeDir()
		if err != nil {
			fmt.Println(err)
		}

		sessions, err := getSessionList(homedir)
		if err != nil {
			fmt.Print(err)
		}

		if len(sessions) > 0 {
			fmt.Println("There are sessions on:")
			for _, v := range sessions {
				fmt.Printf("  %s\n", v)
			}
		} else {
			fmt.Println("There is no session")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getSessionList(dir string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}
	for i, v := range files {
		_, n := filepath.Split(v)
		files[i] = n
	}
	return files, nil
}
