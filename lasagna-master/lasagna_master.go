package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, apt int) int {
	if apt == 0 {
		apt = 2
	}
	return len(layers) * apt
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodlesQuatnity := 0
	sauceQuatnity := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesQuatnity += 50
		}
		if layer == "sauce" {
			sauceQuatnity += 0.2
		}
	}
	return noodlesQuatnity, sauceQuatnity
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(listB []string, listA []string) {
	listA[len(listA)-1] = listB[len(listB)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, quantity int) []float64 {
	var scaledRecipe []float64

	for _, v := range recipe {
		scaledRecipe = append(scaledRecipe, v*(float64(quantity)/2))
	}
	return scaledRecipe

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
