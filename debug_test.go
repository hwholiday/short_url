package main

import (
	"testing"
)

func TestA(t *testing.T) {
	var a1 = 0
	a1 |= 1 << 2 //0100
	t.Log(2222)
	t.Error(a1)
	t.Error(1)
}
