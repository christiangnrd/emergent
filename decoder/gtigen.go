// Code generated by "goki generate ./... -add-types"; DO NOT EDIT.

package decoder

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.ActivationFunc",
	ShortName:  "decoder.ActivationFunc",
	IDName:     "activation-func",
	Doc:        "",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.Linear",
	ShortName:  "decoder.Linear",
	IDName:     "linear",
	Doc:        "Linear is a linear neural network, which can be configured with a custom\nactivation function. By default it will use the identity function.\nIt learns using the delta rule for each output unit.",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"LRate", &gti.Field{Name: "LRate", Type: "float32", LocalType: "float32", Doc: "learning rate", Directives: gti.Directives{}, Tag: "def:\"0.1\""}},
		{"Layers", &gti.Field{Name: "Layers", Type: "[]github.com/emer/emergent/v2/decoder.Layer", LocalType: "[]Layer", Doc: "layers to decode", Directives: gti.Directives{}, Tag: ""}},
		{"Units", &gti.Field{Name: "Units", Type: "[]github.com/emer/emergent/v2/decoder.LinearUnit", LocalType: "[]LinearUnit", Doc: "unit values -- read this for decoded output", Directives: gti.Directives{}, Tag: ""}},
		{"NInputs", &gti.Field{Name: "NInputs", Type: "int", LocalType: "int", Doc: "number of inputs -- total sizes of layer inputs", Directives: gti.Directives{}, Tag: ""}},
		{"NOutputs", &gti.Field{Name: "NOutputs", Type: "int", LocalType: "int", Doc: "number of outputs -- total sizes of layer inputs", Directives: gti.Directives{}, Tag: ""}},
		{"Inputs", &gti.Field{Name: "Inputs", Type: "[]float32", LocalType: "[]float32", Doc: "input values, copied from layers", Directives: gti.Directives{}, Tag: ""}},
		{"ValsTsrs", &gti.Field{Name: "ValsTsrs", Type: "map[string]*goki.dev/etable/v2/etensor.Float32", LocalType: "map[string]*etensor.Float32", Doc: "for holding layer values", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Weights", &gti.Field{Name: "Weights", Type: "goki.dev/etable/v2/etensor.Float32", LocalType: "etensor.Float32", Doc: "synaptic weights: outer loop is units, inner loop is inputs", Directives: gti.Directives{}, Tag: ""}},
		{"ActivationFn", &gti.Field{Name: "ActivationFn", Type: "github.com/emer/emergent/v2/decoder.ActivationFunc", LocalType: "ActivationFunc", Doc: "activation function", Directives: gti.Directives{}, Tag: ""}},
		{"PoolIndex", &gti.Field{Name: "PoolIndex", Type: "int", LocalType: "int", Doc: "which pool to use within a layer", Directives: gti.Directives{}, Tag: ""}},
		{"Comm", &gti.Field{Name: "Comm", Type: "*github.com/emer/empi/v2/mpi.Comm", LocalType: "*mpi.Comm", Doc: "mpi communicator -- MPI users must set this to their comm -- do direct assignment", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"MPIDWts", &gti.Field{Name: "MPIDWts", Type: "goki.dev/etable/v2/etensor.Float32", LocalType: "etensor.Float32", Doc: "delta weight changes: only for MPI mode -- outer loop is units, inner loop is inputs", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.Layer",
	ShortName:  "decoder.Layer",
	IDName:     "layer",
	Doc:        "Layer is the subset of emer.Layer that is used by this code",
	Directives: gti.Directives{},

	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.LinearUnit",
	ShortName:  "decoder.LinearUnit",
	IDName:     "linear-unit",
	Doc:        "LinearUnit has variables for Linear decoder unit",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Target", &gti.Field{Name: "Target", Type: "float32", LocalType: "float32", Doc: "target activation value -- typically 0 or 1 but can be within that range too", Directives: gti.Directives{}, Tag: ""}},
		{"Act", &gti.Field{Name: "Act", Type: "float32", LocalType: "float32", Doc: "final activation = sum x * w -- this is the decoded output", Directives: gti.Directives{}, Tag: ""}},
		{"Net", &gti.Field{Name: "Net", Type: "float32", LocalType: "float32", Doc: "net input = sum x * w", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.SoftMax",
	ShortName:  "decoder.SoftMax",
	IDName:     "soft-max",
	Doc:        "SoftMax is a softmax decoder, which is the best choice for a 1-hot classification\nusing the widely-used SoftMax function: https://en.wikipedia.org/wiki/Softmax_function",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Lrate", &gti.Field{Name: "Lrate", Type: "float32", LocalType: "float32", Doc: "learning rate", Directives: gti.Directives{}, Tag: "def:\"0.1\""}},
		{"Layers", &gti.Field{Name: "Layers", Type: "[]github.com/emer/emergent/v2/emer.Layer", LocalType: "[]emer.Layer", Doc: "layers to decode", Directives: gti.Directives{}, Tag: ""}},
		{"NCats", &gti.Field{Name: "NCats", Type: "int", LocalType: "int", Doc: "number of different categories to decode", Directives: gti.Directives{}, Tag: ""}},
		{"Units", &gti.Field{Name: "Units", Type: "[]github.com/emer/emergent/v2/decoder.SoftMaxUnit", LocalType: "[]SoftMaxUnit", Doc: "unit values", Directives: gti.Directives{}, Tag: ""}},
		{"Sorted", &gti.Field{Name: "Sorted", Type: "[]int", LocalType: "[]int", Doc: "sorted list of indexes into Units, in descending order from strongest to weakest -- i.e., Sortedhas the most likely categorization, and its activity is Units].Act", Directives: gti.Directives{}, Tag: ""}},
		{"NInputs", &gti.Field{Name: "NInputs", Type: "int", LocalType: "int", Doc: "number of inputs -- total sizes of layer inputs", Directives: gti.Directives{}, Tag: ""}},
		{"Inputs", &gti.Field{Name: "Inputs", Type: "[]float32", LocalType: "[]float32", Doc: "input values, copied from layers", Directives: gti.Directives{}, Tag: ""}},
		{"Target", &gti.Field{Name: "Target", Type: "int", LocalType: "int", Doc: "current target index of correct category", Directives: gti.Directives{}, Tag: ""}},
		{"ValsTsrs", &gti.Field{Name: "ValsTsrs", Type: "map[string]*goki.dev/etable/v2/etensor.Float32", LocalType: "map[string]*etensor.Float32", Doc: "for holding layer values", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"Weights", &gti.Field{Name: "Weights", Type: "goki.dev/etable/v2/etensor.Float32", LocalType: "etensor.Float32", Doc: "synaptic weights: outer loop is units, inner loop is inputs", Directives: gti.Directives{}, Tag: ""}},
		{"Comm", &gti.Field{Name: "Comm", Type: "*github.com/emer/empi/v2/mpi.Comm", LocalType: "*mpi.Comm", Doc: "mpi communicator -- MPI users must set this to their comm -- do direct assignment", Directives: gti.Directives{}, Tag: "view:\"-\""}},
		{"MPIDWts", &gti.Field{Name: "MPIDWts", Type: "goki.dev/etable/v2/etensor.Float32", LocalType: "etensor.Float32", Doc: "delta weight changes: only for MPI mode -- outer loop is units, inner loop is inputs", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.SoftMaxUnit",
	ShortName:  "decoder.SoftMaxUnit",
	IDName:     "soft-max-unit",
	Doc:        "SoftMaxUnit has variables for softmax decoder unit",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Act", &gti.Field{Name: "Act", Type: "float32", LocalType: "float32", Doc: "final activation = e^Ge / sum e^Ge", Directives: gti.Directives{}, Tag: ""}},
		{"Net", &gti.Field{Name: "Net", Type: "float32", LocalType: "float32", Doc: "net input = sum x * w", Directives: gti.Directives{}, Tag: ""}},
		{"Exp", &gti.Field{Name: "Exp", Type: "float32", LocalType: "float32", Doc: "exp(Net)", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:       "github.com/emer/emergent/v2/decoder.softMaxForSerialization",
	ShortName:  "decoder.softMaxForSerialization",
	IDName:     "soft-max-for-serialization",
	Doc:        "",
	Directives: gti.Directives{},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Weights", &gti.Field{Name: "Weights", Type: "[]float32", LocalType: "[]float32", Doc: "", Directives: gti.Directives{}, Tag: "json:\"weights\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
