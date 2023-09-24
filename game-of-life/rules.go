package gameoflife

type Rules struct {
	LiveToDie  Neighbours
	DeadToLive Neighbours
}

type Neighbours struct {
	Zero  bool
	One   bool
	Two   bool
	Three bool
	Four  bool
	Five  bool
	Six   bool
	Seven bool
	Eight bool
}
