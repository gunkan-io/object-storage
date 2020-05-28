//
// Copyright 2019-2020 Jean-Francois Smigielski
//
// This software is supplied under the terms of the MIT License, a
// copy of which should be located in the distribution where this
// file was obtained (LICENSE.txt). A copy of the license may also be
// found online at https://opensource.org/licenses/MIT.
//

package gunkan

import (
	"context"
)

type IndexClient interface {
	Put(ctx context.Context, key BaseKey, value string) error

	Get(ctx context.Context, key BaseKey) (string, error)

	Delete(ctx context.Context, key BaseKey) error

	List(ctx context.Context, marker BaseKey, max uint32) ([]string, error)
}
