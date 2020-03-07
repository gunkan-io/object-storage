//
// Copyright 2019-2020 Jean-Francois Smigielski
//
// This software is supplied under the terms of the MIT License, a
// copy of which should be located in the distribution where this
// file was obtained (LICENSE.txt). A copy of the license may also be
// found online at https://opensource.org/licenses/MIT.
//

package cmd_blob_store_fs

import (
	"github.com/jfsmig/object-storage/pkg/gunkan"
	"golang.org/x/sys/unix"
)

const (
	flagsCommon   int = unix.O_NOATIME | unix.O_NONBLOCK | unix.O_CLOEXEC
	flagsRO           = flagsCommon | unix.O_RDONLY
	flagsRW           = flagsCommon | unix.O_RDWR
	flagsCreate       = flagsRW | unix.O_EXCL | unix.O_CREAT
	flagsOpenDir      = flagsRO | unix.O_DIRECTORY | unix.O_PATH
	flagsOpenRead     = flagsRO
)

const (
	routeInfo   = "/info"
	routeHealth = "/health"
	routeStatus = "/status"
	routeList   = "/v1/list"
	prefixBlob  = "/v1/blob/"
	infoString  = "gunkan/blob-store-" + gunkan.VersionString
)
