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
	"bufio"
	"fmt"
	"net"
	"path/filepath"
	"strings"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// attachCmd represents the attach command
var attachCmd = &cobra.Command{
	Use:   "attach <session name>",
	Short: "attach to the specified session",
	Long:  `attach to the specified session`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handleAttachCmd(args[0])
	},
}

func init() {
	rootCmd.AddCommand(attachCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attachCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attachCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func handleAttachCmd(name string) {
	homedir, err := socketHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	sessions, err := getSessionList(homedir)
	if err != nil {
		fmt.Println(err)
		return
	}

	matchedSession := make([]string, 0)
	for _, s := range sessions {
		if strings.HasSuffix(s, name) {
			matchedSession = append(matchedSession, s)
		}
		if strings.HasPrefix(s, name) {
			matchedSession = append(matchedSession, s)
		}
	}
	if len(matchedSession) == 0 {
		fmt.Printf("There is no matched session for name: %s\n", name)
		return
	}
	if len(matchedSession) > 1 {
		fmt.Println("There are more than one session matched:")
		for _, s := range matchedSession {
			fmt.Printf("  %s\n", s)
		}
		return
	}

	fpath := filepath.Join(homedir, matchedSession[0])

	sessionAddr, err := net.ResolveUnixAddr("unix", fpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// lname := filepath.Join(homedir, fmt.Sprintf("%d.sock", os.Getpid()))
	// lAddr, err := net.ResolveUnixAddr("unix", lname)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	conn, err := net.DialUnix("unix", nil, sessionAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	app := tview.NewApplication()
	screen := tview.NewTextView().
		SetChangedFunc(func() {
			app.Draw()
		})

	go func() {
		r := bufio.NewReader(conn)
		for {
			b, err := r.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Fprint(screen, b)
		}
	}()

	if err := app.SetRoot(screen, true).Run(); err != nil {
		panic(err)
	}
}
