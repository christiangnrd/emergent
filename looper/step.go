// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package looper

import "github.com/emer/emergent/etime"

// Step manages stepping state.  Also trace flag
type Step struct {
	Time      string `desc:"stepping level, time string"`
	Default   string `desc:"default stepping level"`
	N         int    `desc:"number of times to iterate at StepScope level, no stepping if 0"`
	Cnt       int    `view:"-" desc:"counter for number of times through loop"`
	LoopTrace bool   `desc:"if true, print out a trace of looping stages as they run"`
	FuncTrace bool   `desc:"if true, print out a trace of functions as they run -- implies LoopTrace"`
}

// Set sets the stepping time and n -- 0 = no stepping
// resets counter.
func (st *Step) Set(time string, n int) {
	st.Time = time
	st.N = n
	st.Cnt = 0
}

// Clear resets stepping (sets N = 0)
func (st *Step) Clear() {
	st.N = 0
}

// IsScope checks if given scope is stepping scope
func (st *Step) IsScope(scope etime.ScopeKey) bool {
	_, tm := scope.ModeAndTimeStr()
	return st.Time == tm
}

// StopCheck checks if it is time to stop for this scope
// returns true if so
func (st *Step) StopCheck(scope etime.ScopeKey) bool {
	if st.N <= 0 {
		return false
	}
	if !st.IsScope(scope) {
		return false
	}
	st.Cnt++
	if st.Cnt >= st.N {
		st.Cnt = 0
		return true
	}
	return false
}
