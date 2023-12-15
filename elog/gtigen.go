// Code generated by "goki generate ./... -add-types"; DO NOT EDIT.

package elog

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/elog.WriteFunc",
	ShortName:  "elog.WriteFunc",
	IDName:     "write-func",
	Doc:        "WriteFunc function that computes and sets log values\nThe Context provides information typically needed for logging",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/elog.Context",
	ShortName:  "elog.Context",
	IDName:     "context",
	Doc:        "Context provides the context for logging Write functions.\nSetContext must be called on Logs to set the Stats and Net values\nProvides various convenience functions for setting log values\nand other commonly-used operations.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Logs", &gti.Field{Name: "Logs", Type: "*github.com/emer/emergent/v2/elog.Logs", LocalType: "*Logs", Doc: "pointer to the Logs object with all log data", Directives: gti.Directives{}, Tag: ""}},
		{"Stats", &gti.Field{Name: "Stats", Type: "*github.com/emer/emergent/v2/estats.Stats", LocalType: "*estats.Stats", Doc: "pointer to stats", Directives: gti.Directives{}, Tag: ""}},
		{"Net", &gti.Field{Name: "Net", Type: "github.com/emer/emergent/v2/emer.Network", LocalType: "emer.Network", Doc: "network", Directives: gti.Directives{}, Tag: ""}},
		{"Di", &gti.Field{Name: "Di", Type: "int", LocalType: "int", Doc: "data parallel index for accessing data from network", Directives: gti.Directives{}, Tag: ""}},
		{"Item", &gti.Field{Name: "Item", Type: "*github.com/emer/emergent/v2/elog.Item", LocalType: "*Item", Doc: "current log Item", Directives: gti.Directives{}, Tag: ""}},
		{"Scope", &gti.Field{Name: "Scope", Type: "github.com/emer/emergent/v2/etime.ScopeKey", LocalType: "etime.ScopeKey", Doc: "current scope key", Directives: gti.Directives{}, Tag: ""}},
		{"Mode", &gti.Field{Name: "Mode", Type: "github.com/emer/emergent/v2/etime.Modes", LocalType: "etime.Modes", Doc: "current scope eval mode (if standard)", Directives: gti.Directives{}, Tag: ""}},
		{"Time", &gti.Field{Name: "Time", Type: "github.com/emer/emergent/v2/etime.Times", LocalType: "etime.Times", Doc: "current scope timescale (if standard)", Directives: gti.Directives{}, Tag: ""}},
		{"LogTable", &gti.Field{Name: "LogTable", Type: "*github.com/emer/emergent/v2/elog.LogTable", LocalType: "*LogTable", Doc: "LogTable with extra data for the table", Directives: gti.Directives{}, Tag: ""}},
		{"Table", &gti.Field{Name: "Table", Type: "*goki.dev/etable/v2/etable.Table", LocalType: "*etable.Table", Doc: "current table to record value to", Directives: gti.Directives{}, Tag: ""}},
		{"Row", &gti.Field{Name: "Row", Type: "int", LocalType: "int", Doc: "current row in table to write to", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/elog.WriteMap",
	ShortName:  "elog.WriteMap",
	IDName:     "write-map",
	Doc:        "WriteMap holds log writing functions for scope keys",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/elog.Item",
	ShortName:  "elog.Item",
	IDName:     "item",
	Doc:        "Item describes one item to be logged -- has all the info\nfor this item, across all scopes where it is relevant.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Name", &gti.Field{Name: "Name", Type: "string", LocalType: "string", Doc: "name of column -- must be unique for a table", Directives: gti.Directives{}, Tag: ""}},
		{"Type", &gti.Field{Name: "Type", Type: "goki.dev/etable/v2/etensor.Type", LocalType: "etensor.Type", Doc: "data type, using etensor types which are isomorphic with arrow.Type", Directives: gti.Directives{}, Tag: ""}},
		{"CellShape", &gti.Field{Name: "CellShape", Type: "[]int", LocalType: "[]int", Doc: "shape of a single cell in the column (i.e., without the row dimension) -- for scalars this is nil -- tensor column will add the outer row dimension to this shape", Directives: gti.Directives{}, Tag: ""}},
		{"DimNames", &gti.Field{Name: "DimNames", Type: "[]string", LocalType: "[]string", Doc: "names of the dimensions within the CellShape -- 'Row' will be added to outer dimension", Directives: gti.Directives{}, Tag: ""}},
		{"Write", &gti.Field{Name: "Write", Type: "github.com/emer/emergent/v2/elog.WriteMap", LocalType: "WriteMap", Doc: "holds Write functions for different scopes.  After processing, the scope key will be a single mode and time, from Scope(mode, time), but the initial specification can lists for each, or the All* option, if there is a Write function that works across scopes", Directives: gti.Directives{}, Tag: ""}},
		{"Plot", &gti.Field{Name: "Plot", Type: "bool", LocalType: "bool", Doc: "Whether or not to plot it", Directives: gti.Directives{}, Tag: ""}},
		{"Range", &gti.Field{Name: "Range", Type: "goki.dev/etable/v2/minmax.F64", LocalType: "minmax.F64", Doc: "The minimum and maximum values, for plotting", Directives: gti.Directives{}, Tag: ""}},
		{"FixMin", &gti.Field{Name: "FixMin", Type: "bool", LocalType: "bool", Doc: "Whether to fix the minimum in the display", Directives: gti.Directives{}, Tag: ""}},
		{"FixMax", &gti.Field{Name: "FixMax", Type: "bool", LocalType: "bool", Doc: "Whether to fix the maximum in the display", Directives: gti.Directives{}, Tag: ""}},
		{"ErrCol", &gti.Field{Name: "ErrCol", Type: "string", LocalType: "string", Doc: "Name of other item that has the error bar values for this item -- for plotting", Directives: gti.Directives{}, Tag: ""}},
		{"TensorIdx", &gti.Field{Name: "TensorIdx", Type: "int", LocalType: "int", Doc: "index of tensor to plot -- defaults to 0 -- use -1 to plot all", Directives: gti.Directives{}, Tag: ""}},
		{"Color", &gti.Field{Name: "Color", Type: "string", LocalType: "string", Doc: "specific color for plot -- uses default ordering of colors if empty", Directives: gti.Directives{}, Tag: ""}},
		{"Modes", &gti.Field{Name: "Modes", Type: "map[string]bool", LocalType: "map[string]bool", Doc: "map of eval modes that this item has a Write function for", Directives: gti.Directives{}, Tag: ""}},
		{"Times", &gti.Field{Name: "Times", Type: "map[string]bool", LocalType: "map[string]bool", Doc: "map of times that this item has a Write function for", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/elog.Logs",
	ShortName:  "elog.Logs",
	IDName:     "logs",
	Doc:        "Logs contains all logging state and API for doing logging.\ndo AddItem to add any number of items, at different eval mode, time scopes.\nEach Item has its own Write functions, at each scope as neeeded.\nThen call CreateTables to generate log Tables from those items.\nCall Log with a scope to add a new row of data to the log\nand ResetLog to reset the log to empty.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Tables", &gti.Field{Name: "Tables", Type: "map[github.com/emer/emergent/v2/etime.ScopeKey]*github.com/emer/emergent/v2/elog.LogTable", LocalType: "map[etime.ScopeKey]*LogTable", Doc: "Tables storing log data, auto-generated from Items.", Directives: gti.Directives{}, Tag: ""}},
		{"MiscTables", &gti.Field{Name: "MiscTables", Type: "map[string]*goki.dev/etable/v2/etable.Table", LocalType: "map[string]*etable.Table", Doc: "holds additional tables not computed from items -- e.g., aggregation results, intermediate computations, etc", Directives: gti.Directives{}, Tag: ""}},
		{"Items", &gti.Field{Name: "Items", Type: "[]*github.com/emer/emergent/v2/elog.Item", LocalType: "[]*Item", Doc: "A list of the items that should be logged. Each item should describe one column that you want to log, and how.  Order in list determines order in logs.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Context", &gti.Field{Name: "Context", Type: "github.com/emer/emergent/v2/elog.Context", LocalType: "Context", Doc: "context information passed to logging Write functions -- has all the information needed to compute and write log values -- is updated for each item in turn", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Modes", &gti.Field{Name: "Modes", Type: "map[string]bool", LocalType: "map[string]bool", Doc: "All the eval modes that appear in any of the items of this log.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Times", &gti.Field{Name: "Times", Type: "map[string]bool", LocalType: "map[string]bool", Doc: "All the timescales that appear in any of the items of this log.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"ItemIdxMap", &gti.Field{Name: "ItemIdxMap", Type: "map[string]int", LocalType: "map[string]int", Doc: "map of item indexes by name, for rapid access to items if they need to be modified after adding.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"TableOrder", &gti.Field{Name: "TableOrder", Type: "[]github.com/emer/emergent/v2/etime.ScopeKey", LocalType: "[]etime.ScopeKey", Doc: "sorted order of table scopes", Directives: gti.Directives{}, Tag: "view:\"-\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/elog.LogTable",
	ShortName:  "elog.LogTable",
	IDName:     "log-table",
	Doc:        "LogTable contains all the data for one log table",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Table", &gti.Field{Name: "Table", Type: "*goki.dev/etable/v2/etable.Table", LocalType: "*etable.Table", Doc: "Actual data stored.", Directives: gti.Directives{}, Tag: ""}},
		{"Meta", &gti.Field{Name: "Meta", Type: "map[string]string", LocalType: "map[string]string", Doc: "arbitrary meta-data for each table, e.g., hints for plotting: Plot = false to not plot, XAxisCol, LegendCol", Directives: gti.Directives{}, Tag: ""}},
		{"IdxView", &gti.Field{Name: "IdxView", Type: "*goki.dev/etable/v2/etable.IdxView", LocalType: "*etable.IdxView", Doc: "Index View of the table -- automatically updated when a new row of data is logged to the table.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"NamedViews", &gti.Field{Name: "NamedViews", Type: "map[string]*goki.dev/etable/v2/etable.IdxView", LocalType: "map[string]*etable.IdxView", Doc: "named index views onto the table that can be saved and used across multiple items -- these are reset to nil after a new row is written -- see NamedIdxView funtion for more details.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"File", &gti.Field{Name: "File", Type: "*os.File", LocalType: "*os.File", Doc: "File to store the log into.", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"WroteHeaders", &gti.Field{Name: "WroteHeaders", Type: "bool", LocalType: "bool", Doc: "true if headers for File have already been written", Directives: gti.Directives{}, Tag: "view:\"-\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
