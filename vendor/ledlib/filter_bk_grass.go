package ledlib

import (
	"ledlib/util"
	"time"
)

type FilterBkGrass struct {
	filterObjects *FilterObjects
}

func NewFilterBkGrass(canvas LedCanvas) LedCanvas {
	filter := FilterBkGrass{}
	filter.filterObjects = NewFilterObjects(canvas)

	duration := 100 * time.Millisecond
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		"/asset/image/grass/grass1.png", 0, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		"/asset/image/grass/grass2.png", 1, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		"/asset/image/grass/grass1.png", 2, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		"/asset/image/grass/grass2.png", 3, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		"/asset/image/grass/grass1.png", 4, duration))
	filter.filterObjects.Append(NewObjectScrolledBitmap(
		"/asset/image/grass/grass3.png", 5, duration))

	return &filter
}

func (f *FilterBkGrass) Show(c util.ImmutableImage3D, param LedCanvasParam) {
	f.filterObjects.Show(c, param)
}
