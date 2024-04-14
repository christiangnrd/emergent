// Copyright (c) 2022, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package egui

import (
	"cogentcore.org/core/core"
	"github.com/emer/emergent/v2/actrf"
	"github.com/emer/etable/v2/etensor"
	"github.com/emer/etable/v2/etview"
)

// Grid gets tensor grid view of given name, creating if not yet made
func (gui *GUI) Grid(name string) *etview.TensorGrid {
	if gui.Grids == nil {
		gui.Grids = make(map[string]*etview.TensorGrid)
	}
	tsr, ok := gui.Grids[name]
	if !ok {
		tsr = &etview.TensorGrid{}
		gui.Grids[name] = tsr
	}
	return tsr
}

// SetGrid sets tensor grid view to given name
func (gui *GUI) SetGrid(name string, tg *etview.TensorGrid) {
	if gui.Grids == nil {
		gui.Grids = make(map[string]*etview.TensorGrid)
	}
	gui.Grids[name] = tg
}

// ConfigRasterGrid configures a raster grid for given layer name.
// Uses Raster_laynm and given Tensor that has the raster data.
func (gui *GUI) ConfigRasterGrid(lay *core.Layout, laynm string, rast *etensor.Float32) *etview.TensorGrid {
	tg := gui.Grid(laynm)
	tg.SetName(laynm + "Raster")
	core.NewLabel(lay, laynm, laynm+":")
	lay.AddChild(tg)
	core.NewSpace(lay, laynm+"_spc")
	rast.SetMetaData("grid-fill", "1")
	tg.SetTensor(rast)
	return tg
}

// SaveActRFGrid stores the given TensorGrid in Grids under given name,
// and configures the grid view for ActRF viewing.
func (gui *GUI) SaveActRFGrid(tg *etview.TensorGrid, name string) {
	gui.SetGrid(name, tg)
}

// AddActRFGridTabs adds tabs for each of the ActRFs.
func (gui *GUI) AddActRFGridTabs(arfs *actrf.RFs) {
	for _, rf := range arfs.RFs {
		nm := rf.Name
		tf := gui.Tabs.NewTab(nm)
		tg := etview.NewTensorGrid(tf)
		gui.SaveActRFGrid(tg, nm)
	}
}

// ViewActRFs displays act rfs into tensor Grid views previously configured
func (gui *GUI) ViewActRFs(atf *actrf.RFs) {
	for _, rf := range atf.RFs {
		nm := rf.Name
		tg := gui.Grid(nm)
		if tg.Tensor == nil {
			rf := atf.RFByName(nm)
			tg.SetTensor(&rf.NormRF)
		} else {
			tg.NeedsRender()
		}
	}
}
