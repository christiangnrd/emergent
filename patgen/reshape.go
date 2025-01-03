// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package patgen

import (
	"log"
	"reflect"

	"cogentcore.org/core/core"
	"cogentcore.org/lab/table"
	"cogentcore.org/lab/tensor"
)

// ReshapeCpp fixes C++ emergent table shape which is reversed from Go.
// it switches the dimension order in the given table, for all columns
// that are float 2D or 4D columns -- assumes these are layer patterns
// and names dimensions accordingly.
func ReshapeCpp(dt *table.Table) {
	for _, cl := range dt.Columns.Values {
		shp := cl.Shape().Sizes
		if cl.NumDims() == 3 && (cl.DataType() == reflect.Float32 || cl.DataType() == reflect.Float64) {
			revshp := []int{shp[0], shp[2], shp[1]} // [0] = row
			cl.SetShapeSizes(revshp...)
		}
		if cl.NumDims() == 5 && (cl.DataType() == reflect.Float32 || cl.DataType() == reflect.Float64) {
			revshp := []int{shp[0], shp[4], shp[3], shp[2], shp[1]} // [0] = row
			cl.SetShapeSizes(revshp...)
		}
	}
}

// ReshapeCppFile fixes C++ emergent table shape which is reversed from Go.
// It loads file from fname and saves to fixnm
func ReshapeCppFile(dt *table.Table, fname, fixnm string) {
	err := dt.OpenCSV(core.Filename(fname), tensor.Tab)
	if err != nil {
		log.Println(err)
		return
	}
	ReshapeCpp(dt)
	dt.SaveCSV(core.Filename(fixnm), tensor.Tab, true)
}
