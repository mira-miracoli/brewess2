package models

type Malt struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	EBC   string `json:"EBC"`
}

func GetMalt() Malt {
	var malt Malt
	return malt
}

func GetMalts() []Malt {
	var malts []Malt
	return malts
}
