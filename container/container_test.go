package container

import (
	"math/rand"
	"testing"
)

func TestContainer_Bind(t *testing.T) {

	data := [][]interface{}{{
		"name", ClosureBinding(func(app *Container) interface{} {
			return `huiren`
		}),
		`huiren`,
	}}

	c := NewContainer()

	for i := range data {

		abstract := data[i][0].(string)
		binding := data[i][1].(Binding)
		expect := data[i][2]

		c.Bind(abstract, binding)

		result := c.Make(abstract)

		if result != expect {
			t.Errorf("error")
		}
	}
}

func TestContainer_Singleton(t *testing.T) {

	c := NewContainer()

	c.Singleton(`notRandom`, ClosureBinding(func(c *Container) interface{} {
		return rand.Int63()
	}))

	num1 := c.Make(`notRandom`)
	num2 := c.Make(`notRandom`)

	if num1 != num2 {
		t.Error("error", num1, num2)
	}
}

func TestContainer_Alias(t *testing.T) {
	c := NewContainer()

	c.Bind(`name`, ClosureBinding(func(c *Container) interface{} {
		return `huiren`
	}))

	c.Alias(`name`, `alias_name`)

	name := c.Make(`name`)
	aliasName := c.Make(`alias_name`)

	if name != aliasName {
		t.Error("error", name, aliasName)
	}
}
