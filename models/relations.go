package models

type RecipeMaltRelation struct {
	ID       int    `json:"id"`
	RecipeID string `json:"title"`
	EBC      string `json:"EBC"`
	Amount   int    `json:"amount"`
}

func GetRecipeMaltRelation() RecipeMaltRelation {
	var recipe_malt_relation RecipeMaltRelation
	return recipe_malt_relation
}

func GetRecipeMaltRelations() []RecipeMaltRelation {
	var recipe_hop_relations []RecipeMaltRelation
	return recipe_hop_relations
}
