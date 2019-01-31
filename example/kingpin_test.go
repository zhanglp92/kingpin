package example

import (
	"testing"

	"github.com/zhanglp92/kingpin"
)

func TestKingpin(t *testing.T) {
	a := kingpin.Flag("name-test", "help-test").Default("true").Bool()
	b := kingpin.Flag("name-test", "help-test").Default("true").String()
	c := kingpin.Flag("name-test", "help-test").Default("3.46608ms").Duration()

	t.Log("Bool: ", a, *a)
	t.Log("String: ", b, *b)
	t.Log("Du: ", c, *c)
}
