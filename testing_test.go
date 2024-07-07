// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/4/25

package fst

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func newMyTesting(t *testing.T) *myTesting {
	return &myTesting{
		t: t,
	}
}

var _ Testing = (*myTesting)(nil)
var _ Helper = (*myTesting)(nil)

type myTesting struct {
	t          *testing.T
	runFail    bool
	msg        string
	want       bool
	lastCaller string
}

func (m *myTesting) Helper() {
	if m.lastCaller != "" && m.runFail != m.want {
		m.check()
	}
	m.runFail = false
	m.msg = ""
	_, file, lineNo, _ := runtime.Caller(2)
	m.lastCaller = fmt.Sprintf("%s:%d", filepath.Base(file), lineNo)
}

func (m *myTesting) check() {
	m.t.Helper()
	if m.want {
		if m.runFail {
			m.t.Fatalf("%s want success, but not", m.lastCaller)
		}
	} else {
		if !m.runFail {
			m.t.Fatalf("%s want fail, but not", m.lastCaller)
		}
	}
}

func (m *myTesting) Fatalf(format string, args ...any) {
	m.runFail = true
	m.msg = fmt.Sprintf(format, args...)
}

func (m *myTesting) Success(fn func(t Testing)) {
	m.t.Helper()
	m.want = true
	fn(m)
	m.check()
}

func (m *myTesting) Fail(fn func(t Testing)) {
	m.want = false
	fn(m)
	m.check()
}
