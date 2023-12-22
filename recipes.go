package main

type Recipe struct {
	Name        string
	Ingredients map[string]int
}

var recipes = []Recipe{
	{
		Name:        "Tomato Soup",
		Ingredients: map[string]int{"Crop": 2},
	},
}

func (inv *Inventory) HasIngredients(recipe Recipe) bool {
	requiredCrops, exists := recipe.Ingredients["Crop"]
	if !exists {
		return false
	}

	return inv.Crops >= requiredCrops
}
func (inv *Inventory) CraftRecipe(recipe Recipe) bool {
	requiredCrops, exists := recipe.Ingredients["Crop"]
	if !exists || inv.Crops < requiredCrops {
		return false
	}

	inv.Crops -= requiredCrops

	return true

}
func (inv *Inventory) RemoveItems(ingredients map[string]int) {
	for item, qty := range ingredients {
		inv.Items[item] -= qty
	}
}
