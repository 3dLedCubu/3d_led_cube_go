package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"time"
)

type FilterBkMountain struct {
	filterObjects     *FilterObjects
	filterObjectsSnow *FilterObjects
}

func NewFilterBkMountain(canvas LedCanvas) LedCanvas {
	filter := FilterBkMountain{}
	filter.filterObjects = NewFilterObjects(canvas)
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/mountain/mountain2.png", 6, 300*time.Millisecond, true))
	filter.filterObjects.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/mountain/mountain1.png", 7, 300*time.Millisecond, true))
	filter.filterObjectsSnow = NewFilterObjects(canvas)
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/mountain/mountain2-s.png", 6, 300*time.Millisecond, true))
	filter.filterObjectsSnow.Append(NewObjectHorizontalScrolledBitmap(
		"/asset/image/mountain/mountain1-s.png", 7, 300*time.Millisecond, true))

	return &filter
}

func (f *FilterBkMountain) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	if param.HasEffect("filter-bk-snows") {
		f.filterObjectsSnow.Show(c, param)
	} else {
		f.filterObjects.Show(c, param)
	}

}
