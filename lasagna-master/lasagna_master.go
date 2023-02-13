package lasagna

const defaultPreparationTime int = 2

// Estimates how long it takes to cook a lasagna based on layers and time for each layer.
func PreparationTime(layers []string, layerPreparationTime int) int {
	layersNumber := len(layers)

	if layerPreparationTime == 0 {
		layerPreparationTime = defaultPreparationTime
	}

	return layersNumber * layerPreparationTime
}

const noodlePerLayer int = 50
const saucePerLayer float64 = 0.2

// Countes how much noodles and sauce needed based on the number of layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	var noodleLayers int
	var sauceLayers int

	for _, v := range layers {
		switch v {
		case "noodles":
			noodleLayers += 1
		case "sauce":
			sauceLayers += 1
		default:
			continue
		}
	}

	noodles = noodleLayers * noodlePerLayer
	sauce = float64(sauceLayers) * saucePerLayer

	return
}

// Adds friend's secret ingredient to ownIngredients slice as the last element.
func AddSecretIngredient(friendIngredients, ownIngredients []string) {
	secretIngredient := friendIngredients[len(friendIngredients)-1]
	ownIngredients[len(ownIngredients)-1] = secretIngredient
}

// Counts how much of each ingredient grams is necessary for the provided number of portions.
func ScaleRecipe(quantities []float64, portions int) (currentQuantities []float64) {
	// Create a copy of quantities array.
	currentQuantities = make([]float64, len(quantities))
	copy(currentQuantities, quantities)

	// Update quantites based on the number of portions.
	for i := range currentQuantities {
		currentQuantities[i] *= float64(portions) / 2
	}

	return currentQuantities
}
