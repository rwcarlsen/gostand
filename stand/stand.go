package stand

import "reflect"

var stands []Stand
var facs []Facility

func RegisterStand(s Stand) {
	stands = append(stands, s)
}

func RegisterFacility(f Facility) {
	facs = append(facs, f)
}

func RunAll() map[string]map[string]Result {
	results := make(map[string]map[string]Result)
	for _, s := range stands {
		standresults := make(map[string]Result)
		for _, f := range facs {
			s.SetFac(f)
			fname := reflect.TypeOf(f).Name()
			standresults[fname] = s.Run()
		}
		sname := reflect.TypeOf(s).Name()
		results[sname] = standresults
	}
	return results
}

type Facility interface {
	Tick(Stand)
}

type Result interface{}

type Stand interface {
	// Request asks the stand for a resource like r. The returned resource
	// is owned by the caller and may or may not be like r.
	Request(r *Resource) *Resource
	// Offer returns true if r is accepted for trade. The offerer no longer owns r.
	Offer(r *Resource) bool
	// Time returns the current simulation time
	Time() int
	// Run executes all tests by the stand on the specified facility. The
	// returned result must be marshallable to JSON.
	Run() Result
	// SetFac makes f the facility being tested by this stand.
	SetFac(f Facility)
}

type Resource struct {
	Type string
	Qty  float64
}
