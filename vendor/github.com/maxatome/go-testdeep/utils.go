// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package testdeep

import (
	"reflect"
	"time"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
)

// getTime returns the time.Time that is inside got or that can be
// converted from got contents.
func getTime(ctx ctxerr.Context, got reflect.Value, mustConvert bool) (time.Time, *ctxerr.Error) {
	var (
		gotIf interface{}
		ok    bool
	)
	if mustConvert {
		gotIf, ok = dark.GetInterface(got.Convert(timeType), true)
	} else {
		gotIf, ok = dark.GetInterface(got, true)
	}
	if !ok {
		if ctx.BooleanError {
			return time.Time{}, ctxerr.BooleanError
		}
		return time.Time{}, &ctxerr.Error{
			Message: "cannot compare",
			Summary: ctxerr.NewSummary("unexported field that cannot be overridden"),
		}
	}
	return gotIf.(time.Time), nil
}
