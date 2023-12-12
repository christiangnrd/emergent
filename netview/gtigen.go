// Code generated by "goki generate"; DO NOT EDIT.

package netview

import (
	"sync"

	"goki.dev/colors/colormap"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/giv"
	"goki.dev/goosi/events"
	"goki.dev/gti"
	"goki.dev/ki/v2"
	"goki.dev/ordmap"
	"goki.dev/xyz"
)

// Scene3DType is the [gti.Type] for [Scene3D]
var Scene3DType = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/netview.Scene3D",
	ShortName:  "netview.Scene3D",
	IDName:     "scene-3-d",
	Doc:        "Scene3D is a Widget for managing the 3D Scene",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"NetView", &gti.Field{Name: "NetView", Type: "*github.com/emer/emergent/v2/netview.NetView", LocalType: "*NetView", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Scene3D", &gti.Field{Name: "Scene3D", Type: "goki.dev/gi/v2/xyzv.Scene3D", LocalType: "xyzv.Scene3D", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &Scene3D{},
})

// NewScene3D adds a new [Scene3D] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewScene3D(par ki.Ki, name ...string) *Scene3D {
	return par.NewChild(Scene3DType, name...).(*Scene3D)
}

// KiType returns the [*gti.Type] of [Scene3D]
func (t *Scene3D) KiType() *gti.Type {
	return Scene3DType
}

// New returns a new [*Scene3D] value
func (t *Scene3D) New() ki.Ki {
	return &Scene3D{}
}

// SetNetView sets the [Scene3D.NetView]
func (t *Scene3D) SetNetView(v *NetView) *Scene3D {
	t.NetView = v
	return t
}

// SetTooltip sets the [Scene3D.Tooltip]
func (t *Scene3D) SetTooltip(v string) *Scene3D {
	t.Tooltip = v
	return t
}

// SetClass sets the [Scene3D.Class]
func (t *Scene3D) SetClass(v string) *Scene3D {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [Scene3D.PriorityEvents]
func (t *Scene3D) SetPriorityEvents(v []events.Types) *Scene3D {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [Scene3D.CustomContextMenu]
func (t *Scene3D) SetCustomContextMenu(v func(m *gi.Scene)) *Scene3D {
	t.CustomContextMenu = v
	return t
}

// LayNameType is the [gti.Type] for [LayName]
var LayNameType = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/netview.LayName",
	ShortName:  "netview.LayName",
	IDName:     "lay-name",
	Doc:        "LayName is the Layer name as a Text2D within the NetView",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"NetView", &gti.Field{Name: "NetView", Type: "*github.com/emer/emergent/v2/netview.NetView", LocalType: "*NetView", Doc: "our netview", Directives: gti.Directives{}, Tag: "copy:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Text2D", &gti.Field{Name: "Text2D", Type: "goki.dev/xyz.Text2D", LocalType: "xyz.Text2D", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &LayName{},
})

// NewLayName adds a new [LayName] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewLayName(par ki.Ki, name ...string) *LayName {
	return par.NewChild(LayNameType, name...).(*LayName)
}

// KiType returns the [*gti.Type] of [LayName]
func (t *LayName) KiType() *gti.Type {
	return LayNameType
}

// New returns a new [*LayName] value
func (t *LayName) New() ki.Ki {
	return &LayName{}
}

// SetNetView sets the [LayName.NetView]:
// our netview
func (t *LayName) SetNetView(v *NetView) *LayName {
	t.NetView = v
	return t
}

// SetMat sets the [LayName.Mat]
func (t *LayName) SetMat(v xyz.Material) *LayName {
	t.Mat = v
	return t
}

// SetText sets the [LayName.Text]
func (t *LayName) SetText(v string) *LayName {
	t.Text = v
	return t
}

// LayObjType is the [gti.Type] for [LayObj]
var LayObjType = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/netview.LayObj",
	ShortName: "netview.LayObj",
	IDName:    "lay-obj",
	Doc:       "LayObj is the Layer 3D object within the NetView",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"LayName", &gti.Field{Name: "LayName", Type: "string", LocalType: "string", Doc: "name of the layer we represent", Directives: gti.Directives{}, Tag: ""}},
		{"NetView", &gti.Field{Name: "NetView", Type: "*github.com/emer/emergent/v2/netview.NetView", LocalType: "*NetView", Doc: "our netview", Directives: gti.Directives{}, Tag: "copy:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Solid", &gti.Field{Name: "Solid", Type: "goki.dev/xyz.Solid", LocalType: "xyz.Solid", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods:  ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
	Instance: &LayObj{},
})

// NewLayObj adds a new [LayObj] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewLayObj(par ki.Ki, name ...string) *LayObj {
	return par.NewChild(LayObjType, name...).(*LayObj)
}

// KiType returns the [*gti.Type] of [LayObj]
func (t *LayObj) KiType() *gti.Type {
	return LayObjType
}

// New returns a new [*LayObj] value
func (t *LayObj) New() ki.Ki {
	return &LayObj{}
}

// SetLayName sets the [LayObj.LayName]:
// name of the layer we represent
func (t *LayObj) SetLayName(v string) *LayObj {
	t.LayName = v
	return t
}

// SetNetView sets the [LayObj.NetView]:
// our netview
func (t *LayObj) SetNetView(v *NetView) *LayObj {
	t.NetView = v
	return t
}

// SetMat sets the [LayObj.Mat]
func (t *LayObj) SetMat(v xyz.Material) *LayObj {
	t.Mat = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/netview.NetData",
	ShortName: "netview.NetData",
	IDName:    "net-data",
	Doc:       "NetData maintains a record of all the network data that has been displayed\nup to a given maximum number of records (updates), using efficient ring index logic\nwith no copying to store in fixed-sized buffers.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Net", &gti.Field{Name: "Net", Type: "github.com/emer/emergent/v2/emer.Network", LocalType: "emer.Network", Doc: "the network that we're viewing", Directives: gti.Directives{}, Tag: "json:\"-\""}},
		{"NoSynData", &gti.Field{Name: "NoSynData", Type: "bool", LocalType: "bool", Doc: "copied from Params -- do not record synapse level data -- turn this on for very large networks where recording the entire synaptic state would be prohibitive", Directives: gti.Directives{}, Tag: ""}},
		{"PrjnLay", &gti.Field{Name: "PrjnLay", Type: "string", LocalType: "string", Doc: "name of the layer with unit for viewing projections (connection / synapse-level values)", Directives: gti.Directives{}, Tag: ""}},
		{"PrjnUnIdx", &gti.Field{Name: "PrjnUnIdx", Type: "int", LocalType: "int", Doc: "1D index of unit within PrjnLay for for viewing projections", Directives: gti.Directives{}, Tag: ""}},
		{"PrjnType", &gti.Field{Name: "PrjnType", Type: "string", LocalType: "string", Doc: "copied from NetView Params: if non-empty, this is the type projection to show when there are multiple projections from the same layer -- e.g., Inhib, Lateral, Forward, etc", Directives: gti.Directives{}, Tag: "inactive:\"+\""}},
		{"UnVars", &gti.Field{Name: "UnVars", Type: "[]string", LocalType: "[]string", Doc: "the list of unit variables saved", Directives: gti.Directives{}, Tag: ""}},
		{"UnVarIdxs", &gti.Field{Name: "UnVarIdxs", Type: "map[string]int", LocalType: "map[string]int", Doc: "index of each variable in the Vars slice", Directives: gti.Directives{}, Tag: ""}},
		{"SynVars", &gti.Field{Name: "SynVars", Type: "[]string", LocalType: "[]string", Doc: "the list of synaptic variables saved", Directives: gti.Directives{}, Tag: ""}},
		{"SynVarIdxs", &gti.Field{Name: "SynVarIdxs", Type: "map[string]int", LocalType: "map[string]int", Doc: "index of synaptic variable in the SynVars slice", Directives: gti.Directives{}, Tag: ""}},
		{"Ring", &gti.Field{Name: "Ring", Type: "github.com/emer/emergent/v2/ringidx.Idx", LocalType: "ringidx.Idx", Doc: "the circular ring index -- Max here is max number of values to store, Len is number stored, and Idx(Len-1) is the most recent one, etc", Directives: gti.Directives{}, Tag: ""}},
		{"MaxData", &gti.Field{Name: "MaxData", Type: "int", LocalType: "int", Doc: "max data parallel data per unit", Directives: gti.Directives{}, Tag: ""}},
		{"LayData", &gti.Field{Name: "LayData", Type: "map[string]*github.com/emer/emergent/v2/netview.LayData", LocalType: "map[string]*LayData", Doc: "the layer data -- map keyed by layer name", Directives: gti.Directives{}, Tag: ""}},
		{"UnMinPer", &gti.Field{Name: "UnMinPer", Type: "[]float32", LocalType: "[]float32", Doc: "unit var min values for each Ring.Max * variable", Directives: gti.Directives{}, Tag: ""}},
		{"UnMaxPer", &gti.Field{Name: "UnMaxPer", Type: "[]float32", LocalType: "[]float32", Doc: "unit var max values for each Ring.Max * variable", Directives: gti.Directives{}, Tag: ""}},
		{"UnMinVar", &gti.Field{Name: "UnMinVar", Type: "[]float32", LocalType: "[]float32", Doc: "min values for unit variables", Directives: gti.Directives{}, Tag: ""}},
		{"UnMaxVar", &gti.Field{Name: "UnMaxVar", Type: "[]float32", LocalType: "[]float32", Doc: "max values for unit variables", Directives: gti.Directives{}, Tag: ""}},
		{"SynMinVar", &gti.Field{Name: "SynMinVar", Type: "[]float32", LocalType: "[]float32", Doc: "min values for syn variables", Directives: gti.Directives{}, Tag: ""}},
		{"SynMaxVar", &gti.Field{Name: "SynMaxVar", Type: "[]float32", LocalType: "[]float32", Doc: "max values for syn variables", Directives: gti.Directives{}, Tag: ""}},
		{"Counters", &gti.Field{Name: "Counters", Type: "[]string", LocalType: "[]string", Doc: "counter strings", Directives: gti.Directives{}, Tag: ""}},
		{"RasterCtrs", &gti.Field{Name: "RasterCtrs", Type: "[]int", LocalType: "[]int", Doc: "raster counter values", Directives: gti.Directives{}, Tag: ""}},
		{"RasterMap", &gti.Field{Name: "RasterMap", Type: "map[int]int", LocalType: "map[int]int", Doc: "map of raster counter values to record numbers", Directives: gti.Directives{}, Tag: ""}},
		{"RastCtr", &gti.Field{Name: "RastCtr", Type: "int", LocalType: "int", Doc: "dummy raster counter when passed a -1 -- increments and wraps around", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"OpenJSON", &gti.Method{Name: "OpenJSON", Doc: "OpenJSON opens colors from a JSON-formatted file.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"SaveJSON", &gti.Method{Name: "SaveJSON", Doc: "SaveJSON saves colors to a JSON-formatted file.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"error", &gti.Field{Name: "error", Type: "error", LocalType: "error", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
	}),
})

// NetViewType is the [gti.Type] for [NetView]
var NetViewType = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/netview.NetView",
	ShortName:  "netview.NetView",
	IDName:     "net-view",
	Doc:        "NetView is a GoGi Widget that provides a 3D network view using the GoGi gi3d\n3D framework.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Net", &gti.Field{Name: "Net", Type: "github.com/emer/emergent/v2/emer.Network", LocalType: "emer.Network", Doc: "the network that we're viewing", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Var", &gti.Field{Name: "Var", Type: "string", LocalType: "string", Doc: "current variable that we're viewing", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Di", &gti.Field{Name: "Di", Type: "int", LocalType: "int", Doc: "current data parallel index di, for networks capable of processing input patterns in parallel.", Directives: gti.Directives{}, Tag: ""}},
		{"Vars", &gti.Field{Name: "Vars", Type: "[]string", LocalType: "[]string", Doc: "the list of variables to view", Directives: gti.Directives{}, Tag: ""}},
		{"SynVars", &gti.Field{Name: "SynVars", Type: "[]string", LocalType: "[]string", Doc: "list of synaptic variables", Directives: gti.Directives{}, Tag: ""}},
		{"SynVarsMap", &gti.Field{Name: "SynVarsMap", Type: "map[string]int", LocalType: "map[string]int", Doc: "map of synaptic variable names to index", Directives: gti.Directives{}, Tag: ""}},
		{"VarParams", &gti.Field{Name: "VarParams", Type: "map[string]*github.com/emer/emergent/v2/netview.VarParams", LocalType: "map[string]*VarParams", Doc: "parameters for the list of variables to view", Directives: gti.Directives{}, Tag: ""}},
		{"CurVarParams", &gti.Field{Name: "CurVarParams", Type: "*github.com/emer/emergent/v2/netview.VarParams", LocalType: "*VarParams", Doc: "current var params -- only valid during Update of display", Directives: gti.Directives{}, Tag: "json:\"-\" xml:\"-\" view:\"-\""}},
		{"Params", &gti.Field{Name: "Params", Type: "github.com/emer/emergent/v2/netview.Params", LocalType: "Params", Doc: "parameters controlling how the view is rendered", Directives: gti.Directives{}, Tag: ""}},
		{"ColorMap", &gti.Field{Name: "ColorMap", Type: "*goki.dev/colors/colormap.Map", LocalType: "*colormap.Map", Doc: "color map for mapping values to colors -- set by name in Params", Directives: gti.Directives{}, Tag: ""}},
		{"ColorMapVal", &gti.Field{Name: "ColorMapVal", Type: "*goki.dev/gi/v2/giv.ColorMapValue", LocalType: "*giv.ColorMapValue", Doc: "color map value representing ColorMap", Directives: gti.Directives{}, Tag: ""}},
		{"RecNo", &gti.Field{Name: "RecNo", Type: "int", LocalType: "int", Doc: "record number to display -- use -1 to always track latest, otherwise in range", Directives: gti.Directives{}, Tag: ""}},
		{"LastCtrs", &gti.Field{Name: "LastCtrs", Type: "string", LocalType: "string", Doc: "last non-empty counters string provided -- re-used if no new one", Directives: gti.Directives{}, Tag: ""}},
		{"Data", &gti.Field{Name: "Data", Type: "github.com/emer/emergent/v2/netview.NetData", LocalType: "NetData", Doc: "contains all the network data with history", Directives: gti.Directives{}, Tag: ""}},
		{"DataMu", &gti.Field{Name: "DataMu", Type: "sync.RWMutex", LocalType: "sync.RWMutex", Doc: "mutex on data access", Directives: gti.Directives{}, Tag: "view:\"-\" copy:\"-\" json:\"-\" xml:\"-\""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Layout", &gti.Field{Name: "Layout", Type: "goki.dev/gi/v2/gi.Layout", LocalType: "gi.Layout", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{
		{"Current", &gti.Method{Name: "Current", Doc: "Current records the current state of the network, including synaptic values,\nand updates the display.  Use this when switching to NetView tab after network\nhas been running while viewing another tab, because the network state\nis typically not recored then.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"SaveWeights", &gti.Method{Name: "SaveWeights", Doc: "SaveWeights saves the network weights -- when called with giv.CallMethod\nit will auto-prompt for filename", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"OpenWeights", &gti.Method{Name: "OpenWeights", Doc: "OpenWeights opens the network weights -- when called with giv.CallMethod\nit will auto-prompt for filename", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"filename", &gti.Field{Name: "filename", Type: "goki.dev/gi/v2/gi.FileName", LocalType: "gi.FileName", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{})}},
		{"ShowNonDefaultParams", &gti.Method{Name: "ShowNonDefaultParams", Doc: "ShowNonDefaultParams shows a dialog of all the parameters that\nare not at their default values in the network.  Useful for setting params.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"ShowAllParams", &gti.Method{Name: "ShowAllParams", Doc: "ShowAllParams shows a dialog of all the parameters in the network.", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"ShowKeyLayerParams", &gti.Method{Name: "ShowKeyLayerParams", Doc: "ShowKeyLayerParams shows a dialog with a listing for all layers in the network,\nof the most important layer-level params (specific to each algorithm)", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
		{"ShowKeyPrjnParams", &gti.Method{Name: "ShowKeyPrjnParams", Doc: "ShowKeyPrjnParams shows a dialog with a listing for all Recv projections in the network,\nof the most important projection-level params (specific to each algorithm)", Directives: gti.Directives{
			&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
		}, Args: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}), Returns: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
			{"string", &gti.Field{Name: "string", Type: "string", LocalType: "string", Doc: "", Directives: gti.Directives{}, Tag: ""}},
		})}},
	}),
	Instance: &NetView{},
})

// NewNetView adds a new [NetView] with the given name
// to the given parent. If the name is unspecified, it defaults
// to the ID (kebab-case) name of the type, plus the
// [ki.Ki.NumLifetimeChildren] of the given parent.
func NewNetView(par ki.Ki, name ...string) *NetView {
	return par.NewChild(NetViewType, name...).(*NetView)
}

// KiType returns the [*gti.Type] of [NetView]
func (t *NetView) KiType() *gti.Type {
	return NetViewType
}

// New returns a new [*NetView] value
func (t *NetView) New() ki.Ki {
	return &NetView{}
}

// SetDi sets the [NetView.Di]:
// current data parallel index di, for networks capable of processing input patterns in parallel.
func (t *NetView) SetDi(v int) *NetView {
	t.Di = v
	return t
}

// SetVars sets the [NetView.Vars]:
// the list of variables to view
func (t *NetView) SetVars(v []string) *NetView {
	t.Vars = v
	return t
}

// SetSynVars sets the [NetView.SynVars]:
// list of synaptic variables
func (t *NetView) SetSynVars(v []string) *NetView {
	t.SynVars = v
	return t
}

// SetSynVarsMap sets the [NetView.SynVarsMap]:
// map of synaptic variable names to index
func (t *NetView) SetSynVarsMap(v map[string]int) *NetView {
	t.SynVarsMap = v
	return t
}

// SetVarParams sets the [NetView.VarParams]:
// parameters for the list of variables to view
func (t *NetView) SetVarParams(v map[string]*VarParams) *NetView {
	t.VarParams = v
	return t
}

// SetCurVarParams sets the [NetView.CurVarParams]:
// current var params -- only valid during Update of display
func (t *NetView) SetCurVarParams(v *VarParams) *NetView {
	t.CurVarParams = v
	return t
}

// SetParams sets the [NetView.Params]:
// parameters controlling how the view is rendered
func (t *NetView) SetParams(v Params) *NetView {
	t.Params = v
	return t
}

// SetColorMap sets the [NetView.ColorMap]:
// color map for mapping values to colors -- set by name in Params
func (t *NetView) SetColorMap(v *colormap.Map) *NetView {
	t.ColorMap = v
	return t
}

// SetColorMapVal sets the [NetView.ColorMapVal]:
// color map value representing ColorMap
func (t *NetView) SetColorMapVal(v *giv.ColorMapValue) *NetView {
	t.ColorMapVal = v
	return t
}

// SetRecNo sets the [NetView.RecNo]:
// record number to display -- use -1 to always track latest, otherwise in range
func (t *NetView) SetRecNo(v int) *NetView {
	t.RecNo = v
	return t
}

// SetLastCtrs sets the [NetView.LastCtrs]:
// last non-empty counters string provided -- re-used if no new one
func (t *NetView) SetLastCtrs(v string) *NetView {
	t.LastCtrs = v
	return t
}

// SetData sets the [NetView.Data]:
// contains all the network data with history
func (t *NetView) SetData(v NetData) *NetView {
	t.Data = v
	return t
}

// SetDataMu sets the [NetView.DataMu]:
// mutex on data access
func (t *NetView) SetDataMu(v sync.RWMutex) *NetView {
	t.DataMu = v
	return t
}

// SetTooltip sets the [NetView.Tooltip]
func (t *NetView) SetTooltip(v string) *NetView {
	t.Tooltip = v
	return t
}

// SetClass sets the [NetView.Class]
func (t *NetView) SetClass(v string) *NetView {
	t.Class = v
	return t
}

// SetPriorityEvents sets the [NetView.PriorityEvents]
func (t *NetView) SetPriorityEvents(v []events.Types) *NetView {
	t.PriorityEvents = v
	return t
}

// SetCustomContextMenu sets the [NetView.CustomContextMenu]
func (t *NetView) SetCustomContextMenu(v func(m *gi.Scene)) *NetView {
	t.CustomContextMenu = v
	return t
}

// SetStackTop sets the [NetView.StackTop]
func (t *NetView) SetStackTop(v int) *NetView {
	t.StackTop = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/netview.RasterParams",
	ShortName: "netview.RasterParams",
	IDName:    "raster-params",
	Doc:       "RasterParams holds parameters controlling the raster plot view",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"On", &gti.Field{Name: "On", Type: "bool", LocalType: "bool", Doc: "if true, show a raster plot over time, otherwise units", Directives: gti.Directives{}, Tag: ""}},
		{"XAxis", &gti.Field{Name: "XAxis", Type: "bool", LocalType: "bool", Doc: "if true, the raster counter (time) is plotted across the X axis -- otherwise the Z depth axis", Directives: gti.Directives{}, Tag: ""}},
		{"Max", &gti.Field{Name: "Max", Type: "int", LocalType: "int", Doc: "maximum count for the counter defining the raster plot", Directives: gti.Directives{}, Tag: ""}},
		{"UnitSize", &gti.Field{Name: "UnitSize", Type: "float32", LocalType: "float32", Doc: "size of a single unit, where 1 = full width and no space.. 1 default", Directives: gti.Directives{}, Tag: "min:\"0.1\" max:\"1\" step:\"0.1\" def:\"1\""}},
		{"UnitHeight", &gti.Field{Name: "UnitHeight", Type: "float32", LocalType: "float32", Doc: "height multiplier for units, where 1 = full height.. 0.2 default", Directives: gti.Directives{}, Tag: "min:\"0.1\" max:\"1\" step:\"0.1\" def:\"0.2\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/netview.Params",
	ShortName: "netview.Params",
	IDName:    "params",
	Doc:       "Params holds parameters controlling how the view is rendered",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Raster", &gti.Field{Name: "Raster", Type: "github.com/emer/emergent/v2/netview.RasterParams", LocalType: "RasterParams", Doc: "raster plot parameters", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"NoSynData", &gti.Field{Name: "NoSynData", Type: "bool", LocalType: "bool", Doc: "do not record synapse level data -- turn this on for very large networks where recording the entire synaptic state would be prohibitive", Directives: gti.Directives{}, Tag: ""}},
		{"PrjnType", &gti.Field{Name: "PrjnType", Type: "string", LocalType: "string", Doc: "if non-empty, this is the type projection to show when there are multiple projections from the same layer -- e.g., Inhib, Lateral, Forward, etc", Directives: gti.Directives{}, Tag: ""}},
		{"MaxRecs", &gti.Field{Name: "MaxRecs", Type: "int", LocalType: "int", Doc: "maximum number of records to store to enable rewinding through prior states", Directives: gti.Directives{}, Tag: "min:\"1\""}},
		{"NVarCols", &gti.Field{Name: "NVarCols", Type: "int", LocalType: "int", Doc: "number of variable columns", Directives: gti.Directives{}, Tag: ""}},
		{"UnitSize", &gti.Field{Name: "UnitSize", Type: "float32", LocalType: "float32", Doc: "size of a single unit, where 1 = full width and no space.. .9 default", Directives: gti.Directives{}, Tag: "min:\"0.1\" max:\"1\" step:\"0.1\" def:\"0.9\""}},
		{"LayNmSize", &gti.Field{Name: "LayNmSize", Type: "float32", LocalType: "float32", Doc: "size of the layer name labels -- entire network view is unit sized", Directives: gti.Directives{}, Tag: "min:\"0.01\" max:\".1\" step:\"0.01\" def:\"0.05\""}},
		{"ColorMap", &gti.Field{Name: "ColorMap", Type: "goki.dev/gi/v2/giv.ColorMapName", LocalType: "giv.ColorMapName", Doc: "name of color map to use", Directives: gti.Directives{}, Tag: ""}},
		{"ZeroAlpha", &gti.Field{Name: "ZeroAlpha", Type: "float32", LocalType: "float32", Doc: "opacity (0-1) of zero values -- greater magnitude values become increasingly opaque on either side of this minimum", Directives: gti.Directives{}, Tag: "min:\"0\" max:\"1\" step:\"0.1\" def:\"0.5\""}},
		{"NetView", &gti.Field{Name: "NetView", Type: "*github.com/emer/emergent/v2/netview.NetView", LocalType: "*NetView", Doc: "our netview, for update method", Directives: gti.Directives{}, Tag: "copy:\"-\" json:\"-\" xml:\"-\" view:\"-\""}},
		{"NFastSteps", &gti.Field{Name: "NFastSteps", Type: "int", LocalType: "int", Doc: "the number of records to jump for fast forward/backward", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/netview.VarParams",
	ShortName: "netview.VarParams",
	IDName:    "var-params",
	Doc:       "VarParams holds parameters for display of each variable",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Var", &gti.Field{Name: "Var", Type: "string", LocalType: "string", Doc: "name of the variable", Directives: gti.Directives{}, Tag: ""}},
		{"ZeroCtr", &gti.Field{Name: "ZeroCtr", Type: "bool", LocalType: "bool", Doc: "keep Min - Max centered around 0, and use negative heights for units -- else use full min-max range for height (no negative heights)", Directives: gti.Directives{}, Tag: ""}},
		{"Range", &gti.Field{Name: "Range", Type: "goki.dev/etable/v2/minmax.Range32", LocalType: "minmax.Range32", Doc: "range to display", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
		{"MinMax", &gti.Field{Name: "MinMax", Type: "goki.dev/etable/v2/minmax.F32", LocalType: "minmax.F32", Doc: "if not using fixed range, this is the actual range of data", Directives: gti.Directives{}, Tag: "view:\"inline\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
