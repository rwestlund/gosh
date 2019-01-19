// Copyright (c) 2019, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func handleLine(line string) {
	var parts = strings.Split(line, " ")
	if parts[0] == "" {
		return
	}
	// Check built-ins.
	switch parts[0] {
	case "cd":
		cd(parts[1:])
	default:
		// Look for a binary.
		var bin = findBinary(parts[0])
		if bin == "" {
			fmt.Println("binary not found")
			return
		}
		runBinary(bin, parts[1:])
	}
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
