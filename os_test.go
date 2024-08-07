// Copyright(C) 2024 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2024/7/7

package fst

import "testing"

func TestFileExists(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		FileExists(t, "os.go")
		FileNotExists(t, "os1.go")
		FileNotExists(t, "../fst/")
	})
	mt.Fail(func(t Testing) {
		FileExists(t, "os1.go")
		FileNotExists(t, "os.go")
	})
}

func TestDirExists(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		DirExists(t, "../fst/")
		DirNotExists(t, "demo")
	})
	mt.Fail(func(t Testing) {
		DirExists(t, "demo")
		DirNotExists(t, "../fst/")
	})
}
