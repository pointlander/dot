// Copyright 2020 The Dot Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func BenchmarkVMDot16(t *testing.B) {
	a, b := getVector16(), getVector16()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		vmdot16(a, b)
	}
}

func BenchmarkMDot16(t *testing.B) {
	a, b := getVector16(), getVector16()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		mdot16(a, b)
	}
}

func BenchmarkDot16(t *testing.B) {
	a, b := getVector16(), getVector16()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		dot16(a, b)
	}
}

func BenchmarkVMDot32(t *testing.B) {
	a, b := getVector32(), getVector32()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		vmdot32(a, b)
	}
}

func BenchmarkVDot32(t *testing.B) {
	a, b := getVector32(), getVector32()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		dot32(a, b)
	}
}

func BenchmarkDot32(t *testing.B) {
	a, b := getVector32(), getVector32()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		_dot32(a, b)
	}
}
