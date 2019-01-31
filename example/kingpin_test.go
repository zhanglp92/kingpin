package example

import (
	"testing"

	"github.com/zhanglp92/kingpin"
)

func TestKingpin(t *testing.T) {
	a := kingpin.Flag("name-test", "help-test").Default("true").Bool()
	b := kingpin.Flag("name-test", "help-test").Default("true").String()

	t.Log("Bool: ", a, *a)
	t.Log("String: ", b, *b)
}
