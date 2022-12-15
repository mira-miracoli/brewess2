package models

type Yeast struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Top   string `json:"top"`
}

func GetYeast() Yeast {
	var yeast Yeast
	return yeast
}

func GetYeasts() []Yeast {
	var yeasts []Yeast
	return yeasts
}
