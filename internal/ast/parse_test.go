package ast

import (
	"testing"
)

func TestParse(t *testing.T) {
	script := "int age = 45+2; age= 20; age+10*2;"
	tree, err := Parse(script)
	if err != nil {
		t.FailNow()
	}
	Dump(tree, "")
}
