package models

type Hop struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Iso   string `json:"iso"`
}

func GetHop() Hop {
	var hop Hop
	return hop
}

func GetHops() []Hop {
	var hops []Hop
	return hops
}
