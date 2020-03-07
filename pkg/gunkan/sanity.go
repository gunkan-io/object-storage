//
// Copyright 2019-2020 Jean-Francois Smigielski
//
// This software is supplied under the terms of the MIT License, a
// copy of which should be located in the distribution where this
// file was obtained (LICENSE.txt). A copy of the license may also be
// found online at https://opensource.org/licenses/MIT.
//

package gunkan

func ValidateBucketName(n string) bool {
	return len(n) > 0 && len(n) < 1024
}

func ValidateContentName(n string) bool {
	return len(n) > 0 && len(n) < 1024
}
