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

var withA, withBorrow bool
var testCount, columnCount, maxNumber int

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
	testCmd.Flags().BoolVarP(&withA, "answer", "a", false, "generate subjects with answer")
	testCmd.Flags().BoolVarP(&withBorrow, "borrow", "b", false, "generate subjects only with borrow and give")
	testCmd.Flags().IntVarP(&maxNumber, "max", "m", 20, "specify max number in subjects")
	testCmd.Flags().IntVarP(&testCount, "number", "n", 20, "specify count number for test subjects")
	testCmd.Flags().IntVarP(&columnCount, "columns", "c", 5, "specify column count for print out")
}

func testCmdHandler(cmd *cobra.Command, args []string) {
	rand.Seed(time.Now().Unix())

	addMap := make(map[string]bool)
	subMap := make(map[string]bool)

	addedNumber := maxNumber
	addNumber := maxNumber - 10
	if withBorrow {
		addedNumber = maxNumber - 11
		addNumber = maxNumber - 11
	}
	for i := 1; i <= addedNumber; i++ {
		if i == 10 {
			continue
		}
		for j := 2; j <= addNumber; j++ {
			if j == 10 {
				continue
			}
			sum := i + j
			if sum > 20 {
				continue
			}
			if sum > 10 {
				if withA {
					addMap[fmt.Sprintf("%-2d + %-2d = %-2d", i, j, sum)] = true
				} else {
					addMap[fmt.Sprintf("%-2d + %-2d = %-2s", i, j, "?")] = true
				}
			}
		}
	}

	subNumber := maxNumber
	if withBorrow {
		subNumber = maxNumber - 11
	}

	for i := 11; i <= maxNumber; i++ {
		for j := 2; j <= subNumber; j++ {
			sub := i - j
			if sub <= 0 || sub == 10 || j == 10 {
				continue
			}
			if withBorrow && sub >= 10 {
				continue
			}
			if withA {
				subMap[fmt.Sprintf("%-2d - %-2d = %-2d", i, j, sub)] = true
			} else {
				subMap[fmt.Sprintf("%-2d - %-2d = %-2s", i, j, "?")] = true
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
			fmt.Printf("%16s", subject)
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
