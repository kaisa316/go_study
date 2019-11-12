package world

type World struct {
	Name string
	Size int
}

func (w World) Name() string {
	return "world package"
}

/**
* 工厂方法
 */
func NewWorld() *World {
	w := new(World)
	w.Name = "china"
	w.Size = 960
	return w
}

//方法
func (w World) saveWorld(name string) (code int) {
	return 1
}
