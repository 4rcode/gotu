package gotu_test

import (
	"testing"

	"github.com/4rcode/gotu"
)

func TestAssertWith(t *testing.T) {
	var assert = gotu.AssertWith(t)

	type args []interface{}

	for tc, td := range map[string]struct {
		test bool
		args
		helper   int
		expected string
	}{
		"trueNoArgs": {
			true, nil, 0, ""},
		"trueMsgOnly": {
			true, args{"msg"}, 0, ""},
		"trueValOnly": {
			true, args{1, "abc"}, 0, ""},
		"trueAllArgs": {
			true, args{1, "a string", "equal to", "abc"}, 0, ""},
		"falseNoArgs": {
			false, nil, 1, `

false assertion

`},
		"falseMsgOnly": {
			false, args{"msg"}, 1, `

msg

`},
		"falseValOnly": {
			false, args{1, "abc"}, 1, `

PROVIDED:
          1
EXPECTED:
          "abc"

`},
		"falseAllArgs": {
			false, args{1, "a string", "equal to", "abc"}, 1, `

PROVIDED:
          1
EXPECTED:
          a string equal to
          "abc"

`},
	} {
		t.Log(tc)

		var mock _mock
		var test = gotu.AssertWith(&mock)
		var result = test(td.test, td.args...)
		var provided = mock.String()

		assert(result == td.test, result, td.test)
		assert(mock.helper == td.helper, mock.helper, td.helper)
		assert(provided == td.expected, provided, td.expected)

		mock = _mock{}
		test = gotu.AssertWith(mock.Helper, mock.Error)
		result = test(td.test, td.args...)
		provided = mock.String()

		assert(result == td.test, result, td.test)
		assert(mock.helper == td.helper, mock.helper, td.helper)
		assert(provided == td.expected, provided, td.expected)
	}
}
