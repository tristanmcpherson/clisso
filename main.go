/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package main

import (
	"log"
	"runtime"

	"github.com/mattn/go-colorable"

	"github.com/allcloud-io/clisso/cmd"
)

// This variable is used by the "version" command and is set during build.
var version = "undefined"

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Handle terminal colors on Windows machines.
	if runtime.GOOS == "windows" {
		log.SetOutput(colorable.NewColorableStdout())
	}

	cmd.Execute(version)
}
