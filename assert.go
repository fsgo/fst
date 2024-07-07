// Copyright(C) 2023 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2023/9/8

package fst

import (
	"slices"
	"strings"
)

func Equal[T any](t Testing, got T, want T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !equal(got, want) {
		t.Fatalf("Not equal: \nrunFail:  %#v\n want: %#v", got, want)
	}
}

func NotEqual[T any](t Testing, got T, want T) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if equal(got, want) {
		t.Fatalf("Should not equal: %#v", want)
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

func StringContains(t Testing, got string, want string) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !strings.Contains(got, want) {
		t.Fatalf("%#v does not contains %#v", got, want)
	}
}

func StringNotContains(t Testing, got string, want string) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if strings.Contains(got, want) {
		t.Fatalf("%#v should not contains %#v", got, want)
	}
}

func SliceContains[S ~[]E, E comparable](t Testing, got S, want E) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if !slices.Contains(got, want) {
		t.Fatalf("%#v does not contains %#v", got, want)
	}
}

func SliceNotContains[S ~[]E, E comparable](t Testing, got S, want E) {
	if h, ok := t.(Helper); ok {
		h.Helper()
	}
	if slices.Contains(got, want) {
		t.Fatalf("%#v should not contains %#v", got, want)
	}
}
