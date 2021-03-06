package ledlib

import (
	"3d_led_cube_go/ledlib/util"
	"sync"
)

/*
	LedObjectとそのlifetimeを管理する
	以下の機能を提供
	- LedObjectとlifetimeの登録(複数可能)
	- lifetime終了チェック
	- すべてのObjectの列挙&Draw(LedManagedObject) on goroutine
	- マルチスレッドアクセス可能
	- Not Singleton
	-

	LedObject IntefaceにIsOutOfCanvasを実装
	 OutOfCanvasなObjectは、管理から外す処理


*/

type LedManagedObject interface {
	IsExpired() bool
	Draw(cube util.Image3D)
}

type FilterObjects struct {
	canvas  LedCanvas
	objects []LedManagedObject
	cube    util.Image3D
	mutex   *sync.Mutex
}

func NewFilterObjects(canvas LedCanvas) *FilterObjects {

	filter := FilterObjects{}
	filter.canvas = canvas
	filter.objects = make([]LedManagedObject, 0)
	filter.cube = NewLedImage3D()
	filter.mutex = &sync.Mutex{}

	return &filter
}

func (l *FilterObjects) Len() int {
	return len(l.objects)
}

func (f *FilterObjects) Append(obj LedManagedObject) {
	f.objects = append(f.objects, obj)
}

func (f *FilterObjects) ConcurrentShow(cube util.ImmutableImage3D, param LedCanvasParam) {
}

func (f *FilterObjects) Show(cube util.ImmutableImage3D, param LedCanvasParam) {

	f.cube = cube.Copy()

	actives := make([]int, 0, len(f.objects))

	util.ConcurrentEnum(0, len(f.objects), func(i int) {
		if !f.objects[i].IsExpired() {
			f.mutex.Lock()
			actives = append(actives, i)
			f.mutex.Unlock()
			f.objects[i].Draw(f.cube)
		}
	})

	newobjects := make([]LedManagedObject, len(actives))
	for i, target := range actives {
		newobjects[i] = f.objects[target]
	}
	f.objects = newobjects
	f.canvas.Show(f.cube, param)
}
