// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bitsuint64

import (
	"fmt"
	"sync/atomic"
)

type Bits64 uint64

func (bt Bits64) String() string {
	return fmt.Sprintf("0b%0b", bt)
}

func (bt *Bits64) SetBitAt(n uint8) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval | (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) ClearBitAt(n uint8) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval &^ (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) NegBitAt(n uint8) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval ^ (1 << n)
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) TestBitAt(n uint8) bool {
	val := atomic.LoadUint64((*uint64)(bt)) & (1 << n)
	return val != 0
}

func (bt *Bits64) SetBitsByUint64(v uint64) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval | v
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) ClearBitsByUint64(v uint64) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval &^ v
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) NegBitsByUint64(v uint64) {
	var success bool
	for !success {
		oldval := atomic.LoadUint64((*uint64)(bt))
		newval := oldval ^ v
		success = atomic.CompareAndSwapUint64((*uint64)(bt), oldval, newval)
	}
}

func (bt *Bits64) TestBitsByUint64(v uint64) bool {
	val := atomic.LoadUint64((*uint64)(bt)) & v
	return val != 0
}

func (bt *Bits64) SetByUint64(v uint64) {
	atomic.StoreUint64((*uint64)(bt), v)
}
func (bt *Bits64) GetUint64() uint64 {
	return atomic.LoadUint64((*uint64)(bt))
}

func (bt *Bits64) Countbits64() int {
	v := bt.GetUint64()
	v -= (v >> 1) & 0x5555555555555555
	v = (v & 0x3333333333333333) + ((v >> 2) & 0x3333333333333333)
	v = (v + (v >> 4)) & 0x0f0f0f0f0f0f0f0f
	return int((v * 0x0101010101010101) >> 56)
}
