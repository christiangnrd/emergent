// Code generated by "goki generate ./... -add-types"; DO NOT EDIT.

package ringidx

import (
	"goki.dev/gti"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/ringidx.FIx",
	ShortName: "ringidx.FIx",
	IDName:    "f-ix",
	Doc:       "FIx is a fixed-length ring index structure -- does not grow\nor shrink dynamically.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gosl", Directive: "start", Args: []string{"ringidx"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Zi", &gti.Field{Name: "Zi", Type: "uint32", LocalType: "uint32", Doc: "the zero index position -- where logical 0 is in physical buffer", Directives: gti.Directives{}, Tag: ""}},
		{"Len", &gti.Field{Name: "Len", Type: "uint32", LocalType: "uint32", Doc: "the length of the buffer -- wraps around at this modulus", Directives: gti.Directives{}, Tag: ""}},
		{"pad", &gti.Field{Name: "pad", Type: "uint32", LocalType: "uint32", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

var _ = gti.AddType(&gti.Type{
	Name:      "github.com/emer/emergent/v2/ringidx.Idx",
	ShortName: "ringidx.Idx",
	IDName:    "idx",
	Doc:       "Idx is the ring index structure, maintaining starting index and length\ninto a ring-buffer with maximum length Max.  Max must be > 0 and Len <= Max.\nWhen adding new items would overflow Max, starting index is shifted over\nto overwrite the oldest items with the new ones.  No moving is ever\nrequired -- just a fixed-length buffer of size Max.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "go", Directive: "generate", Args: []string{"goki", "generate", "-add-types"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"StIdx", &gti.Field{Name: "StIdx", Type: "int", LocalType: "int", Doc: "the starting index where current data starts -- the oldest data is at this index, and continues for Len items, wrapping around at Max, coming back up at most to StIdx-1", Directives: gti.Directives{}, Tag: ""}},
		{"Len", &gti.Field{Name: "Len", Type: "int", LocalType: "int", Doc: "the number of items stored starting at StIdx.  Capped at Max", Directives: gti.Directives{}, Tag: ""}},
		{"Max", &gti.Field{Name: "Max", Type: "int", LocalType: "int", Doc: "the maximum number of items that can be stored in this ring", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})
