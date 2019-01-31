package kingpin

import "strconv"

// Application ...
type Application struct {
	name, help string
	values     []string
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

// Bool ...
func (a *Application) Bool() (target *bool) {
	target = new(bool)

	if len(a.values) <= 0 {
		return
	}
	*target, _ = strconv.ParseBool(a.values[0])

	return
}