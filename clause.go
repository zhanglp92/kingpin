package kingpin

// desc: 封装flag

import (
	"strconv"
)

// Clause ...
type Clause struct {
	name, help string
	values     []string
}

// Flag ...
func Flag(name, help string) *Clause {
	return &Clause{name: name, help: help}
}

// Default ...
func (c *Clause) Default(values ...string) *Clause {
	return &Clause{values: values}
}

func (c *Clause) String() (target *string) {
	target = new(string)
	if len(c.values) <= 0 {
		return
	}
	*target = c.values[0]

	return
}

// Bool ...
func (c *Clause) Bool() (target *bool) {
	target = new(bool)

	if len(c.values) <= 0 {
		return
	}
	*target, _ = strconv.ParseBool(c.values[0])

	return
}
