package country

type China struct {
	name string
	area float32
	flag string
}

//Name get name
func (c China) Name() string {
	return c.name
}
