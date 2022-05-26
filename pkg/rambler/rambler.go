// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package rambler

import (
	"arod-im/pkg/murmur3"
	"time"
)

const (
	SequenceBit = 12
	//SessionTypeBit = 5
	MaxSeq = 0xFFF
)

var m = [32]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A',
	'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
	'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
}

type Rambler struct {
	ret int64
}

func NewRambler() *Rambler {
	return &Rambler{}
}

func (r *Rambler) GetSeqID(sessionID []byte) string {
	return r.generate(sessionID)
}

func (r *Rambler) generate(sessionID []byte) string {
	highBits := time.Now().UnixMilli()

	seq := r.getMessageSeq()
	highBits = highBits << SequenceBit
	highBits = highBits | seq

	//highBits = highBits << SessionTypeBit
	//highBits = highBits | int64(sessionType&0xF)

	id := int64(murmur3.Sum64(sessionID)) & 0x1FFFFFF

	res := make([]byte, 16)
	for i := 15; i >= 11; i-- {
		res[i] = m[id&0x1F]
		id >>= 5
	}
	for i := 10; i >= 0; i-- {
		res[i] = m[highBits&0x1F]
		highBits >>= 5
	}

	return string(res)
}
func (r *Rambler) getMessageSeq() int64 {
	r.ret++
	r.ret %= MaxSeq
	return r.ret
}
