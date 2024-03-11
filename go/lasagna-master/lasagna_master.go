package lasagna

func PreparationTime(layers []string, averageTime int) int {
	if averageTime <= 0 {
		averageTime = 2
	}
	return len(layers) * averageTime
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQuantities []float64
	for _, v := range quantities {
		newQuantities = append(newQuantities, ((v / 2) * float64(portions)))
	}
	return newQuantities
}
