package main

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"
	"testing/quick"
)

func TestThing(t *testing.T) {
	got := Thing(true)
	if got != false {
		t.Errorf("Thing() = %v; want false", got)
	}
}

func TestEven(t *testing.T) {
	f := func(x int) bool {
		y := Even(x)
		return y%2 == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

func FuzzHex(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := hex.EncodeToString(in)
		out, err := hex.DecodeString(enc)
		if err != nil {
			t.Fatalf("%v: decode: %v", in, err)
		}
		if !bytes.Equal(in, out) {
			t.Fatalf("%v: not equal after round trip: %v", in, out)
		}
	})
}

func FuzzMaybePanic(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Fuzz(func(t *testing.T, i int) {
		actual := MaybePanic(i)
		if actual != i {
			t.Fatalf("MaybePanic: does not act as identity: %v", i)
		}
	})
}
