package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	(myList)[len(myList)-1] = (friendsList)[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))
	copy(result, quantities)

	for i := range result {
		result[i] = result[i] * (float64(portions) / 2.0)
	}
	return result
}