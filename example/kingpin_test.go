package example

import (
	"testing"

	"github.com/zhanglp92/kingpin"
)

func TestKingpin(t *testing.T) {

	kingpin.SetVal("n1", "n1", "nn1")
	kingpin.SetVal("n2", "n2", "nn2")

	a := kingpin.Flag("n1", "help-test").Default("true").Bool()
	b := kingpin.Flag("n2", "help-test").Default("true").String()
	c := kingpin.Flag("n3", "help-test").Default("3.46608ms").Duration()

	t.Log("Flags", kingpin.GetFlags())

	t.Log("Bool: ", a, *a)
	t.Log("String: ", b, *b)
	t.Log("Du: ", c, *c)
}
