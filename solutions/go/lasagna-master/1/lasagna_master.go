package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTimePerLayer int) int {
    factor := avgPrepTimePerLayer
    if avgPrepTimePerLayer <= 0 {
        factor = 2
    }
    return len(layers) * factor
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    const noodlesPerLayer = 50
    const litersOfSaucePerLayer = 0.2

    noodles := 0
    sauce := 0
    
    for _, value := range layers {
        if value == "noodles" {
            noodles++
        } else if value == "sauce" {
            sauce++
        }
    }

    return (noodles * noodlesPerLayer), (float64(sauce) * litersOfSaucePerLayer)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
func ScaleRecipe(quantities []float64, portions int) []float64 {
    scaled := make([]float64, len(quantities))
    copy(scaled, quantities)
    for i := 0; i < len(scaled); i++ {
        scaled[i] = (quantities[i] / 2.0) * float64(portions)
    }
    return scaled
}
