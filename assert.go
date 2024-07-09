// Copyright(C) 2023 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2023/9/8

package fst

import (
	"cmp"
	"errors"
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

func Less[T cmp.Ordered](t Testing, x T, y T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if cmp.Compare(x, y) != -1 {
		t.Fatalf(`"%v" is not less than "%v"`, x, y)
	}
}

func LessOrEqual[T cmp.Ordered](t Testing, x T, y T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if cmp.Compare(x, y) == 1 {
		t.Fatalf(`"%v" is not less than or equal to "%v"`, x, y)
	}
}

func Greater[T cmp.Ordered](t Testing, x T, y T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if cmp.Compare(x, y) != 1 {
		t.Fatalf(`"%v" is not greater than "%v"`, x, y)
	}
}

func GreaterOrEqual[T cmp.Ordered](t Testing, x T, y T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if cmp.Compare(x, y) == -1 {
		t.Fatalf(`"%v" is not greater than or equal to "%v"`, x, y)
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

func ErrorIs(t Testing, err error, target error) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if errors.Is(err, target) {
		return
	}

	var expectedText string
	if target != nil {
		expectedText = target.Error()
	}

	chain := buildErrorChainString(err)
	t.Fatalf("Target error should be in err chain:\n"+
		"expected: %q\n"+
		"in chain: %s", expectedText, chain,
	)
}

func NotErrorIs(t Testing, err error, target error) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !errors.Is(err, target) {
		return
	}

	var expectedText string
	if target != nil {
		expectedText = target.Error()
	}

	chain := buildErrorChainString(err)
	t.Fatalf("Target error should not be in err chain:\n"+
		"expected: %q\n"+
		"in chain: %s", expectedText, chain,
	)
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
		t.Fatalf("Expected nil, but got: %#v", got)
	}
}

func NotNil(t Testing, got any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if isNil(got) {
		t.Fatalf("Expected value not to be nil")
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

func SamePtr(t Testing, expected any, actual any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !samePointers(expected, actual) {
		t.Fatalf("Not same: \n"+
			"expected: %p %#v\n"+
			"actual  : %p %#v", expected, expected, actual, actual)
	}
}

func NotSamePtr(t Testing, expected any, actual any) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if samePointers(expected, actual) {
		t.Fatalf("Expected and actual point to the same object: %p %#v", expected, expected)
	}
}
