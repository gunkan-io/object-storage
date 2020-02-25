//
// Copyright 2019 Jean-Francois Smigielski
//
// This software is supplied under the terms of the MIT License, a
// copy of which should be located in the distribution where this
// file was obtained (LICENSE.txt). A copy of the license may also be
// found online at https://opensource.org/licenses/MIT.
//

package cmd_blob_client

import (
	"errors"
	"github.com/jfsmig/object-storage/pkg/blob-client"
	"github.com/jfsmig/object-storage/pkg/blob-model"
	"github.com/spf13/cobra"
)

func DelCommand() *cobra.Command {
	var cfg config

	client := &cobra.Command{
		Use:     "del",
		Aliases: []string{"delete", "remove", "rm", "erase"},
		Short:   "Delete BLOBs from a service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New(("Missing Blob ID"))
			}
			client, err := gunkan_blob_client.Dial(cfg.url)
			if err != nil {
				return err
			}
			if len(args) == 1 {
				id := args[0]
				err = delOne(client, id)
				debug(id, err)
				return err
			}

			strongError := false
			for _, id := range args {
				err = delOne(client, id)
				debug(id, err)
				if err != gunkan_blob_client.ErrNotFound {
					strongError = true
				}
			}
			if strongError {
				err = errors.New("Unrecoverable error")
			} else {
				err = nil
			}
			return err
		},
	}

	client.Flags().StringVar(&cfg.url, "url", "", "IP:PORT endpoint of the service to contact")

	return client
}

func delOne(client gunkan_blob_client.Client, strid string) error {
	var err error
	var id gunkan_blob_model.Id

	if err = id.Decode(strid); err != nil {
		return err
	} else {
		return client.Delete(id)
	}
}
