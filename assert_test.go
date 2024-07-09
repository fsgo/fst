// Copyright(C) 2023 github.com/fsgo  All Rights Reserved.
// Author: hidu <duv123@gmail.com>
// Date: 2023/9/8

package fst

import (
	"fmt"
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

func TestSamePtr(t *testing.T) {
	type TStruct struct {
		_ int
	}
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		v1 := &TStruct{}
		v2 := v1
		SamePtr(t, v1, v2)
		v3 := &TStruct{}
		NotSamePtr(t, v1, v3)
	})
	mt.Fail(func(t Testing) {
		SamePtr(t, &TStruct{}, &TStruct{})
		v1 := &TStruct{}
		v2 := v1
		NotSamePtr(t, v1, v2)
	})
}

func TestLess(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		Less(t, 1, 2)
		Less(t, "a", "b")
		Less(t, 0.1, 0.2)
		Less(t, uint32(1), uint32(2))
	})
	mt.Fail(func(t Testing) {
		Less(t, 3, 2)
		Less(t, "c", "b")
		Less(t, 0.3, 0.2)
		Less(t, uint32(3), uint32(2))
	})
}

func TestLessOrEqual(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		LessOrEqual(t, 1, 2)
		LessOrEqual(t, 2, 2)

		LessOrEqual(t, "a", "b")
		LessOrEqual(t, "b", "b")

		LessOrEqual(t, 0.1, 0.2)
		LessOrEqual(t, 0.2, 0.2)

		LessOrEqual(t, uint32(1), uint32(2))
		LessOrEqual(t, uint32(2), uint32(2))
	})
	mt.Fail(func(t Testing) {
		LessOrEqual(t, 3, 2)
		LessOrEqual(t, "c", "b")
		LessOrEqual(t, 0.3, 0.2)
		LessOrEqual(t, uint32(3), uint32(2))
	})
}

func TestGreater(t *testing.T) {
	type intA int
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		Greater(t, 3, 2)
		Greater(t, intA(3), intA(2))
		Greater(t, "c", "b")
		Greater(t, 0.3, 0.2)
		Greater(t, uint32(3), uint32(2))
	})
	mt.Fail(func(t Testing) {
		Greater(t, 1, 2)
		Greater(t, 2, 2)

		Greater(t, "a", "b")
		Greater(t, "b", "b")

		Greater(t, 0.2, 0.2)
		Greater(t, 0.1, 0.2)

		Greater(t, uint32(2), uint32(2))
		Greater(t, uint32(1), uint32(2))
	})
}

func TestGreaterOrEqual(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		GreaterOrEqual(t, 3, 2)
		GreaterOrEqual(t, 3, 3)

		GreaterOrEqual(t, "c", "b")
		GreaterOrEqual(t, "c", "c")

		GreaterOrEqual(t, 0.3, 0.2)
		GreaterOrEqual(t, 0.3, 0.3)

		GreaterOrEqual(t, uint32(3), uint32(2))
		GreaterOrEqual(t, uint32(3), uint32(3))
	})
	mt.Fail(func(t Testing) {
		GreaterOrEqual(t, 1, 2)
		GreaterOrEqual(t, "a", "b")
		GreaterOrEqual(t, 0.1, 0.2)
		GreaterOrEqual(t, uint32(1), uint32(2))
	})
}

func TestErrorIs(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		ErrorIs(t, io.EOF, io.EOF)
		ErrorIs(t, fmt.Errorf("%w ,ok", io.EOF), io.EOF)
	})
	mt.Fail(func(t Testing) {
		ErrorIs(t, nil, io.EOF)
		ErrorIs(t, io.EOF, fmt.Errorf("%w ,ok", io.EOF))
	})
}

func TestNotErrorIs(t *testing.T) {
	mt := newMyTesting(t)
	mt.Success(func(t Testing) {
		NotErrorIs(t, nil, io.EOF)
		NotErrorIs(t, io.EOF, fmt.Errorf("%w ,ok", io.EOF))
	})
	mt.Fail(func(t Testing) {
		NotErrorIs(t, io.EOF, io.EOF)
	})
}
