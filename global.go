package kingpin

// Flag ...
func Flag(name, help string) *Application {
	return &Application{name: name, help: help}
}
