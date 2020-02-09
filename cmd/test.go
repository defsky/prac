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
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var withA bool
var testCount int
var columnCount int

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Generate test subject",
	Long: `test subcommand is used to
generate test subjects`,
	Run: testCmdHandler,
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	testCmd.Flags().BoolVarP(&withA, "answer", "a", false, "generate expression with answer")
	testCmd.Flags().IntVarP(&testCount, "number", "n", 20, "specify count number for test subjects")
	testCmd.Flags().IntVarP(&columnCount, "columns", "c", 5, "specify column count for print out")
}

func testCmdHandler(cmd *cobra.Command, args []string) {
	rand.Seed(time.Now().Unix())

	addMap := make(map[string]bool)
	subMap := make(map[string]bool)

	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			sum := i + j
			if sum > 10 {
				if withA {
					addMap[fmt.Sprintf("%d + %d = %d", i, j, sum)] = true
				} else {
					addMap[fmt.Sprintf("%d + %d = ?", i, j)] = true
				}
			}
		}
	}

	for i := 11; i < 20; i++ {
		for j := 2; j < 10; j++ {
			sub := i - j
			if sub < 10 {
				if withA {
					subMap[fmt.Sprintf("%d - %d = %d", i, j, sub)] = true
				} else {
					subMap[fmt.Sprintf("%d - %d = ?", i, j)] = true
				}
			}
		}
	}

	n := 0
	willPickAdd := true

	for i := 0; i < testCount; i++ {
		var expMap map[string]bool
		var subject string

		if willPickAdd {
			expMap = addMap
		} else {
			expMap = subMap
		}

		exps := make([]string, 0)
		for exp, valid := range expMap {
			if valid {
				exps = append(exps, exp)
			}
		}
		if len(exps) > 0 {
			subject = randPick(exps)
			expMap[subject] = false

			if n%columnCount == 0 && n != 0 {
				fmt.Print("\n\n")
			}
			fmt.Printf("%12s", subject)
			n++
		}

		willPickAdd = !willPickAdd
	}
	fmt.Print("\n")
}

func randPick(src []string) string {
	idx := rand.Intn(len(src))

	return src[idx]
}
