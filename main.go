// Copyright (c) 2019, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/pkg/errors"
)

// The global environment that everything needs.
var env struct {
	usr      *user.User
	hostname string
	cwd      string
}

func main() {
	var err error
	env.hostname, err = os.Hostname()
	if err != nil {
		log.Fatal(errors.Wrap(err, "When getting hostname"))
	}
	env.usr, err = user.Current()
	if err != nil {
		log.Fatal(errors.Wrap(err, "When getting username"))
	}
	updateCwd()

	printPrompt()
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		handleLine(line)
		printPrompt()
	}
	if scanner.Err() != nil {
		log.Fatal(errors.Wrap(scanner.Err(), "When scanning input"))
	}
}

func updateCwd() {
	var err error
	env.cwd, err = os.Getwd()
	if err != nil {
		log.Println(errors.Wrap(err, "When getting cwd"))
	}
}
func printPrompt() {
	var cwd = abbreviatePath(env.cwd)
	fmt.Fprintf(os.Stdout, "%s@%s %s> ", env.usr.Username, env.hostname, cwd)
}

// Convert e.g. /usr/home/foo/bar to /u/h/f/bar.
func abbreviatePath(path string) string {
	var parts = strings.Split(path, "/")
	if len(parts) < 2 {
		return path
	}
	for i := range parts[1:len(parts)] {
		if len(parts[i]) > 0 {
			parts[i] = string(parts[i][0])
		}
	}
	return strings.Join(parts, "/")
}
