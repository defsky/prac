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
	"os"
	"os/user"
	"path/filepath"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var isDetached bool
var cmdFile string

const baseDir string = "/var/run/xtelnet"

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new <session name>",
	Short: "create a new session",
	Long:  `create a new xtelnet session`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("main start")

		pid, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if err != 0 {
			fmt.Printf("%v\n", err)
			return
		}

		if pid == 0 {
			child(args[0])
		} else {

			if isDetached {
				fmt.Println("start session in detached mode")
			}

			sessionName := fmt.Sprintf("%d.%s", pid, args[0])

			fmt.Println("Session name: ", sessionName)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().BoolVarP(&isDetached, "detach", "d", false, "create new session in detached status")
	newCmd.Flags().StringVarP(&cmdFile, "file", "f", "", "specify startup command file name")
}

func child(name string) {
	fname, err := socketFileName(name)
	if err != nil {
		return
	}

	unixAddr, err := net.ResolveUnixAddr("unix", fname)
	if err != nil {
		fmt.Println(err)
	}
	l, err := net.ListenUnix("unix", unixAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	l.SetUnlinkOnClose(true)

	os.Chmod(fname, os.ModeSocket|os.FileMode(0600))

	for {
		conn, err := l.AcceptUnix()
		if err != nil {
			conn.Close()
			continue
		}
		
		go handleIncoming(conn)
	}
}

func handleIncoming(conn *net.UnixConn) {
	defer conn.Close()

	w := bufio.NewWriter(conn)

	w.WriteString("hello\n")
	w.Flush()

	counter := 0
	for {
		_, err := w.WriteString(fmt.Sprintf("Counter: %d\n", counter))
		if err != nil {
			break
		}
		w.Flush()

		time.Sleep(time.Second)
		counter++
	}
}

func mkdirIfNotExist(path string, mode os.FileMode) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, mode)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}

func socketFileName(name string) (string, error) {
	homedir, err := socketHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%d.%s", homedir, os.Getpid(), name), nil
}

func socketHomeDir() (string, error) {
	err := mkdirIfNotExist(baseDir, os.ModeDir|os.FileMode(0775))
	if err != nil {
		return "", err
	}

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homedirName := fmt.Sprintf("S-%s", currentUser.Username)
	dirname := filepath.Join(baseDir, homedirName)

	err = mkdirIfNotExist(dirname, os.ModeDir|os.FileMode(0700))
	if err != nil {
		return "", err
	}
	return dirname, nil
}
