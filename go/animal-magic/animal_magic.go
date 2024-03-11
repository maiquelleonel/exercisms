package chance

import (
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(19) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return 0 + rand.Float64()*12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	for i := range animals {
		r := rand.Intn(len(animals))
		animals[i], animals[r] = animals[r], animals[i]
	}
	return animals
}
