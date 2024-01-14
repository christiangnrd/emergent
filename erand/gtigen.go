// Code generated by "goki generate"; DO NOT EDIT.

package erand

import (
	"goki.dev/gti"
)

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/erand.Rand", IDName: "rand", Doc: "Rand provides an interface with most of the standard\nrand.Rand methods, to support the use of either the\nglobal rand generator or a separate Rand source.\nAll methods take an optional thr argument for the\nthread number in the parallel threaded case.\nThe gosl.slrand generator for example supports this.\nIf not used set to 0 or -1"})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/erand.SysRand", IDName: "sys-rand", Doc: "SysRand supports the system random number generator\nfor either a separate rand.Rand source, or, if that\nis nil, the global rand stream.", Fields: []gti.Field{{Name: "Rand", Doc: "if non-nil, use this random number source instead of the global default one"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/erand.RndParams", IDName: "rnd-params", Doc: "RndParams provides parameterized random number generation according to different distributions\nand variance, mean params", Directives: []gti.Directive{{Tool: "go", Directive: "generate", Args: []string{"goki", "generate"}}, {Tool: "git", Directive: "add"}}, Fields: []gti.Field{{Name: "Dist", Doc: "distribution to generate random numbers from"}, {Name: "Mean", Doc: "mean of random distribution -- typically added to generated random variants"}, {Name: "Var", Doc: "variability parameter for the random numbers (gauss = standard deviation, not variance; uniform = half-range, others as noted in RndDists)"}, {Name: "Par", Doc: "extra parameter for distribution (depends on each one)"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/erand.RndDists", IDName: "rnd-dists", Doc: "RndDists are different random number distributions", Directives: []gti.Directive{{Tool: "enums", Directive: "enum"}}})

var _ = gti.AddType(&gti.Type{Name: "github.com/emer/emergent/v2/erand.Seeds", IDName: "seeds", Doc: "Seeds is a set of random seeds, typically used one per Run"})
