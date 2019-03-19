// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package one

import (
	"cycle/two"

	"github.com/integration-system/dep/gps"
)

var (
	A = gps.Solve
	B = two.A
)
