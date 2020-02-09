/*
Copyright © 2020 def<defsky@qq.com>

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

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Generate subtraction expression",
	Long: `sub subcommand is used
generate subtraction expression`,
	Run: subCmdHandler,
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func subCmdHandler(cmd *cobra.Command, args []string) {
	fmt.Println("sub called")
	n := 0
	for i := 11; i < 20; i++ {
		for j := 2; j < 10; j++ {
			sub := i - j
			if sub < 10 {
				if n%6 == 0 && n != 0 {
					fmt.Print("\n\n")
				}
				fmt.Printf("%12s", fmt.Sprintf("%d - %d = %d", i, j, sub))
				n++
			}
		}
	}
	fmt.Print("\n")
}
