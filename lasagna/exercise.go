package lasagna

const NOODLES_G = 50
const SAUCE_L = 0.2

func PreparationTime(layers []string, avgPrep int) int {
	if avgPrep == 0 {
		return len(layers) * 2
	}

	return len(layers) * avgPrep
}

func Quantities(layers []string) (int, float64) {
	noodlesCount := 0
	sauceCount := 0

	for i := 0; i < len(layers); i++ {
		layer := layers[i]

		if layer == "noodles" {
			noodlesCount++
		} else if layer == "sauce" {
			sauceCount++
		}
	}

	return noodlesCount * NOODLES_G, float64(sauceCount) * SAUCE_L
}

func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]

	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))

	multiplier := float64(portions) / 2

	for i := 0; i < len(quantities); i++ {
		scaled[i] = quantities[i] * multiplier
	}

	return scaled
}
