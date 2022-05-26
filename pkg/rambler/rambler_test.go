// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package rambler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkGetID(b *testing.B) {
	r := NewRambler()
	for i := 0; i < b.N; i++ {
		r.GetSeqID([]byte("100001"))
	}
}

func TestGetSeqID(t *testing.T) {
	r := NewRambler()
	m := r.GetSeqID([]byte("100001"))
	n := r.GetSeqID([]byte("100001"))
	assert.Less(t, m, n)
}
