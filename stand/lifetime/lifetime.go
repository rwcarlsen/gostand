package lifetime

import (
	"github.com/rwcarlsen/gostand/stand"
)

const (
	month    = 1
	year     = 12 * month
	maxdur   = 1000 * year
	deathdur = 50 * year
)

type Stand struct {
	offers, requests   bool
	lastOffer, lastReq int
	time               int
	fac                stand.Facility
}

func (s *Stand) Time() int {
	return s.time
}

func (s *Stand) Request(r *stand.Resource) *stand.Resource {
	s.requests = true
	s.lastOffer = s.time
	return r
}

func (s *Stand) Offer(r *stand.Resource) bool {
	s.offers = true
	s.lastReq = s.time
	return true
}

func (s *Stand) Run() stand.Result {
	s.offers, s.requests = false, false
	s.lastOffer, s.lastReq = -1, -1

	deathTime := -1
	for s.time = 0; s.time < maxdur; s.time++ {
		s.fac.Tick(s)

		deadOff := s.offers && (s.time-s.lastOffer >= deathdur) || !s.offers
		deadReq := s.requests && (s.time-s.lastReq >= deathdur) || !s.requests
		deadNothing := !s.requests && !s.offers && s.time >= deathdur
		if deadNothing {
			deathTime = 0
			break
		} else if deadOff && deadReq && s.time > deathdur {
			deathTime = max(s.lastOffer, s.lastReq)
			break
		}
	}

	return &Result{deathTime, s.lastOffer, s.lastReq}
}

func (s *Stand) SetFac(f stand.Facility) {
	s.fac = f
}

type Result struct {
	DeathTime   int
	OfferStop   int
	RequestStop int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
