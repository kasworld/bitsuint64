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

import "testing"

func TestSetBit(t *testing.T) {
	var b Bits64
	t.Logf("%v", b)

	(&b).SetBitAt(5)
	t.Logf("SetBitAt(5) %v", b)
	(&b).SetBitsByUint64(0x0f0f0f)
	t.Logf("SetBitsByUint64(0x0f0f0f) %v", b)

	(&b).ClearBitAt(2)
	t.Logf("ClearBitAt(2) %v", b)
	(&b).ClearBitsByUint64(0x0f)
	t.Logf("ClearBitsByUint64(0x0f) %v", b)

	b.SetBitsByUint64(0)
	t.Logf("%v", b)
	(&b).NegBitAt(5)
	t.Logf("NegBitAt(5) %v", b)
	(&b).NegBitAt(5)
	t.Logf("NegBitAt(5) %v", b)
	(&b).NegBitsByUint64(0xf0f0f0f)
	t.Logf("NegBitsByUint64(0xf0f0f0f) %v", b)

	b.SetBitsByUint64(0)
	t.Logf("TestBitAt(7) %v", (&b).TestBitAt(7))
	t.Logf("TestBitsByUint64(0x704) %v", (&b).TestBitsByUint64(0x704))
}
