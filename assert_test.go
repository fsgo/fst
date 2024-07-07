// Copyright(C) 2023 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2023/9/8

package fst

import (
	"io"
	"testing"
)

func TestEqual(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		Equal(t, 1, 1)
		Equal(t, "a", "a")
		Equal(t, []int{1}, []int{1})
		NotEqual(t, "a", "b")
		NotEqual(t, 1, 2)
	})

	mt.Fail(func(t Testing) {
		Equal(t, 1, 2)
		Equal(t, "a", "b")
		NotEqual(t, 1, 1)
	})
}

func TestError(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		Error(t, io.EOF)
		NoError(t, nil)
		var err1 error
		NoError(t, err1)
	})
	mt.Fail(func(t Testing) {
		NoError(t, io.EOF)
		Error(t, nil)
	})
}

func TestTrue(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		True(t, true)
		False(t, false)
	})
	mt.Fail(func(t Testing) {
		True(t, false)
		False(t, true)
	})
}

func TestNil(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		Nil(t, nil)
		NotNil(t, 1)
	})
	mt.Fail(func(t Testing) {
		Nil(t, 1)
		NotNil(t, nil)
	})
}

func TestEmpty(t *testing.T) {
	type TStruct struct {
		x int
	}
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		Empty(t, nil)
		Empty(t, 0)
		Empty(t, false)
		Empty(t, make(chan int))
		Empty(t, "")
		Empty(t, [1]int{})
		Empty(t, TStruct{})

		NotEmpty(t, 1)
		NotEmpty(t, true)
		NotEmpty(t, [1]int{1})
		NotEmpty(t, "ok")
		NotEmpty(t, TStruct{x: 1})

		var v1 *TStruct
		Empty(t, v1)

		Empty(t, &TStruct{})
	})
	mt.Fail(func(t Testing) {
		Empty(t, 1)
		Empty(t, true)

		NotEmpty(t, false)
		NotEmpty(t, 0)
	})
}

func TestStringContains(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		StringContains(t, "hello", "h")
		StringNotContains(t, "hello", "a")
	})
	mt.Fail(func(t Testing) {
		StringContains(t, "hello", "a")
		StringNotContains(t, "hello", "h")
	})
}

func TestSliceContains(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		SliceContains(t, []int{1, 2}, 1)
		SliceNotContains(t, []int{1, 2}, 3)
	})
	mt.Fail(func(t Testing) {
		SliceContains(t, []int{1, 2}, 3)
		SliceNotContains(t, []int{1, 2}, 1)
	})
}
