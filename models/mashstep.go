package models

type Mashstep struct {
	ID          int `json:"id"`
	Time        int `json:"time"`
	Temperature int `json:"temperature"`
}

func GetMashstep() Mashstep {
	var mashstep Mashstep
	return mashstep
}

func GetMashsteps() []Mashstep {
	var mashsteps []Mashstep
	return mashsteps
}
