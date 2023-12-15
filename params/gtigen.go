// Code generated by "goki generate ./... -add-types"; DO NOT EDIT.

package params

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/params.FlexVal",
	ShortName:  "params.FlexVal",
	IDName:     "flex-val",
	Doc:        "FlexVal is a specific flexible value for the Flex parameter map\nthat implements the StylerObj interface for CSS-style selection logic",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Nm", &gti.Field{Name: "Nm", Type: "string", LocalType: "string", Doc: "name of this specific object -- matches #Name selections", Directives: gti.Directives{}, Tag: ""}},
		{"Type", &gti.Field{Name: "Type", Type: "string", LocalType: "string", Doc: "type name of this object -- matches plain TypeName selections", Directives: gti.Directives{}, Tag: ""}},
		{"Cls", &gti.Field{Name: "Cls", Type: "string", LocalType: "string", Doc: "space-separated list of class name(s) -- match the .Class selections", Directives: gti.Directives{}, Tag: ""}},
		{"Obj", &gti.Field{Name: "Obj", Type: "any", LocalType: "any", Doc: "actual object with data that is set by the parameters", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/params.Flex",
	ShortName:  "params.Flex",
	IDName:     "flex",
	Doc:        "Flex supports arbitrary named parameter values that can be set\nby a Set of parameters, as a map of any objects.\nFirst initialize the map with set of names and a type to create\nblank values, then apply the Set to it.",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/params.History",
	ShortName:  "params.History",
	IDName:     "history",
	Doc:        "The params.History interface records history of parameters applied\nto a given object.",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/params.HistoryImpl",
	ShortName:  "params.HistoryImpl",
	IDName:     "history-impl",
	Doc:        "HistoryImpl implements the History interface.  Implementing object can\njust pass calls to a HistoryImpl field.",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.HyperVals",
	ShortName: "params.HyperVals",
	IDName:    "hyper-vals",
	Doc:       "HyperVals is a string-value map for storing hyperparameter values",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Hypers",
	ShortName: "params.Hypers",
	IDName:    "hypers",
	Doc:       "Hypers is a parallel structure to Params which stores information relevant\nto hyperparameter search as well as the values.\nUse the key \"Val\" for the default value. This is equivalant to the value in\nParams. \"Min\" and \"Max\" guid the range, and \"Sigma\" describes a Gaussian.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Params",
	ShortName: "params.Params",
	IDName:    "params",
	Doc:       "Params is a name-value map for parameter values that can be applied\nto any numeric type in any object.\nThe name must be a dot-separated path to a specific parameter, e.g., Prjn.Learn.Lrate\nThe first part of the path is the overall target object type, e.g., \"Prjn\" or \"Layer\",\nwhich is used for determining if the parameter applies to a given object type.\n\nAll of the params in one map must apply to the same target type because\nonly the first item in the map (which could be any due to order randomization)\nis used for checking the type of the target.  Also, they all fall within the same\nSel selector scope which is used to determine what specific objects to apply the\nparameters to.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Sel",
	ShortName: "params.Sel",
	IDName:    "sel",
	Doc:       "params.Sel specifies a selector for the scope of application of a set of\nparameters, using standard css selector syntax (. prefix = class, # prefix = name,\nand no prefix = type)",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Sel", &gti.Field{Name: "Sel", Type: "string", LocalType: "string", Doc: "selector for what to apply the parameters to, using standard css selector syntax: .Example applies to anything with a Class tag of 'Example', #Example applies to anything with a Name of 'Example', and Example with no prefix applies to anything of type 'Example'", Directives: gti.Directives{}, Tag: "width:\"30\""}},
		{"Desc", &gti.Field{Name: "Desc", Type: "string", LocalType: "string", Doc: "description of these parameter values -- what effect do they have?  what range was explored?  it is valuable to record this information as you explore the params.", Directives: gti.Directives{}, Tag: "width:\"60\""}},
		{"Params", &gti.Field{Name: "Params", Type: "github.com/emer/emergent/v2/params.Params", LocalType: "Params", Doc: "parameter values to apply to whatever matches the selector", Directives: gti.Directives{}, Tag: "view:\"no-inline\""}},
		{"Hypers", &gti.Field{Name: "Hypers", Type: "github.com/emer/emergent/v2/params.Hypers", LocalType: "Hypers", Doc: "Put your hyperparams here", Directives: gti.Directives{}, Tag: ""}},
		{"NMatch", &gti.Field{Name: "NMatch", Type: "int", LocalType: "int", Doc: "number of times this selector matched a target during the last Apply process -- a warning is issued for any that remain at 0 -- see Sheet SelMatchReset and SelNoMatchWarn methods", Directives: gti.Directives{}, Tag: "tableview:\"-\" toml:\"-\" json:\"-\" xml:\"-\" inactive:\"+\""}},
		{"SetName", &gti.Field{Name: "SetName", Type: "string", LocalType: "string", Doc: "name of current Set being applied", Directives: gti.Directives{}, Tag: "tableview:\"-\" toml:\"-\" json:\"-\" xml:\"-\" inactive:\"+\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Sheet",
	ShortName: "params.Sheet",
	IDName:    "sheet",
	Doc:       "Sheet is a CSS-like style-sheet of params.Sel values, each of which represents\na different set of specific parameter values applied according to the Sel selector:\n.Class #Name or Type.\n\nThe order of elements in the Sheet list is critical, as they are applied\nin the order given by the list (slice), and thus later Sel's can override\nthose applied earlier.  Thus, you generally want to have more general Type-level\nparameters listed first, and then subsequently more specific ones (.Class and #Name)\n\nThis is the highest level of params that has an Apply method -- above this level\napplication must be done under explicit program control.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Sheets",
	ShortName: "params.Sheets",
	IDName:    "sheets",
	Doc:       "Sheets is a map of named sheets -- used in the Set",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Set",
	ShortName: "params.Set",
	IDName:    "set",
	Doc:       "Set is a collection of Sheets that constitute a coherent set of parameters --\na particular specific configuration of parameters, which the user selects to use.\nThe Set name is stored in the Sets map from which it is typically accessed.\nA good strategy is to have a \"Base\" set that has all the best parameters so far,\nand then other sets can modify relative to that one.  It is up to the Sim code to\napply parameter sets in whatever order is desired.\n\nWithin a params.Set, multiple different params.Sheets can be organized,\nwith each CSS-style sheet achieving a relatively complete parameter styling\nof a given element of the overal model, e.g., \"Network\", \"Sim\", \"Env\".\nOr Network could be further broken down into \"Learn\" vs. \"Act\" etc,\nor according to different brain areas (\"Hippo\", \"PFC\", \"BG\", etc).\nAgain, this is entirely at the discretion of the modeler and must be\nperformed under explict program control, especially because order is so critical.\n\nNote that there is NO deterministic ordering of the Sheets due to the use of\na Go map structure, which specifically randomizes order, so simply iterating over them\nand applying may produce unexpected results -- it is better to lookup by name.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Desc", &gti.Field{Name: "Desc", Type: "string", LocalType: "string", Doc: "description of this param set -- when should it be used?  how is it different from the other sets?", Directives: gti.Directives{}, Tag: "width:\"60\""}},
		{"Sheets", &gti.Field{Name: "Sheets", Type: "github.com/emer/emergent/v2/params.Sheets", LocalType: "Sheets", Doc: "Sheet's grouped according to their target and / or function, e.g.,", Directives: gti.Directives{}, Tag: "Network\" for all the network params (or \"Learn\" vs. \"Act\" for more fine-grained), and \"Sim\" for overall simulation control parameters, \"Env\" for environment parameters, etc.  It is completely up to your program to lookup these names and apply them as appropriate\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/params.Sets",
	ShortName: "params.Sets",
	IDName:    "sets",
	Doc:       "Sets is a collection of Set's that can be chosen among\ndepending on different desired configurations etc.  Thus, each Set\nrepresents a collection of different possible specific configurations,\nand different such configurations can be chosen by name to apply as desired.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{}},
	},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/params.Styler",
	ShortName:  "params.Styler",
	IDName:     "styler",
	Doc:        "The params.Styler interface exposes TypeName, Class, and Name methods\nthat allow the params.Sel CSS-style selection specifier to determine\nwhether a given parameter applies.\nAdding Set versions of Name and Class methods is a good idea but not\nneeded for this interface, so they are not included here.",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/params.StylerObj",
	ShortName:  "params.StylerObj",
	IDName:     "styler-obj",
	Doc:        "The params.StylerObj interface extends Styler to include an arbitary\nfunction to access the underlying object type.",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
