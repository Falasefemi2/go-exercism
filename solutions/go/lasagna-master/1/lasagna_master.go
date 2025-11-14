package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageLayer int) int {
    if averageLayer == 0 {
        return 4
    }
    return len(layers) * averageLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	n := 0
	s := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			n++
		}
		if layer == "sauce" {
			s++
		}
	}
	return n * 50, s * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(a, b []string) {
	secret := a[len(a)-1]
	b[len(b)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newSlice := make([]float64, len(quantities))
	for i, q := range quantities {
		newSlice[i] = q * float64(portions) / 2
	}
	return newSlice
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
