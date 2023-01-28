package help

import "testing"

func TestCondition(t *testing.T) {
	t.Log(If(1 > 2, "one", "two"))
}
