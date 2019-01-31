package kingpin

import (
	sk "gopkg.in/alecthomas/kingpin.v2"
)

// Parse ...
func Parse() {
	if Kingpin() {
		sk.Parse()
	}
}
