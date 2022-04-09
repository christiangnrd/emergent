// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package looper

import "github.com/emer/emergent/etime"

// Loop represents one level of looping, with arbitrary functions
// called at 3 different points in the loop, corresponding to a
// do..while loop logic, with no initialization, which is necessary
// to ensure reentrant steppability.  In Go, the logic looks like this:
//
// for {
//    for { <subloops here> } // drills down levels for each subloop
//    Main()                  // Main is called after subloops -- increment counters!
//    if Stop() {
//        break
//    }
// }
// End()                      // Reset counters here so next pass starts over
//
type Loop struct {
	Scope etime.ScopeKey `desc:"scope level of this loop"`
	Main  Funcs          `desc:"main functions to call inside each iteration, after looping at lower level for non-terminal levels -- any counters should be incremented here"`
	Stop  BoolFuncs      `desc:"functions that cause the loop to stop -- if any return true, it stops"`
	End   Funcs          `desc:"functions to run at the end of the loop, after it has stopped.  counters should be reset here, so next iteration starts over."`
}

func NewLoop(sc etime.ScopeKey) *Loop {
	return &Loop{Scope: sc}
}
