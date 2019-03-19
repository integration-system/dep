// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package never

import (
	"sort"

	"github.com/integration-system/dep/gps"
)

var (
	M = sort.Strings
	_ = gps.Solve
)
