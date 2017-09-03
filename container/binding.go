package container

type Binding interface {
	//	实例化实现
	Build(*Container, ...interface{}) interface{}
}

type ClosureBinding func(*Container) interface{}

func (b ClosureBinding) Build(c *Container, params ...interface{}) interface{} {
	return b(c)
}
