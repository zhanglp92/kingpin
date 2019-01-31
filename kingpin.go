package kingpin

import (
	"github.com/zhanglp92/kingpin/clause"
	"github.com/zhanglp92/kingpin/global"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Parse ...
func Parse() {
	if global.Kingpin() {
		kingpin.Parse()
	}
}

// Flag ...
func Flag(name, help string) *clause.Clause {
	return clause.Flag(name, help)
}
