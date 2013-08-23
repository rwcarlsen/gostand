package lifetime

import (
	"math"
	"testing"

	"github.com/rwcarlsen/gostand/stand"
)

var res1 = &stand.Resource{"bananas", 563}
var res2 = &stand.Resource{"monkeys", 1}

type Fac struct {
	ReqFreq, OffFreq int
	ReqStop, OffStop int
}

func (f *Fac) Tick(s stand.Stand) {
	t := s.Time()
	if t < f.ReqStop && t%f.ReqFreq == 0 {
		s.Request(res1)
	}
	if t < f.OffStop && t%f.OffFreq == 0 {
		s.Offer(res1)
	}
}

func TestEternal(t *testing.T) {
	f := &Fac{1, 1, math.MaxInt32, math.MaxInt32}

	s := new(Stand)

	s.SetFac(f)
	results := s.Run().(*Result)

	if death := results.DeathTime; death != -1 {
		t.Errorf("Value of DeathTime - expected %v, got %v", -1, death)
	}
}

func TestNever(t *testing.T) {
	f := &Fac{1, 1, 0, 0}

	s := new(Stand)

	s.SetFac(f)
	results := s.Run().(*Result)

	if death := results.DeathTime; death != 0 {
		t.Errorf("Value of DeathTime - expected %v, got %v", 0, death)
	}
}
