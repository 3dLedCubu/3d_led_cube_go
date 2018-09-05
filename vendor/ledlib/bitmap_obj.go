package ledlib

import (
	"image/png"
	"ledlib/util"
	"log"
	"os"
)

type BitmapObj struct {
	cube util.CubeImage
}

func NewBitmapObj(paths []string) *BitmapObj {
	bmp := BitmapObj{NewLedCubeImage()}
	bmp.load(paths)
	return &bmp
}

func (b *BitmapObj) load(paths []string) {

	for z, path := range paths {
		if path == "" {
			continue
		}
		reader, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer reader.Close()
		m, err := png.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		width, height := m.Bounds().Dx(), m.Bounds().Dy()
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if m == nil {
					continue
				}
				b.cube.SetAt(x, y, z, util.NewFromColorColor(m.At(x, y)))
			}
		}
	}
}

func (b *BitmapObj) Draw(canvas ILedCanvas) {
	canvas.Show(b.cube)
}
