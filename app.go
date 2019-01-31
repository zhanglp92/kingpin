package kingpin

import (
	"strconv"
	"time"
)

// Application ...
type Application struct {
	name, help string
	values     []string
}

// Action ...
func (a *Application) Action(action Action) *Application {
	return a
}

// Flag ...
func (a *Application) Flag(name, help string) *Application {
	return &Application{name: name, help: help}
}

// Default ...
func (a *Application) Default(values ...string) *Application {
	return &Application{values: values}
}

func (a *Application) String() (target *string) {
	target = new(string)
	if len(a.values) <= 0 {
		return
	}
	*target = a.values[0]

	return
}

// StringVar ...
func (a *Application) StringVar(target *string) {
	if len(a.values) > 0 {
		*target = a.values[0]
	}
}

// Bool ...
func (a *Application) Bool() (target *bool) {
	target = new(bool)

	if len(a.values) <= 0 {
		return
	}
	*target, _ = strconv.ParseBool(a.values[0])

	return
}

// Int ...
func (a *Application) Int() (target *int) {
	target = new(int)

	if len(a.values) <= 0 {
		return
	}
	n, _ := strconv.ParseInt(a.values[0], 10, 0)
	*target = int(n)

	return
}

// Duration ...
func (a *Application) Duration() (target *time.Duration) {
	target = new(time.Duration)

	if len(a.values) <= 0 {
		return
	}
	*target, _ = time.ParseDuration(a.values[0])

	return
}
