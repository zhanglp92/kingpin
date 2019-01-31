package kingpin

// desc: 封装flag

import (
	"strconv"

	sk "gopkg.in/alecthomas/kingpin.v2"
)

// Clause ...
type Clause struct {
	c *sk.FlagClause

	name, help string
	values     []string
}

// Flag ...
func Flag(name, help string) *Clause {
	c := &Clause{name: name, help: help}

	if Kingpin() {
		c.c = sk.Flag(name, help)
	}
	return c
}

// Default ...
func (c *Clause) Default(values ...string) *Clause {
	cc := &Clause{values: values}

	if Kingpin() {
		cc.c = c.c.Default(values...)
	}
	return cc
}

func (c *Clause) String() (target *string) {
	if Kingpin() {
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
	if Kingpin() {
		return c.c.Bool()
	}

	target = new(bool)

	if len(c.values) <= 0 {
		return
	}
	*target, _ = strconv.ParseBool(c.values[0])

	return
}
