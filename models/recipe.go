package models

type Recipe struct {
	Id               int
	Title            string  `json:"title"`
	BasicInfo        string  `json:"basic_info"`
	HopInfo          string  `json:"hop_info"`
	MaltInfo         string  `json:"malt_info"`
	MashInfo         string  `json:"mash_info"`
	FermentationInfo string  `json:"fermentation_info"`
	IBU              float64 `json:"ibu"`
	EBC              float64 `json:"ebc"`
	OGTarget         float64 `json:"og_target"` // specifies the targeted original gravity in %sacc
	CastWorth        float64 `json:"cast_worth"`
	CookingTime      float64 `json:"cooking_time"`
	SHA              float64 `json:"sha"`
}

func GetRecipe() Recipe {
	var recipe Recipe
	return recipe
}

func GetRecipes() []Recipe {
	var recipes []Recipe
	return recipes
}
