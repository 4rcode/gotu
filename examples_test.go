package gotu_test

import (
	"fmt"

	"github.com/4rcode/gotu"
)

var t = fmt.Println

func Example() {
	var assert = gotu.AssertWith(t)

	var (
		a = 1
		b = 2
	)

	assert(a == b, a, b)
	// output:
	// PROVIDED:
	//           1
	// EXPECTED:
	//           2
}

func Example_withDetails() {
	var assert = gotu.AssertWith(t)

	var (
		a = 1
		b = 2
	)

	assert(a > b, a, "greater than", b)
	// output:
	// PROVIDED:
	//           1
	// EXPECTED:
	//           greater than
	//           2
}

func Example_withMessageOnly() {
	var assert = gotu.AssertWith(t)

	var (
		a = 1
		b = 2
	)

	assert(a == b, "incorrect value")
	// output:
	// incorrect value
}

func Example_withNoArguments() {
	var assert = gotu.AssertWith(t)

	var (
		a = 1
		b = 2
	)

	assert(a == b)
	// output:
	// false assertion
}
