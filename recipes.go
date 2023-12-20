package main

import "fmt"

type Recipe struct {
	Name        string
	Ingredients map[string]int
	Result      string
}

var recipes = []Recipe{
	{
		Name:        "Tomato Soup",
		Ingredients: map[string]int{"Tomato": 2, "Greens": 1},
		Result:      "CraftedItem1",
	},
}

func (inv *Inventory) HasIngredients(ingredients map[string]int) bool {
	for item, qty := range ingredients {
		if inv.Items[item] < qty {
			return false
		}
	}
	return true
}

func (inv *Inventory) RemoveItems(ingredients map[string]int) {
	for item, qty := range ingredients {
		inv.Items[item] -= qty
	}
}

// got help from chatgpt with this function
func (m *mapGame) CraftItem(recipe Recipe) {
	if m.playerInventory.HasIngredients(recipe.Ingredients) {
		m.playerInventory.RemoveItems(recipe.Ingredients)
		m.playerInventory.Items[recipe.Result]++
		fmt.Println("Crafted:", recipe.Result)
	} else {
		fmt.Println("Missing ingredients for:", recipe.Name)
	}
}
func (m *mapGame) getSelectedRecipe() Recipe {
	if m.selectedRecipeIndex < 0 || m.selectedRecipeIndex >= len(recipes) {
		return Recipe{} // Return empty recipe if out of bounds
	}
	return recipes[m.selectedRecipeIndex]
}
