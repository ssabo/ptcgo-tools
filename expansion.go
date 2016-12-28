package main

type expansion struct {
	name     string
	series   string
	id       string
	count    int
	standard bool
	expanded bool
	legacy   bool
}

func (e expansion) Name() string {
	return e.name
}

func (e expansion) Series() string {
	return e.series
}

func (e expansion) Id() string {
	return e.id
}

func (e expansion) Count() int {
	return e.count
}

func (e expansion) isStandard() bool {
	return e.standard
}

func (e expansion) isExpanded() bool {
	return e.expanded
}

func (e expansion) isLegacy() bool {
	return e.expanded
}

func Expansions() []expansion {
	expansions := []expansion{
		// expansion{"XY", "xy-series", "xy1", 146, false, true, true},
		// expansion{"XY-Flashfire", "xy-series", "xy2", 109, false, true, true},
		// expansion{"XY-Furious Fists", "xy-series", "xy3", 113, false, true, true},
		// expansion{"XY-Phantom Forces", "xy-series", "xy4", 119, false, true, true},
		// expansion{"XY-Primal Clash", "xy-series", "xy5", 160, true, true, true},
		// expansion{"Double Crisis", "xy-series", "dc1", 34, true, true, true},
		// expansion{"XY-Roaring Skies", "xy-series", "xy6", 110, true, true, true},
		// expansion{"XY-Ancient Origins", "xy-series", "xy7", 100, true, true, true},
		// expansion{"XY-BREAKthrough", "xy-series", "xy8", 164. true, true, true},
		// expansion{"XY-BREAKpoint", "xy-series", "xy9", 123, true, true, true},
		expansion{"Generations", "xy-series", "g1", 83, true, true, true},
		// expansion{"Fates Collide", "xy-series", "xy10", 120, true, true, true},
		// expansion{"Steam Siege", "xy-series", "xy11", 114, true, true, true},
		// expansion{"Evolutions", "xy-series", "xy12", 109, true, true, true},
	}
	return expansions
}
