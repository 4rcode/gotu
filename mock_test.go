package gotu_test

import (
	"fmt"
	"strings"
)

type _mock struct {
	strings.Builder
	helper int
}

func (t *_mock) Error(args ...interface{}) {
	fmt.Fprintln(&t.Builder, args...)
}

func (t *_mock) Helper() {
	t.helper++
}
