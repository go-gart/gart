package container

import "fmt"

type (
	Instance interface{}

	Concrete interface{}
)

type Container struct {
	aliases map[string]string

	resolved map[string]bool

	bindings map[string]Binding
	shares   map[string]bool

	instances map[string]Instance
}

func NewContainer() *Container {
	instance := &Container{}

	instance.aliases = make(map[string]string)
	instance.bindings = make(map[string]Binding)
	instance.resolved = make(map[string]bool)
	instance.shares = make(map[string]bool)

	instance.instances = make(map[string]Instance)

	return instance
}

func (c *Container) Empty() bool {
	panic("implement me")
}

func (c *Container) Size() int {
	panic("implement me")
}

func (c *Container) Clear() {
	panic("implement me")
}

func (c *Container) Values() []interface{} {
	panic("implement me")
}

func (c *Container) Bind(abstract string, concrete Binding) {

	c.dropStaleInstance(abstract)

	c.bindings[abstract] = concrete
}

func (c *Container) Share(abstract string) {
	c.shares[abstract] = true
}

func (c *Container) Singleton(abstract string, builder Binding) {
	c.Bind(abstract, builder)
	c.Share(abstract)
}

func (c *Container) Instance(abstract string, instance interface{}) {
	c.instances[abstract] = instance
}

func (c *Container) dropStaleInstance(abstract string) {
	c.bindings[abstract] = nil
	c.shares[abstract] = false
	c.aliases[abstract] = ""
}

//	Alias a type to a different name.
//	为 abstract 提供 alias 作为别名
func (c *Container) Alias(abstract string, alias string) {
	c.aliases[alias] = abstract
}

func (c *Container) Make(abstract string, params ...interface{}) interface{} {
	abstract = c.getAlias(abstract)

	if c.instances[abstract] != nil {
		return c.instances[abstract]
	}

	concrete := c.getConcrete(abstract)

	obj := c.Build(concrete, params)

	if c.IsShared(abstract) {
		c.instances[abstract] = obj
	}

	return obj
}

func (c *Container) getAlias(abstract string) string {

	for len(c.aliases[abstract]) > 0 {
		abstract = c.aliases[abstract]
	}

	return abstract
}

func (c *Container) getConcrete(abstract string) Binding {

	binding := c.bindings[abstract]

	if binding == nil {
		panic(fmt.Sprintf(`unregisted binding [%s]`, abstract))
	}

	return binding
}

func (c *Container) Build(binding Binding, params []interface{}) interface{} {
	return binding.Build(c)
}

func (c *Container) isBound() bool {
	return false
}

func (c *Container) BindInstance(string, interface{}) {
	panic("implement me")
}

func (c *Container) IsBound(string) bool {
	panic("implement me")
}

func (c *Container) IsResolved(string) bool {
	panic("implement me")
}

func (c *Container) IsShared(abstract string) bool {
	return c.shares[abstract]
}

func (c *Container) IsAlias(string) bool {
	panic("implement me")
}

func (c *Container) Resolving(string) {
	panic("implement me")
}

func (c *Container) Resolved(string) {
	panic("implement me")
}
