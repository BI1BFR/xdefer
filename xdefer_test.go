package xdefer

import (
	"fmt"
	"testing"
)

func printFn(s string) func() {
	return func() {
		fmt.Println(s)
	}
}

func add(d *Defer) {
	d.Defer(printFn("from add()"))
}

func TestAdd(t *testing.T) {
	var xd Defer
	defer xd.Exec()

	xd.Defer(printFn("before add()"))
	add(&xd)
	xd.Defer(printFn("after add()"))
	fmt.Println("end of TestAdd()")
}

func create() (xd Defer) {
	xd.Defer(printFn("from create()"))
	return
}

func TestCreate(t *testing.T) {
	xd := create()
	defer xd.Exec()

	xd.Defer(printFn("from TestCreate()"))
	fmt.Println("end of TestCreate()")
}
