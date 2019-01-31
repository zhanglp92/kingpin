package clause

// desc: 封装flag

import (
	"strconv"

	"github.com/zhanglp92/kingpin/global"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Clause ...
type Clause struct {
	c *kingpin.FlagClause

	name, help string
	values     []string
}

// Flag ...
func Flag(name, help string) *Clause {
	c := &Clause{name: name, help: help}

	if global.Kingpin() {
		c.c = kingpin.Flag(name, help)
	}
	return c
}

// Default ...
func (c *Clause) Default(values ...string) *Clause {
	cc := &Clause{values: values}

	if global.Kingpin() {
		cc.c = c.c.Default(values...)
	}
	return cc
}

func (c *Clause) String() (target *string) {
	if global.Kingpin() {
		return c.c.String()
	}

	target = new(string)
	if len(c.values) <= 0 {
		return
	}
	*target = c.values[0]

	return
}

// Bool ...
func (c *Clause) Bool() (target *bool) {
	if global.Kingpin() {
		return c.c.Bool()
	}

	target = new(bool)

	if len(c.values) <= 0 {
		return
	}
	*target, _ = strconv.ParseBool(c.values[0])

	return
}
