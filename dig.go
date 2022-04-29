package main

import (
	"fmt"
	"go.uber.org/dig"
)

type Config struct {
	Loop int
}

func (c *Config) Foo() {
	fmt.Println("foo")
}

func main()  {
	c := dig.New()
	_ = c.Provide(func() *Config {
		return new(Config)
	})
	_ = c.Invoke(func(c *Config) {
		c.Foo()
	})
}
