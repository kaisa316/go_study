package world

type world struct {
	Name string
	Size int
}

/**
* 工厂方法
 */
func NewWorld() *world {
	w := new(world)
	w.Name = "china"
	w.Size = 960
	return w
}