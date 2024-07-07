// Copyright(C) 2023 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2023/9/8

package fst

import (
	"slices"
	"strings"
)

func Equal[T any](t Testing, expected T, actual T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !equal(expected, actual) {
		t.Fatalf("Not equal: \nrunFail:  %#v\n actual: %#v", expected, actual)
	}
}

func NotEqual[T any](t Testing, expected T, actual T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if equal(expected, actual) {
		t.Fatalf("Should not equal: %#v", actual)
	}
}

func Error(t Testing, err error) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if err == nil {
		t.Fatalf("An error is expected but runFail nil.")
	}
}

func NoError(t Testing, err error) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if err != nil {
		t.Fatalf("Received unexpected error:\n%+v", err)
	}
}

func True(t Testing, got bool) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !got {
		t.Fatalf("Should be true")
	}
}

func False(t Testing, got bool) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if got {
		t.Fatalf("Should be false")
	}
}

func Nil(t Testing, got any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !isNil(got) {
		t.Fatalf("%#v runFail not nil, but should have", got)
	}
}

func NotNil(t Testing, got any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if isNil(got) {
		t.Fatalf("%#v runFail nil, but should not have", got)
	}
}

func Empty(t Testing, got any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}

	if !isEmpty(got) {
		t.Fatalf("Should be empty, but was %v", got)
	}
}

func NotEmpty(t Testing, got any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if isEmpty(got) {
		t.Fatalf("Should NOT be empty, but was %v", got)
	}
}

func StringContains(t Testing, s string, substr string) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !strings.Contains(s, substr) {
		t.Fatalf("%#v does not substr %#v", s, substr)
	}
}

func StringNotContains(t Testing, s string, substr string) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if strings.Contains(s, substr) {
		t.Fatalf("%#v should not substr %#v", s, substr)
	}
}

func SliceContains[S ~[]E, E comparable](t Testing, values S, item E) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !slices.Contains(values, item) {
		t.Fatalf("%#v does not contains %#v", values, item)
	}
}

func SliceNotContains[S ~[]E, E comparable](t Testing, values S, item E) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if slices.Contains(values, item) {
		t.Fatalf("%#v should not contains %#v", values, item)
	}
}

func SamePtr[T any](t Testing, expected T, actual T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !samePointers(expected, actual) {
		t.Fatalf("Not same: \n"+
			"expected: %p %#v\n"+
			"actual  : %p %#v", expected, expected, actual, actual)
	}
}

func NotSamePtr[T any](t Testing, expected T, actual T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if samePointers(expected, actual) {
		t.Fatalf("Expected and actual point to the same object: %p %#v", expected, expected)
	}
}
