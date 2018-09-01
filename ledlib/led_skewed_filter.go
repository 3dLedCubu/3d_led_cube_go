package ledlib

import "math"

const T = 4

type LedSkewedFilter struct {
	canvas ILedCanvas
	cube   LedCube
	yt     float64
	zt     float64
	ys     float64
	zs     float64
	yc     float64
	zc     float64
}

func NewLedSkewedFilter(canvas ILedCanvas) ILedCanvas {
	return &LedSkewedFilter{canvas, NewLedCube(), 0, 0, 0, 0, 0, 0}
}

func (f *LedSkewedFilter) PreShow() {
	f.cube.Clear()
	f.canvas.PreShow()
	f.yt += 0.02 * 0.15
	f.zt += 0.02 * 0.15
	f.ys = math.Sin(f.yt * 3.14 * T)
	f.yc = math.Cos(f.yt * 3.14 * T)
	f.zs = math.Sin(f.zt * 3.14 * T)
	f.zc = math.Cos(f.zt * 3.14 * T)
}

func (f *LedSkewedFilter) Show(c LedCube) {
	var dx float64 = LedWidth / 2.0
	var dy float64 = LedHeight / 4.0 * 3
	var dz float64 = LedDepth / 2.0

	c.ConcurrentForEach(func(x, y, z int, c Color32) {
		//        if 'y' in self.axis:
		xx, yy, zz := ((float64(x)-dx)*f.yc+(float64(z)-dz)*f.ys)+dx, float64(y), (-(float64(x)-dx)*f.ys+(float64(z)-dz)*f.yc)+dz

		xx, yy, zz = ((float64(x)-dx)*f.zc+(float64(y)-dy)*f.zs)+dx, (-(float64(x)-dx)*f.zs+(float64(y)-dy)*f.zc)+dy, float64(z)
		f.cube.SetAt(int(math.Round(xx)), int(math.Round(yy)), int(math.Round(zz)), c)
	})
	f.canvas.Show(f.cube)
}
