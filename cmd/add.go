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

var startInt, endInt int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Generate addition expression",
	Long: `add subcommand is used 
to generate addition expression`,
	Args: cobra.MaximumNArgs(1),
	Run:  addCmdHandler,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&startInt, "start", "s", 2, "start number")
	addCmd.Flags().IntVarP(&endInt, "end", "e", 9, "end number")
}

func addCmdHandler(cmd *cobra.Command, args []string) {
	fname := ""
	if len(args) > 0 {
		fname = args[0]
		fmt.Printf("result will be Writen in file:%s\n", fname)
	}
	n := 0
	for i := startInt; i <= endInt; i++ {
		for j := startInt; j <= endInt; j++ {
			sum := i + j
			if sum > 10 {
				if n%6 == 0 && n != 0 {
					fmt.Printf("\n\n")
				}
				fmt.Printf("%12s", fmt.Sprintf("%d + %d = %d", i, j, sum))
				n++
			}
		}
	}
	fmt.Print("\n")
}
