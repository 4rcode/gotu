package gotu

import (
	"fmt"
	"strings"
)

// [Assert] is responsible for logging an error in response to a false
// assertion.
//
// When the assertion is true, Assert simply returns true.
//
// When the assertion is false, Assert uses the supplied context arguments to
// report an error message, then it returns false.
type Assert func(bool, ...interface{}) bool

// [AssertWith] is responsible for building a new [Assert], given some context
// arguments.
//
// By default, the only necessary context argument is an instance of type
// [testing.TB]
//
// However, any type with the following optional methods can be provided:
//
//	Helper()
//	Error(...interface{})
//
// Alternatively, functions with the following signatures can be provided:
//
//	/*
//	 * helper functions (e.g. t.Helper)
//	 */
//
//	    func()
//
//	/*
//	 * testing.TB and log functions (e.g. t.Fatal, t.Skip, or log.Println)
//	 */
//
//	    func(...interface{})
//
//	/*
//	 * print functions (e.g. fmt.Println)
//	 */
//
//	    func(...interface{}) (int, error)
func AssertWith(context ...interface{}) Assert {
	return func(assertion bool, arguments ...interface{}) bool {
		if assertion {
			return assertion
		}

		var text = format(arguments...)

		for _, ctx := range context {
			if helper, _ := ctx.(interface {
				Helper()
			}); helper != nil {
				helper.Helper()
			}

			if logger, _ := ctx.(interface {
				Error(...interface{})
			}); logger != nil {
				logger.Error(text)
			}

			switch logger := ctx.(type) {
			case func():
				logger()
			case func(...interface{}):
				logger(text)
			case func(...interface{}) (int, error):
				logger(text)
			}
		}

		return assertion
	}
}

func format(arguments ...interface{}) string {
	var length = len(arguments)

	if length < 1 {
		return "\n\nfalse assertion\n"
	}

	if length == 1 {
		return "\n\n" + fmt.Sprintln(arguments...)
	}

	var w strings.Builder

	var last = length - 1
	var details = arguments[1:last]

	fmt.Fprintf(&w, "\n\nPROVIDED:\n%10s%#v\nEXPECTED:\n", "", arguments[0])

	if len(details) > 0 {
		fmt.Fprintf(&w, "%10s", "")
		fmt.Fprintln(&w, details...)
	}

	fmt.Fprintf(&w, "%10s%#v\n", "", arguments[last])

	return w.String()
}
