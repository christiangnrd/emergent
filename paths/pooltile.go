// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paths

import (
	"fmt"
	"log"

	"cogentcore.org/core/math32"
	"cogentcore.org/core/math32/minmax"
	"cogentcore.org/core/math32/vecint"
	"cogentcore.org/lab/tensor"
	"github.com/emer/emergent/v2/edge"
	"github.com/emer/emergent/v2/efuns"
)

// PoolTile implements tiled 2D connectivity between pools within layers, where
// a 2D rectangular receptive field (defined over pools, not units) is tiled
// across the sending layer pools, with specified level of overlap.
// Pools are the outer-most two dimensions of a 4D layer shape.
// 2D layers are assumed to have 1x1 pool.
// This is a standard form of convolutional connectivity, where pools are
// the filters and the outer dims are locations filtered.
// Various initial weight / scaling patterns are also available -- code
// must specifically apply these to the receptive fields.
type PoolTile struct {

	// reciprocal topographic connectivity -- logic runs with recv <-> send -- produces symmetric back-pathway or topo path when sending layer is larger than recv
	Recip bool

	// size of receptive field tile, in terms of pools on the sending layer
	Size vecint.Vector2i

	// how many pools to skip in tiling over sending layer -- typically 1/2 of Size
	Skip vecint.Vector2i

	// starting pool offset for lower-left corner of first receptive field in sending layer
	Start vecint.Vector2i

	// if true, pool coordinates wrap around sending shape -- otherwise truncated at edges, which can lead to assymmetries in connectivity etc
	Wrap bool

	// gaussian topographic weights / scaling parameters for full receptive field width. multiplies any other factors present
	GaussFull GaussTopo

	// gaussian topographic weights / scaling parameters within individual sending pools (i.e., unit positions within their parent pool drive distance for gaussian) -- this helps organize / differentiate units more within pools, not just across entire receptive field. multiplies any other factors present
	GaussInPool GaussTopo

	// sigmoidal topographic weights / scaling parameters for full receptive field width.  left / bottom half have increasing sigmoids, and second half decrease.  Multiplies any other factors present (only used if Gauss versions are not On!)
	SigFull SigmoidTopo

	// sigmoidal topographic weights / scaling parameters within individual sending pools (i.e., unit positions within their parent pool drive distance for sigmoid) -- this helps organize / differentiate units more within pools, not just across entire receptive field. multiplies any other factors present  (only used if Gauss versions are not On!).  left / bottom half have increasing sigmoids, and second half decrease.
	SigInPool SigmoidTopo

	// min..max range of topographic weight values to generate
	TopoRange minmax.F32
}

func NewPoolTile() *PoolTile {
	pt := &PoolTile{}
	pt.Defaults()
	return pt
}

// NewPoolTileRecip creates a new PoolTile that is a recip version of given ff feedforward one
func NewPoolTileRecip(ff *PoolTile) *PoolTile {
	pt := &PoolTile{}
	*pt = *ff
	pt.Recip = true
	return pt
}

func (pt *PoolTile) Defaults() {
	pt.Size.Set(4, 4)
	pt.Skip.Set(2, 2)
	pt.Start.Set(-1, -1)
	pt.Wrap = true
	pt.TopoRange.Min = 0.8
	pt.TopoRange.Max = 1
	pt.GaussFull.Defaults()
	pt.GaussInPool.Defaults()
	pt.SigFull.Defaults()
	pt.SigInPool.Defaults()
	pt.GaussFull.On = true
	pt.GaussInPool.On = true
}

func (pt *PoolTile) Name() string {
	return "PoolTile"
}

func (pt *PoolTile) Connect(send, recv *tensor.Shape, same bool) (sendn, recvn *tensor.Int32, cons *tensor.Bool) {
	if pt.Recip {
		return pt.ConnectRecip(send, recv, same)
	}
	sendn, recvn, cons = NewTensors(send, recv)
	sNtot := send.Len()
	sNpY := send.DimSize(0)
	sNpX := send.DimSize(1)
	rNpY := recv.DimSize(0)
	rNpX := recv.DimSize(1)
	sNu := 1
	rNu := 1
	if send.NumDims() == 4 {
		sNu = send.DimSize(2) * send.DimSize(3)
	} else {
		sNpY = 1
		sNpX = 1
		sNu = send.DimSize(0) * send.DimSize(1)
	}
	if recv.NumDims() == 4 {
		rNu = recv.DimSize(2) * recv.DimSize(3)
	} else {
		rNpY = 1
		rNpX = 1
		rNu = recv.DimSize(0) * recv.DimSize(1)
	}
	rnv := recvn.Values
	snv := sendn.Values
	var clip bool
	for rpy := 0; rpy < rNpY; rpy++ {
		for rpx := 0; rpx < rNpX; rpx++ {
			rpi := rpy*rNpX + rpx
			ris := rpi * rNu
			for fy := 0; fy < pt.Size.Y; fy++ {
				spy := pt.Start.Y + rpy*pt.Skip.Y + fy
				if spy, clip = edge.Edge(spy, sNpY, pt.Wrap); clip {
					continue
				}
				for fx := 0; fx < pt.Size.X; fx++ {
					spx := pt.Start.X + rpx*pt.Skip.X + fx
					if spx, clip = edge.Edge(spx, sNpX, pt.Wrap); clip {
						continue
					}
					spi := spy*sNpX + spx
					sis := spi * sNu
					for rui := 0; rui < rNu; rui++ {
						ri := ris + rui
						for sui := 0; sui < sNu; sui++ {
							si := sis + sui
							off := ri*sNtot + si
							if off < cons.Len() && ri < len(rnv) && si < len(snv) {
								// if !pt.SelfCon && same && ri == si {
								// 	continue
								// }
								cons.Values.Set(true, off)
								rnv[ri]++
								snv[si]++
							}
						}
					}
				}
			}
		}
	}
	return
}

func (pt *PoolTile) ConnectRecip(send, recv *tensor.Shape, same bool) (sendn, recvn *tensor.Int32, cons *tensor.Bool) {
	sendn, recvn, cons = NewTensors(send, recv)
	// all these variables are swapped: s from recv, r from send
	rNtot := send.Len()
	sNpY := recv.DimSize(0)
	sNpX := recv.DimSize(1)
	rNpY := send.DimSize(0)
	rNpX := send.DimSize(1)
	sNu := 1
	rNu := 1
	if recv.NumDims() == 4 {
		sNu = recv.DimSize(2) * recv.DimSize(3)
	} else {
		sNpY = 1
		sNpX = 1
		sNu = recv.DimSize(0) * recv.DimSize(1)
	}
	if send.NumDims() == 4 {
		rNu = send.DimSize(2) * send.DimSize(3)
	} else {
		rNpY = 1
		rNpX = 1
		rNu = send.DimSize(0) * send.DimSize(1)
	}
	snv := recvn.Values
	rnv := sendn.Values
	var clip bool
	for rpy := 0; rpy < rNpY; rpy++ {
		for rpx := 0; rpx < rNpX; rpx++ {
			rpi := rpy*rNpX + rpx
			ris := rpi * rNu
			for fy := 0; fy < pt.Size.Y; fy++ {
				spy := pt.Start.Y + rpy*pt.Skip.Y + fy
				if spy, clip = edge.Edge(spy, sNpY, pt.Wrap); clip {
					continue
				}
				for fx := 0; fx < pt.Size.X; fx++ {
					spx := pt.Start.X + rpx*pt.Skip.X + fx
					if spx, clip = edge.Edge(spx, sNpX, pt.Wrap); clip {
						continue
					}
					spi := spy*sNpX + spx
					sis := spi * sNu
					for sui := 0; sui < sNu; sui++ {
						si := sis + sui
						for rui := 0; rui < rNu; rui++ {
							ri := ris + rui
							off := si*rNtot + ri
							if off < cons.Len() && si < len(snv) && ri < len(rnv) {
								cons.Values.Set(true, off)
								snv[si]++
								rnv[ri]++
							}
						}
					}
				}
			}
		}
	}
	return
}

// HasTopoWeights returns true if some form of topographic weight patterns are set
func (pt *PoolTile) HasTopoWeights() bool {
	return pt.GaussFull.On || pt.GaussInPool.On || pt.SigFull.On || pt.SigInPool.On
}

// TopoWeights sets values in given 4D or 6D tensor according to *Topo settings.
// wts is shaped with first 2 outer-most dims as Y, X of units within layer / pool
// of recv layer (these are units over which topography is defined)
// and remaing 2D or 4D is for receptive field Size by units within pool size for
// sending layer.
func (pt *PoolTile) TopoWeights(send, recv *tensor.Shape, wts *tensor.Float32) error {
	if pt.GaussFull.On || pt.GaussInPool.On {
		if send.NumDims() == 2 {
			return pt.TopoWeightsGauss2D(send, recv, wts)
		} else {
			return pt.TopoWeightsGauss4D(send, recv, wts)
		}
	}
	if pt.SigFull.On || pt.SigInPool.On {
		if send.NumDims() == 2 {
			return pt.TopoWeightsSigmoid2D(send, recv, wts)
		} else {
			return pt.TopoWeightsSigmoid4D(send, recv, wts)
		}
	}
	err := fmt.Errorf("PoolTile:TopoWeights no Gauss or Sig params turned on")
	log.Println(err)
	return err
}

/////////////////////////////////////////////////////
// GaussTopo Wts

// GaussTopo has parameters for Gaussian topographic weights or scaling factors
type GaussTopo struct {

	// use gaussian topographic weights / scaling values
	On bool

	// gaussian sigma (width) in normalized units where entire distance across relevant dimension is 1.0 -- typical useful values range from .3 to 1.5, with .6 default
	Sigma float32 `default:"0.6"`

	// wrap the gaussian around on other sides of the receptive field, with the closest distance being used -- this removes strict topography but ensures a more uniform distribution of weight values so edge units don't have weaker overall weights
	Wrap bool

	// proportion to move gaussian center relative to the position of the receiving unit within its pool: 1.0 = centers span the entire range of the receptive field.  Typically want to use 1.0 for Wrap = true, and 0.8 for false
	CtrMove float32 `default:"0.8,1"`
}

func (gt *GaussTopo) Defaults() {
	gt.Sigma = 0.6
	gt.Wrap = true
	gt.CtrMove = 1
}

func (gt *GaussTopo) ShouldDisplay(field string) bool {
	switch field {
	case "On":
		return true
	default:
		return gt.On
	}
}

// DefWrap sets default wrap parameters (which are overall defaults): CtrMove = 1
func (gt *GaussTopo) DefWrap() {
	gt.Wrap = true
	gt.CtrMove = 1
}

// DefNoWrap sets default no-wrap parameters (CtrMove = .8 instead of 1)
func (gt *GaussTopo) DefNoWrap() {
	gt.Wrap = false
	gt.CtrMove = 0.8
}

// GaussOff turns off gaussian weights
func (pt *PoolTile) GaussOff() {
	pt.GaussFull.On = false
	pt.GaussInPool.On = false
}

// TopoWeightsGauss2D sets values in given 4D tensor according to *Topo settings.
// wts is shaped with first 2 outer-most dims as Y, X of units within layer / pool
// of recv layer (these are units over which topography is defined)
// and remaing 2D is for sending layer size (2D = sender)
func (pt *PoolTile) TopoWeightsGauss2D(send, recv *tensor.Shape, wts *tensor.Float32) error {
	if pt.GaussFull.Sigma == 0 {
		pt.GaussFull.Defaults()
	}
	if pt.GaussInPool.Sigma == 0 {
		pt.GaussInPool.Defaults()
	}
	sNuY := send.DimSize(0)
	sNuX := send.DimSize(1)
	rNuY := recv.DimSize(0) // ok if recv is 2D
	rNuX := recv.DimSize(1)
	if recv.NumDims() == 4 {
		rNuY = recv.DimSize(2)
		rNuX = recv.DimSize(3)
	}
	wts.SetShapeSizes(rNuY, rNuX, sNuY, sNuX)

	fsz := math32.Vec2(float32(sNuX-1), float32(sNuY-1)) // full rf size
	hfsz := fsz.MulScalar(0.5)                           // half rf
	fsig := pt.GaussFull.Sigma * hfsz.X                  // full sigma
	if fsig <= 0 {
		fsig = pt.GaussFull.Sigma
	}

	psz := math32.Vec2(float32(sNuX), float32(sNuY)) // within-pool rf size
	if sNuX > 1 {
		psz.X -= 1
	}
	if sNuY > 1 {
		psz.Y -= 1
	}
	hpsz := psz.MulScalar(0.5)            // half rf
	psig := pt.GaussInPool.Sigma * hpsz.X // pool sigma
	if psig <= 0 {
		psig = pt.GaussInPool.Sigma
	}

	rsz := math32.Vec2(float32(rNuX), float32(rNuY)) // recv units-in-pool size
	if rNuX > 1 {
		rsz.X -= 1
	}
	if rNuY > 1 {
		rsz.Y -= 1
	}

	hrsz := rsz.MulScalar(0.5)
	for ruy := 0; ruy < rNuY; ruy++ {
		for rux := 0; rux < rNuX; rux++ {
			rpos := math32.Vec2(float32(rux), float32(ruy)).Sub(hrsz).Div(hrsz) // -1..1 normalized r unit pos
			rfpos := rpos.MulScalar(pt.GaussFull.CtrMove)
			rppos := rpos.MulScalar(pt.GaussInPool.CtrMove)
			sfctr := rfpos.Mul(hfsz).Add(hfsz) // sending center for full
			spctr := rppos.Mul(hpsz).Add(hpsz) // sending center for within-pool
			for suy := 0; suy < sNuY; suy++ {
				for sux := 0; sux < sNuX; sux++ {
					fwt := float32(1)
					if pt.GaussFull.On {
						sf := math32.Vec2(float32(sux), float32(suy))
						if pt.GaussFull.Wrap {
							sf.X = edge.WrapMinDist(sf.X, fsz.X, sfctr.X)
							sf.Y = edge.WrapMinDist(sf.Y, fsz.Y, sfctr.Y)
						}
						fwt = efuns.GaussVecDistNoNorm(sf, sfctr, fsig)
					}
					pwt := float32(1)
					if pt.GaussInPool.On {
						sp := math32.Vec2(float32(sux), float32(suy))
						if pt.GaussInPool.Wrap {
							sp.X = edge.WrapMinDist(sp.X, psz.X, spctr.X)
							sp.Y = edge.WrapMinDist(sp.Y, psz.Y, spctr.Y)
						}
						pwt = efuns.GaussVecDistNoNorm(sp, spctr, psig)
					}
					wt := fwt * pwt
					rwt := pt.TopoRange.ProjValue(wt)
					wts.Set(rwt, ruy, rux, suy, sux)
				}
			}
		}
	}
	return nil
}

// TopoWeightsGauss4D sets values in given 6D tensor according to *Topo settings.
// wts is shaped with first 2 outer-most dims as Y, X of units within layer / pool
// of recv layer (these are units over which topography is defined)
// and remaing 4D is for receptive field Size by units within pool size for
// sending layer.
func (pt *PoolTile) TopoWeightsGauss4D(send, recv *tensor.Shape, wts *tensor.Float32) error {
	if pt.GaussFull.Sigma == 0 {
		pt.GaussFull.Defaults()
	}
	if pt.GaussInPool.Sigma == 0 {
		pt.GaussInPool.Defaults()
	}
	sNuY := send.DimSize(2)
	sNuX := send.DimSize(3)
	rNuY := recv.DimSize(0) // ok if recv is 2D
	rNuX := recv.DimSize(1)
	if recv.NumDims() == 4 {
		rNuY = recv.DimSize(2)
		rNuX = recv.DimSize(3)
	}
	wts.SetShapeSizes(rNuY, rNuX, pt.Size.Y, pt.Size.X, sNuY, sNuX)

	fsz := math32.Vec2(float32(pt.Size.X*sNuX-1), float32(pt.Size.Y*sNuY-1)) // full rf size
	hfsz := fsz.MulScalar(0.5)                                               // half rf
	fsig := pt.GaussFull.Sigma * hfsz.X                                      // full sigma
	if fsig <= 0 {
		fsig = pt.GaussFull.Sigma
	}

	psz := math32.Vec2(float32(sNuX), float32(sNuY)) // within-pool rf size
	if sNuX > 1 {
		psz.X -= 1
	}
	if sNuY > 1 {
		psz.Y -= 1
	}
	hpsz := psz.MulScalar(0.5)            // half rf
	psig := pt.GaussInPool.Sigma * hpsz.X // pool sigma
	if psig <= 0 {
		psig = pt.GaussInPool.Sigma
	}

	rsz := math32.Vec2(float32(rNuX), float32(rNuY)) // recv units-in-pool size
	if rNuX > 1 {
		rsz.X -= 1
	}
	if rNuY > 1 {
		rsz.Y -= 1
	}
	hrsz := rsz.MulScalar(0.5)
	for ruy := 0; ruy < rNuY; ruy++ {
		for rux := 0; rux < rNuX; rux++ {
			rpos := math32.Vec2(float32(rux), float32(ruy)).Sub(hrsz).Div(hrsz) // -1..1 normalized r unit pos
			rfpos := rpos.MulScalar(pt.GaussFull.CtrMove)
			rppos := rpos.MulScalar(pt.GaussInPool.CtrMove)
			sfctr := rfpos.Mul(hfsz).Add(hfsz) // sending center for full
			spctr := rppos.Mul(hpsz).Add(hpsz) // sending center for within-pool
			for fy := 0; fy < pt.Size.Y; fy++ {
				for fx := 0; fx < pt.Size.X; fx++ {
					for suy := 0; suy < sNuY; suy++ {
						for sux := 0; sux < sNuX; sux++ {
							fwt := float32(1)
							if pt.GaussFull.On {
								sf := math32.Vec2(float32(fx*sNuX+sux), float32(fy*sNuY+suy))
								if pt.GaussFull.Wrap {
									sf.X = edge.WrapMinDist(sf.X, fsz.X, sfctr.X)
									sf.Y = edge.WrapMinDist(sf.Y, fsz.Y, sfctr.Y)
								}
								fwt = efuns.GaussVecDistNoNorm(sf, sfctr, fsig)
							}
							pwt := float32(1)
							if pt.GaussInPool.On {
								sp := math32.Vec2(float32(sux), float32(suy))
								if pt.GaussInPool.Wrap {
									sp.X = edge.WrapMinDist(sp.X, psz.X, spctr.X)
									sp.Y = edge.WrapMinDist(sp.Y, psz.Y, spctr.Y)
								}
								pwt = efuns.GaussVecDistNoNorm(sp, spctr, psig)
							}
							wt := fwt * pwt
							rwt := pt.TopoRange.ProjValue(wt)
							wts.Set(rwt, ruy, rux, fy, fx, suy, sux)
						}
					}
				}
			}
		}
	}
	return nil
}

/////////////////////////////////////////////////////
// SigmoidTopo Wts

// SigmoidTopo has parameters for Gaussian topographic weights or scaling factors
type SigmoidTopo struct {

	// use gaussian topographic weights / scaling values
	On bool

	// gain of sigmoid that determines steepness of curve, in normalized units where entire distance across relevant dimension is 1.0 -- typical useful values range from 0.01 to 0.1
	Gain float32

	// proportion to move gaussian center relative to the position of the receiving unit within its pool: 1.0 = centers span the entire range of the receptive field.  Typically want to use 1.0 for Wrap = true, and 0.8 for false
	CtrMove float32 `default:"0.5,1"`
}

func (gt *SigmoidTopo) Defaults() {
	gt.Gain = 0.05
	gt.CtrMove = 0.5
}

func (gt *SigmoidTopo) ShouldDisplay(field string) bool {
	switch field {
	case "On":
		return true
	default:
		return gt.On
	}
}

// TopoWeightsSigmoid2D sets values in given 4D tensor according to Topo settings.
// wts is shaped with first 2 outer-most dims as Y, X of units within pool
// of recv layer (these are units over which topography is defined)
// and remaing 2D is for sending layer (2D = sender).
func (pt *PoolTile) TopoWeightsSigmoid2D(send, recv *tensor.Shape, wts *tensor.Float32) error {
	if pt.SigFull.Gain == 0 {
		pt.SigFull.Defaults()
	}
	if pt.SigInPool.Gain == 0 {
		pt.SigInPool.Defaults()
	}
	sNuY := send.DimSize(0)
	sNuX := send.DimSize(1)
	rNuY := recv.DimSize(0) // ok if recv is 2D
	rNuX := recv.DimSize(1)
	if recv.NumDims() == 4 {
		rNuY = recv.DimSize(2)
		rNuX = recv.DimSize(3)
	}
	wts.SetShapeSizes(rNuY, rNuX, sNuY, sNuX)

	fsz := math32.Vec2(float32(sNuX-1), float32(sNuY-1)) // full rf size
	hfsz := fsz.MulScalar(0.5)                           // half rf
	fgain := pt.SigFull.Gain * hfsz.X                    // full gain

	psz := math32.Vec2(float32(sNuX), float32(sNuY)) // within-pool rf size
	if sNuX > 1 {
		psz.X -= 1
	}
	if sNuY > 1 {
		psz.Y -= 1
	}
	hpsz := psz.MulScalar(0.5)          // half rf
	pgain := pt.SigInPool.Gain * hpsz.X // pool sigma

	rsz := math32.Vec2(float32(rNuX), float32(rNuY)) // recv units-in-pool size
	if rNuX > 1 {
		rsz.X -= 1
	}
	if rNuY > 1 {
		rsz.Y -= 1
	}
	hrsz := rsz.MulScalar(0.5)
	for ruy := 0; ruy < rNuY; ruy++ {
		for rux := 0; rux < rNuX; rux++ {
			rpos := math32.Vec2(float32(rux), float32(ruy)).Div(hrsz) // 0..2 normalized r unit pos
			sgn := math32.Vec2(1, 1)
			rfpos := rpos.SubScalar(0.5).MulScalar(pt.SigFull.CtrMove).AddScalar(0.5)
			rppos := rpos.SubScalar(0.5).MulScalar(pt.SigInPool.CtrMove).AddScalar(0.5)
			if rpos.X >= 1 { // flip direction half-way through
				sgn.X = -1
				rpos.X = -rpos.X + 1
				rfpos.X = (rpos.X+0.5)*pt.SigFull.CtrMove - 0.5
				rppos.X = (rpos.X+0.5)*pt.SigInPool.CtrMove - 0.5
			}
			if rpos.Y >= 1 {
				sgn.Y = -1
				rpos.Y = -rpos.Y + 1
				rfpos.Y = (rpos.Y+0.5)*pt.SigFull.CtrMove - 0.5
				rfpos.Y = (rpos.Y+0.5)*pt.SigInPool.CtrMove - 0.5
			}
			sfctr := rfpos.Mul(fsz) // sending center for full
			spctr := rppos.Mul(psz) // sending center for within-pool
			for suy := 0; suy < sNuY; suy++ {
				for sux := 0; sux < sNuX; sux++ {
					fwt := float32(1)
					if pt.SigFull.On {
						sf := math32.Vec2(float32(sux), float32(suy))
						sigx := efuns.Logistic(sgn.X*sf.X, fgain, sfctr.X)
						sigy := efuns.Logistic(sgn.Y*sf.Y, fgain, sfctr.Y)
						fwt = sigx * sigy
					}
					pwt := float32(1)
					if pt.SigInPool.On {
						sp := math32.Vec2(float32(sux), float32(suy))
						sigx := efuns.Logistic(sgn.X*sp.X, pgain, spctr.X)
						sigy := efuns.Logistic(sgn.Y*sp.Y, pgain, spctr.Y)
						pwt = sigx * sigy
					}
					wt := fwt * pwt
					rwt := pt.TopoRange.ProjValue(wt)
					wts.Set(rwt, ruy, rux, suy, sux)
				}
			}
		}
	}
	return nil
}

// TopoWeightsSigmoid4D sets values in given 6D tensor according to Topo settings.
// wts is shaped with first 2 outer-most dims as Y, X of units within pool
// of recv layer (these are units over which topography is defined)
// and remaing 2D is for receptive field Size by units within pool size for
// sending layer.
func (pt *PoolTile) TopoWeightsSigmoid4D(send, recv *tensor.Shape, wts *tensor.Float32) error {
	if pt.SigFull.Gain == 0 {
		pt.SigFull.Defaults()
	}
	if pt.SigInPool.Gain == 0 {
		pt.SigInPool.Defaults()
	}
	sNuY := send.DimSize(2)
	sNuX := send.DimSize(3)
	rNuY := recv.DimSize(0) // ok if recv is 2D
	rNuX := recv.DimSize(1)
	if recv.NumDims() == 4 {
		rNuY = recv.DimSize(2)
		rNuX = recv.DimSize(3)
	}
	wts.SetShapeSizes(rNuY, rNuX, pt.Size.Y, pt.Size.X, sNuY, sNuX)

	fsz := math32.Vec2(float32(pt.Size.X*sNuX-1), float32(pt.Size.Y*sNuY-1)) // full rf size
	hfsz := fsz.MulScalar(0.5)                                               // half rf
	fgain := pt.SigFull.Gain * hfsz.X                                        // full gain

	psz := math32.Vec2(float32(sNuX), float32(sNuY)) // within-pool rf size
	if sNuX > 1 {
		psz.X -= 1
	}
	if sNuY > 1 {
		psz.Y -= 1
	}
	hpsz := psz.MulScalar(0.5)          // half rf
	pgain := pt.SigInPool.Gain * hpsz.X // pool sigma

	rsz := math32.Vec2(float32(rNuX), float32(rNuY)) // recv units-in-pool size
	if rNuX > 1 {
		rsz.X -= 1
	}
	if rNuY > 1 {
		rsz.Y -= 1
	}
	hrsz := rsz.MulScalar(0.5)
	for ruy := 0; ruy < rNuY; ruy++ {
		for rux := 0; rux < rNuX; rux++ {
			rpos := math32.Vec2(float32(rux), float32(ruy)).Div(hrsz) // 0..2 normalized r unit pos
			sgn := math32.Vec2(1, 1)
			rfpos := rpos.SubScalar(0.5).MulScalar(pt.SigFull.CtrMove).AddScalar(0.5)
			rppos := rpos.SubScalar(0.5).MulScalar(pt.SigInPool.CtrMove).AddScalar(0.5)
			if rpos.X >= 1 { // flip direction half-way through
				sgn.X = -1
				rpos.X = -rpos.X + 1
				rfpos.X = (rpos.X+0.5)*pt.SigFull.CtrMove - 0.5
				rppos.X = (rpos.X+0.5)*pt.SigInPool.CtrMove - 0.5
			}
			if rpos.Y >= 1 {
				sgn.Y = -1
				rpos.Y = -rpos.Y + 1
				rfpos.Y = (rpos.Y+0.5)*pt.SigFull.CtrMove - 0.5
				rfpos.Y = (rpos.Y+0.5)*pt.SigInPool.CtrMove - 0.5
			}
			sfctr := rfpos.Mul(fsz) // sending center for full
			spctr := rppos.Mul(psz) // sending center for within-pool
			for fy := 0; fy < pt.Size.Y; fy++ {
				for fx := 0; fx < pt.Size.X; fx++ {
					for suy := 0; suy < sNuY; suy++ {
						for sux := 0; sux < sNuX; sux++ {
							fwt := float32(1)
							if pt.SigFull.On {
								sf := math32.Vec2(float32(fx*sNuX+sux), float32(fy*sNuY+suy))
								sigx := efuns.Logistic(sgn.X*sf.X, fgain, sfctr.X)
								sigy := efuns.Logistic(sgn.Y*sf.Y, fgain, sfctr.Y)
								fwt = sigx * sigy
							}
							pwt := float32(1)
							if pt.SigInPool.On {
								sp := math32.Vec2(float32(sux), float32(suy))
								sigx := efuns.Logistic(sgn.X*sp.X, pgain, spctr.X)
								sigy := efuns.Logistic(sgn.Y*sp.Y, pgain, spctr.Y)
								pwt = sigx * sigy
							}
							wt := fwt * pwt
							rwt := pt.TopoRange.ProjValue(wt)
							wts.Set(rwt, ruy, rux, fy, fx, suy, sux)
						}
					}
				}
			}
		}
	}
	return nil
}
