// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rnd

import (
	"math/rand"
	"testing"
)

var __bench_nsamples int
var __bench_ints []int
var __bench_result int

func init() {
	rand.Seed(4321)
	Init(4321)
	__bench_nsamples = 1000
	__bench_ints = make([]int, __bench_nsamples)
	for i := 0; i < __bench_nsamples; i++ {
		__bench_ints[i] = i
	}
}

func Benchmark_go_int(b *testing.B) {
	var res int
	lo, hi := 0, 50
	for i := 0; i < b.N; i++ {
		for j := 0; j < __bench_nsamples; j++ {
			res = Int(lo, hi)
		}
	}
	__bench_result = res
}
