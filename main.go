// Copyright 2020 The Dot Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"math/rand"
)

const Stride = 1024 * 1024

func getVector32() []float32 {
	vector := make([]float32, 1024*1024*256)
	for i := range vector {
		vector[i] = rand.Float32()
	}
	return vector
}

func getVector16() []uint16 {
	vector := make([]uint16, 1024*1024*256)
	for i := range vector {
		bits := math.Float32bits(rand.Float32())
		vector[i] = uint16(bits >> 16)
	}
	return vector
}

func convert(aa []float32, a []uint16) {
	for j, value := range a {
		aa[j] = math.Float32frombits(uint32(value) << 16)
	}
}

func vmdot16(a, b []uint16) float32 {
	var sum float32
	done := make(chan float32, 16)
	aaa, bbb := make(chan []float32, 16), make(chan []float32, 16)
	for i := 0; i < 16; i++ {
		aaa <- make([]float32, Stride)
		bbb <- make([]float32, Stride)
	}
	process := func(i int) {
		aa, bb := <-aaa, <-bbb
		end := i + Stride
		convert(aa, a[i:end])
		convert(bb, b[i:end])
		aaa <- aa
		bbb <- bb
		done <- dot32(aa, bb)
	}
	flight, i := 0, 0
	for i < len(a) && flight < 16 {
		go process(i)
		i += Stride
		flight++
	}
	for i < len(a) {
		sum += <-done
		flight--

		go process(i)
		i += Stride
		flight++
	}
	for i := 0; i < flight; i++ {
		sum += <-done
	}
	return sum
}

func mdot16(a, b []uint16) float32 {
	var sum float32
	done := make(chan float32, 16)
	process := func(i int) {
		end := i + Stride
		done <- dot16(a[i:end], b[i:end])
	}
	flight, i := 0, 0
	for i < len(a) && flight < 16 {
		go process(i)
		i += Stride
		flight++
	}
	for i < len(a) {
		sum += <-done
		flight--

		go process(i)
		i += Stride
		flight++
	}
	for i := 0; i < flight; i++ {
		sum += <-done
	}
	return sum
}

func dot16(a, b []uint16) float32 {
	var sum float32
	for i, value := range a {
		x := math.Float32frombits(uint32(value) << 16)
		y := math.Float32frombits(uint32(b[i]) << 16)
		sum += x * y
	}
	return sum
}

func vmdot32(a, b []float32) float32 {
	var sum float32
	done := make(chan float32, 16)
	process := func(i int) {
		end := i + Stride
		done <- dot32(a[i:end], b[i:end])
	}
	flight, i := 0, 0
	for i < len(a) && flight < 16 {
		go process(i)
		i += Stride
		flight++
	}
	for i < len(a) {
		sum += <-done
		flight--

		go process(i)
		i += Stride
		flight++
	}
	for i := 0; i < flight; i++ {
		sum += <-done
	}
	return sum
}

func _dot32(X, Y []float32) float32 {
	var sum float32
	for i, x := range X {
		sum += x * Y[i]
	}
	return sum
}

func main() {
	a, b := getVector16(), getVector16()
	fmt.Println(vmdot16(a, b))
}
