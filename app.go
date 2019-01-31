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

func (a *Application) getValues() (vals []string) {
	defer func() {
		if len(vals) <= 0 {
			vals = []string{""}
		}
	}()

	if vals = flagsVal[a.name]; len(vals) <= 0 {
		return a.values
	}
	return
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
	a.values = values
	return a
}

func (a *Application) String() (target *string) {
	target = new(string)
	a.StringVar(target)

	return
}

// StringVar ...
func (a *Application) StringVar(target *string) {
	*target = a.getValues()[0]
}

// Bool ...
func (a *Application) Bool() (target *bool) {
	target = new(bool)
	*target, _ = strconv.ParseBool(a.getValues()[0])

	return
}

// Int ...
func (a *Application) Int() (target *int) {
	target = new(int)

	n, _ := strconv.ParseInt(a.getValues()[0], 10, 0)
	*target = int(n)

	return
}

// Duration ...
func (a *Application) Duration() (target *time.Duration) {
	target = new(time.Duration)

	*target, _ = time.ParseDuration(a.getValues()[0])

	return
}
