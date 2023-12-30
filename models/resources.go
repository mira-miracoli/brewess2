package models

type Hop struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" binding:"required"`
	Iso    string `json:"iso" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

type Malt struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	EBC    string `json:"EBC"`
	Amount int    `json:"amount"`
}

type Yeast struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	MinTemp float64 `json:"minTemp"`
	MaxTemp float64 `json:"maxTemp"`
	Top   string  `json:"top"`
}
