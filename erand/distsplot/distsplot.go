// Copyright (c) 2020, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// distsplot plots histograms of random distributions
package main

import (
	"strconv"

	"github.com/emer/emergent/v2/erand"
	"goki.dev/etable/v2/eplot"
	"goki.dev/etable/v2/etable"
	"goki.dev/etable/v2/etensor"
	_ "goki.dev/etable/v2/etview" // include to get gui views
	"goki.dev/etable/v2/histogram"
	"goki.dev/etable/v2/minmax"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/gi/v2/giv"
	"goki.dev/ki/v2/ki"
	"goki.dev/mat32/v2"
)

func main() {
	TheSim.Config()
	gimain.Main(func() { // this starts gui -- requires valid OpenGL display connection (e.g., X11)
		guirun()
	})
}

func guirun() {
	win := TheSim.ConfigGui()
	win.StartEventLoop()
}

// LogPrec is precision for saving float values in logs
const LogPrec = 4

// Sim holds the params, table, etc
type Sim struct {

	// random params
	Dist erand.RndParams

	// number of samples
	NSamp int

	// number of bins in the histogram
	NBins int

	// range for histogram
	Range minmax.F64

	// table for raw data
	Table *etable.Table `view:"no-inline"`

	// histogram of data
	Hist *etable.Table `view:"no-inline"`

	// the plot
	Plot *eplot.Plot2D `view:"-"`

	// main GUI window
	Win *gi.Window `view:"-"`

	// the master toolbar
	ToolBar *gi.ToolBar `view:"-"`
}

// TheSim is the overall state for this simulation
var TheSim Sim

// Config configures all the elements using the standard functions
func (ss *Sim) Config() {
	ss.Dist.Defaults()
	ss.Dist.Dist = erand.Gaussian
	ss.Dist.Mean = 0.5
	ss.Dist.Var = 0.15
	ss.NSamp = 1000000
	ss.NBins = 100
	ss.Range.Set(0, 1)
	ss.Update()
	ss.Table = &etable.Table{}
	ss.Hist = &etable.Table{}
	ss.ConfigTable(ss.Table)
	ss.Run()
}

// Update updates computed values
func (ss *Sim) Update() {
}

// Run generates the data and plots a histogram of results
func (ss *Sim) Run() {
	ss.Update()
	dt := ss.Table

	dt.SetNumRows(ss.NSamp)
	for vi := 0; vi < ss.NSamp; vi++ {
		vl := ss.Dist.Gen(-1)
		dt.SetCellFloat("Val", vi, float64(vl))
	}

	histogram.F64Table(ss.Hist, dt.Cols[0].(*etensor.Float64).Values, ss.NBins, ss.Range.Min, ss.Range.Max)
	if ss.Plot != nil {
		ss.Plot.UpdatePlot()
	}
}

func (ss *Sim) ConfigTable(dt *etable.Table) {
	dt.SetMetaData("name", "Data")
	dt.SetMetaData("read-only", "true")
	dt.SetMetaData("precision", strconv.Itoa(LogPrec))

	sch := etable.Schema{
		{"Val", etensor.FLOAT64, nil, nil},
	}
	dt.SetFromSchema(sch, 0)
}

func (ss *Sim) ConfigPlot(plt *eplot.Plot2D, dt *etable.Table) *eplot.Plot2D {
	plt.Params.Title = "Rand Dist Histogram"
	plt.Params.XAxisCol = "Value"
	plt.Params.Type = eplot.Bar
	plt.Params.XAxisRot = 45
	plt.SetTable(dt)
	// order of params: on, fixMin, min, fixMax, max
	plt.SetColParams("Value", eplot.Off, eplot.FloatMin, 0, eplot.FloatMax, 0)
	plt.SetColParams("Count", eplot.On, eplot.FixMin, 0, eplot.FloatMax, 0)
	return plt
}

// ConfigGui configures the GoGi gui interface for this simulation,
func (ss *Sim) ConfigGui() *gi.Window {
	width := 1600
	height := 1200

	// gi.WinEventTrace = true

	gi.SetAppName("distplot")
	gi.SetAppAbout(`This plots histograms of random distributions. See <a href="https://github.com/emer/emergent/v2">emergent on GitHub</a>.</p>`)

	win := gi.NewMainWindow("distplot", "Plotting Random Distributions", width, height)
	ss.Win = win

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()

	tbar := gi.AddNewToolBar(mfr, "tbar")
	tbar.SetStretchMaxWidth()
	ss.ToolBar = tbar

	split := gi.AddNewSplitView(mfr, "split")
	split.Dim = mat32.X
	split.SetStretchMax()

	sv := giv.AddNewStructView(split, "sv")
	sv.SetStruct(ss)

	tv := gi.AddNewTabView(split, "tv")

	plt := tv.AddNewTab(eplot.KiT_Plot2D, "Histogram").(*eplot.Plot2D)
	ss.Plot = ss.ConfigPlot(plt, ss.Hist)

	split.SetSplits(.3, .7)

	tbar.AddAction(gi.ActOpts{Label: "Run", Icon: "update", Tooltip: "Generate data and plot histogram."}, win.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		ss.Run()
		vp.SetNeedsFullRender()
	})

	tbar.AddAction(gi.ActOpts{Label: "README", Icon: "file-markdown", Tooltip: "Opens your browser on the README file that contains instructions for how to run this model."}, win.This(),
		func(recv, send ki.Ki, sig int64, data interface{}) {
			gi.OpenURL("https://github.com/emer/emergent/v2/blob/master/erand/distplot/README.md")
		})

	vp.UpdateEndNoSig(updt)

	// main menu
	appnm := gi.AppName()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "File", "Edit", "Window"})

	amen := win.MainMenu.ChildByName(appnm, 0).(*gi.Action)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.ChildByName("Edit", 1).(*gi.Action)
	emen.Menu.AddCopyCutPaste(win)

	win.MainMenuUpdated()
	return win
}
