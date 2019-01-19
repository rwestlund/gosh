// Copyright (c) 2019, Randy Westlund. All rights reserved.
// This code is under the BSD-2-Clause license.

package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func cd(args []string) {
	var err error
	if len(args) == 0 {
		err = os.Chdir(env.usr.HomeDir)
	} else {
		err = os.Chdir(args[0])
	}
	if err != nil {
		log.Println(errors.Wrap(err, "When changing dir"))
	}
}
