//
// Copyright 2019-2020 Jean-Francois Smigielski
//
// This software is supplied under the terms of the MIT License, a
// copy of which should be located in the distribution where this
// file was obtained (LICENSE.txt). A copy of the license may also be
// found online at https://opensource.org/licenses/MIT.
//

package main

import (
	"github.com/jfsmig/object-storage/internal/cmd-index-store-rocksdb"
	"log"
)

func main() {
	rootCmd := cmd_index_store_rocksdb.MainCommand()
	rootCmd.Use = "gunkan-index-store-rocksdb"
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Command error:", err)
	}
}
