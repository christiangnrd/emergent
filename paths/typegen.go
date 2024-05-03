// Code generated by "core generate -add-types"; DO NOT EDIT.

package paths

import (
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.Circle", IDName: "circle", Doc: "Circle implements a circular pattern of connectivity between two layers\nwhere the center moves in proportion to receiver position with offset\nand multiplier factors, and a given radius is used (with wrap-around\noptionally).  A corresponding Gaussian bump of TopoWts is available as well.\nMakes for a good center-surround connectivity pattern.\n4D layers are automatically flattened to 2D for this connection.", Fields: []types.Field{{Name: "Radius", Doc: "radius of the circle, in units from center in sending layer"}, {Name: "Start", Doc: "starting offset in sending layer, for computing the corresponding sending center relative to given recv unit position"}, {Name: "Scale", Doc: "scaling to apply to receiving unit position to compute sending center as function of recv unit position"}, {Name: "AutoScale", Doc: "auto-scale sending center positions as function of relative sizes of send and recv layers -- if Start is positive then assumes it is a border, subtracted from sending size"}, {Name: "Wrap", Doc: "if true, connectivity wraps around edges"}, {Name: "TopoWts", Doc: "if true, this path should set gaussian topographic weights, according to following parameters"}, {Name: "Sigma", Doc: "gaussian sigma (width) as a proportion of the radius of the circle"}, {Name: "MaxWt", Doc: "maximum weight value for GaussWts function -- multiplies values"}, {Name: "SelfCon", Doc: "if true, and connecting layer to itself (self pathway), then make a self-connection from unit to itself"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.Full", IDName: "full", Doc: "Full implements full all-to-all pattern of connectivity between two layers", Fields: []types.Field{{Name: "SelfCon", Doc: "if true, and connecting layer to itself (self pathway), then make a self-connection from unit to itself"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.OneToOne", IDName: "one-to-one", Doc: "OneToOne implements point-to-point one-to-one pattern of connectivity between two layers", Fields: []types.Field{{Name: "NCons", Doc: "number of recv connections to make (0 for entire size of recv layer)"}, {Name: "SendStart", Doc: "starting unit index for sending connections"}, {Name: "RecvStart", Doc: "starting unit index for recv connections"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.Pattern", IDName: "pattern", Doc: "Pattern defines a pattern of connectivity between two layers.\nThe pattern is stored efficiently using a bitslice tensor of binary values indicating\npresence or absence of connection between two items.\nA receiver-based organization is generally assumed but connectivity can go either way."})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.PoolOneToOne", IDName: "pool-one-to-one", Doc: "PoolOneToOne implements one-to-one connectivity between pools within layers.\nPools are the outer-most two dimensions of a 4D layer shape.\nIf either layer does not have pools, then if the number of individual\nunits matches the number of pools in the other layer, those are connected one-to-one\notherwise each pool connects to the entire set of other units.\nIf neither is 4D, then it is equivalent to OneToOne.", Fields: []types.Field{{Name: "NPools", Doc: "number of recv pools to connect (0 for entire number of pools in recv layer)"}, {Name: "SendStart", Doc: "starting pool index for sending connections"}, {Name: "RecvStart", Doc: "starting pool index for recv connections"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.PoolRect", IDName: "pool-rect", Doc: "PoolRect implements a rectangular pattern of connectivity between\ntwo 4D layers, in terms of their pool-level shapes,\nwhere the lower-left corner moves in proportion to receiver\npool position with offset and multiplier factors (with wrap-around optionally).", Fields: []types.Field{{Name: "Size", Doc: "size of rectangle (of pools) in sending layer that each receiving unit receives from"}, {Name: "Start", Doc: "starting pool offset in sending layer, for computing the corresponding sending lower-left corner relative to given recv pool position"}, {Name: "Scale", Doc: "scaling to apply to receiving pool osition to compute corresponding position in sending layer of the lower-left corner of rectangle"}, {Name: "AutoScale", Doc: "auto-set the Scale as function of the relative pool sizes of send and recv layers (e.g., if sending layer is 2x larger than receiving, Scale = 2)"}, {Name: "RoundScale", Doc: "if true, use Round when applying scaling factor -- otherwise uses Floor which makes Scale work like a grouping factor -- e.g., .25 will effectively group 4 recv pools with same send position"}, {Name: "Wrap", Doc: "if true, connectivity wraps around all edges if it would otherwise go off the edge -- if false, then edges are clipped"}, {Name: "SelfCon", Doc: "if true, and connecting layer to itself (self pathway), then make a self-connection from unit to itself"}, {Name: "RecvStart", Doc: "starting pool position in receiving layer -- if > 0 then pools below this starting point remain unconnected"}, {Name: "RecvN", Doc: "number of pools in receiving layer to connect -- if 0 then all (remaining after RecvStart) are connected -- otherwise if < remaining then those beyond this point remain unconnected"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.PoolSameUnit", IDName: "pool-same-unit", Doc: "PoolSameUnit connects a given unit to the unit at the same index\nacross all the pools in a layer.\nPools are the outer-most two dimensions of a 4D layer shape.\nThis is most sensible when pools have same numbers of units in send and recv.\nThis is typically used for lateral topography-inducing connectivity\nand can also serve to reduce a pooled layer down to a single pool.\nThe logic works if either layer does not have pools.\nIf neither is 4D, then it is equivalent to OneToOne.", Fields: []types.Field{{Name: "SelfCon", Doc: "if true, and connecting layer to itself (self pathway), then make a self-connection from unit to itself"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.PoolTile", IDName: "pool-tile", Doc: "PoolTile implements tiled 2D connectivity between pools within layers, where\na 2D rectangular receptive field (defined over pools, not units) is tiled\nacross the sending layer pools, with specified level of overlap.\nPools are the outer-most two dimensions of a 4D layer shape.\n2D layers are assumed to have 1x1 pool.\nThis is a standard form of convolutional connectivity, where pools are\nthe filters and the outer dims are locations filtered.\nVarious initial weight / scaling patterns are also available -- code\nmust specifically apply these to the receptive fields.", Fields: []types.Field{{Name: "Recip", Doc: "reciprocal topographic connectivity -- logic runs with recv <-> send -- produces symmetric back-pathway or topo path when sending layer is larger than recv"}, {Name: "Size", Doc: "size of receptive field tile, in terms of pools on the sending layer"}, {Name: "Skip", Doc: "how many pools to skip in tiling over sending layer -- typically 1/2 of Size"}, {Name: "Start", Doc: "starting pool offset for lower-left corner of first receptive field in sending layer"}, {Name: "Wrap", Doc: "if true, pool coordinates wrap around sending shape -- otherwise truncated at edges, which can lead to assymmetries in connectivity etc"}, {Name: "GaussFull", Doc: "gaussian topographic weights / scaling parameters for full receptive field width. multiplies any other factors present"}, {Name: "GaussInPool", Doc: "gaussian topographic weights / scaling parameters within individual sending pools (i.e., unit positions within their parent pool drive distance for gaussian) -- this helps organize / differentiate units more within pools, not just across entire receptive field. multiplies any other factors present"}, {Name: "SigFull", Doc: "sigmoidal topographic weights / scaling parameters for full receptive field width.  left / bottom half have increasing sigmoids, and second half decrease.  Multiplies any other factors present (only used if Gauss versions are not On!)"}, {Name: "SigInPool", Doc: "sigmoidal topographic weights / scaling parameters within individual sending pools (i.e., unit positions within their parent pool drive distance for sigmoid) -- this helps organize / differentiate units more within pools, not just across entire receptive field. multiplies any other factors present  (only used if Gauss versions are not On!).  left / bottom half have increasing sigmoids, and second half decrease."}, {Name: "TopoRange", Doc: "min..max range of topographic weight values to generate"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.GaussTopo", IDName: "gauss-topo", Doc: "GaussTopo has parameters for Gaussian topographic weights or scaling factors", Fields: []types.Field{{Name: "On", Doc: "use gaussian topographic weights / scaling values"}, {Name: "Sigma", Doc: "gaussian sigma (width) in normalized units where entire distance across relevant dimension is 1.0 -- typical useful values range from .3 to 1.5, with .6 default"}, {Name: "Wrap", Doc: "wrap the gaussian around on other sides of the receptive field, with the closest distance being used -- this removes strict topography but ensures a more uniform distribution of weight values so edge units don't have weaker overall weights"}, {Name: "CtrMove", Doc: "proportion to move gaussian center relative to the position of the receiving unit within its pool: 1.0 = centers span the entire range of the receptive field.  Typically want to use 1.0 for Wrap = true, and 0.8 for false"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.SigmoidTopo", IDName: "sigmoid-topo", Doc: "SigmoidTopo has parameters for Gaussian topographic weights or scaling factors", Fields: []types.Field{{Name: "On", Doc: "use gaussian topographic weights / scaling values"}, {Name: "Gain", Doc: "gain of sigmoid that determines steepness of curve, in normalized units where entire distance across relevant dimension is 1.0 -- typical useful values range from 0.01 to 0.1"}, {Name: "CtrMove", Doc: "proportion to move gaussian center relative to the position of the receiving unit within its pool: 1.0 = centers span the entire range of the receptive field.  Typically want to use 1.0 for Wrap = true, and 0.8 for false"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.PoolTileSub", IDName: "pool-tile-sub", Doc: "PoolTileSub implements tiled 2D connectivity between pools within layers, where\na 2D rectangular receptive field (defined over pools, not units) is tiled\nacross the sending layer pools, with specified level of overlap.\nPools are the outer-most two dimensions of a 4D layer shape.\nSub version has sub-pools within each pool to encourage more independent\nrepresentations.\n2D layers are assumed to have 1x1 pool.\nThis is a standard form of convolutional connectivity, where pools are\nthe filters and the outer dims are locations filtered.\nVarious initial weight / scaling patterns are also available -- code\nmust specifically apply these to the receptive fields.", Fields: []types.Field{{Name: "Recip", Doc: "reciprocal topographic connectivity -- logic runs with recv <-> send -- produces symmetric back-pathway or topo path when sending layer is larger than recv"}, {Name: "Size", Doc: "size of receptive field tile, in terms of pools on the sending layer"}, {Name: "Skip", Doc: "how many pools to skip in tiling over sending layer -- typically 1/2 of Size"}, {Name: "Start", Doc: "starting pool offset for lower-left corner of first receptive field in sending layer"}, {Name: "Subs", Doc: "number of sub-pools within each pool"}, {Name: "SendSubs", Doc: "sending layer has sub-pools"}, {Name: "Wrap", Doc: "if true, pool coordinates wrap around sending shape -- otherwise truncated at edges, which can lead to assymmetries in connectivity etc"}, {Name: "GaussFull", Doc: "gaussian topographic weights / scaling parameters for full receptive field width. multiplies any other factors present"}, {Name: "GaussInPool", Doc: "gaussian topographic weights / scaling parameters within individual sending pools (i.e., unit positions within their parent pool drive distance for gaussian) -- this helps organize / differentiate units more within pools, not just across entire receptive field. multiplies any other factors present"}, {Name: "SigFull", Doc: "sigmoidal topographic weights / scaling parameters for full receptive field width.  left / bottom half have increasing sigmoids, and second half decrease.  Multiplies any other factors present (only used if Gauss versions are not On!)"}, {Name: "SigInPool", Doc: "sigmoidal topographic weights / scaling parameters within individual sending pools (i.e., unit positions within their parent pool drive distance for sigmoid) -- this helps organize / differentiate units more within pools, not just across entire receptive field. multiplies any other factors present  (only used if Gauss versions are not On!).  left / bottom half have increasing sigmoids, and second half decrease."}, {Name: "TopoRange", Doc: "min..max range of topographic weight values to generate"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.PoolUnifRnd", IDName: "pool-unif-rnd", Doc: "PoolUnifRnd implements random pattern of connectivity between pools within layers.\nPools are the outer-most two dimensions of a 4D layer shape.\nIf either layer does not have pools, PoolUnifRnd works as UnifRnd does.\nIf probability of connection (PCon) is 1, PoolUnifRnd works as PoolOnetoOne does.", Embeds: []types.Field{{Name: "PoolOneToOne"}, {Name: "UnifRnd"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.Rect", IDName: "rect", Doc: "Rect implements a rectangular pattern of connectivity between two layers\nwhere the lower-left corner moves in proportion to receiver position with offset\nand multiplier factors (with wrap-around optionally).\n4D layers are automatically flattened to 2D for this pathway.", Fields: []types.Field{{Name: "Size", Doc: "size of rectangle in sending layer that each receiving unit receives from"}, {Name: "Start", Doc: "starting offset in sending layer, for computing the corresponding sending lower-left corner relative to given recv unit position"}, {Name: "Scale", Doc: "scaling to apply to receiving unit position to compute corresponding position in sending layer of the lower-left corner of rectangle"}, {Name: "AutoScale", Doc: "auto-set the Scale as function of the relative sizes of send and recv layers (e.g., if sending layer is 2x larger than receiving, Scale = 2)"}, {Name: "RoundScale", Doc: "if true, use Round when applying scaling factor -- otherwise uses Floor which makes Scale work like a grouping factor -- e.g., .25 will effectively group 4 recv units with same send position"}, {Name: "Wrap", Doc: "if true, connectivity wraps around all edges if it would otherwise go off the edge -- if false, then edges are clipped"}, {Name: "SelfCon", Doc: "if true, and connecting layer to itself (self pathway), then make a self-connection from unit to itself"}, {Name: "Recip", Doc: "make the reciprocal of the specified connections -- i.e., symmetric for swapping recv and send"}, {Name: "RecvStart", Doc: "starting position in receiving layer -- if > 0 then units below this starting point remain unconnected"}, {Name: "RecvN", Doc: "number of units in receiving layer to connect -- if 0 then all (remaining after RecvStart) are connected -- otherwise if < remaining then those beyond this point remain unconnected"}}})

var _ = types.AddType(&types.Type{Name: "github.com/emer/emergent/v2/paths.UnifRnd", IDName: "unif-rnd", Doc: "UnifRnd implements uniform random pattern of connectivity between two layers\nusing a permuted (shuffled) list for without-replacement randomness,\nand maintains its own local random number source and seed\nwhich are initialized if Rand == nil -- usually best to keep this\nspecific to each instance of a pathway so it is fully reproducible\nand doesn't interfere with other random number streams.", Fields: []types.Field{{Name: "PCon", Doc: "probability of connection (0-1)"}, {Name: "SelfCon", Doc: "if true, and connecting layer to itself (self pathway), then make a self-connection from unit to itself"}, {Name: "Recip", Doc: "reciprocal connectivity: if true, switch the sending and receiving layers to create a symmetric top-down pathway -- ESSENTIAL to use same RndSeed between two paths to ensure symmetry"}, {Name: "Rand", Doc: "random number source -- is created with its own separate source if nil"}, {Name: "RndSeed", Doc: "the current random seed -- will be initialized to a new random number from the global random stream when Rand is created."}}})
