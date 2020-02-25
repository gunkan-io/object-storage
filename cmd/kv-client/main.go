//
// Copyright 2019 Jean-Francois Smigielski
//
// This software is supplied under the terms of the MIT License, a
// copy of which should be located in the distribution where this
// file was obtained (LICENSE.txt). A copy of the license may also be
// found online at https://opensource.org/licenses/MIT.
//

package main

import (
	"github.com/jfsmig/object-storage/pkg/cmd-kv-client"
	"log"
)

func main() {
	rootCmd := cmd_blob_client.MainCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Command error:", err)
	}
}
