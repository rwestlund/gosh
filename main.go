// Copyright (c) 2019, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	var hostname, err = os.Hostname()
	if err != nil {
		log.Fatal(errors.Wrap(err, "When getting hostname"))
	}
	var usr *user.User
	usr, err = user.Current()
	if err != nil {
		log.Fatal(errors.Wrap(err, "When getting username"))
	}

	printPrompt(usr.Username, hostname)
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		handleLine(line)
		printPrompt(usr.Username, hostname)
	}
	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "When scanning input"))
	}
}

func printPrompt(username, hostname string) {
	var cwd, err = os.Getwd()
	if err != nil {
		log.Println(errors.Wrap(err, "When getting cwd"))
	}
	fmt.Fprintf(os.Stdout, "%s@%s %s> ", username, hostname, cwd)
}

func handleLine(line string) {
	var parts = strings.Split(line, " ")
	if parts[0] == "" {
		return
	}
	var bin = findBinary(parts[0])
	if bin == "" {
		fmt.Println("binary not found")
		return
	}
	runBinary(bin, parts[1:])
}

func findBinary(name string) string {
	var path = os.Getenv("PATH")
	var parts = strings.Split(path, ":")
	for _, dir := range parts {
		var p = filepath.Join(dir, name)
		var stat, err = os.Stat(p)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(errors.Wrap(err, "When statting file path"))
		}
		if stat.IsDir() {
			log.Fatal("path is dir")
		}
		return p
	}
	return ""
}

func runBinary(bin string, args []string) {
	var cmd = exec.Command(bin, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	var err = cmd.Run()
	if err != nil {
		fmt.Println(errors.Wrap(err, "When running binary"))
	}
}
