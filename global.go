package kingpin

import "fmt"

var (
	// 自定义flag
	flags = make(map[string]*Application)

	// 自定义flag的值
	flagsVal = make(map[string][]string)
)

// Flag 添加flag
func Flag(name, help string) *Application {
	return addFlag(&Application{name: name, help: help})
}

// SetVal 自定义flag值
func SetVal(name string, values ...string) {
	flagsVal[name] = values
}

// GetFlags ...
func GetFlags() map[string][]string {
	tar := make(map[string][]string)
	for name, app := range flags {
		tar[name] = app.values
	}

	return tar
}

func addFlag(src *Application) *Application {
	if src == nil {
		return nil
	}

	if flags[src.name] != nil {
		panic(fmt.Sprintf("flag %v is exists", src.name))
	}

	flags[src.name] = src
	return src
}
