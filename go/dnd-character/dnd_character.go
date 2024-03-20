package dndcharacter

import (
	"math"
	"math/rand"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	var dice []int
	dice = append(dice, d6(), d6(), d6(), d6())
	discard := slices.Index(dice, slices.Min(dice))
	dice = append(dice[:discard], dice[discard+1:]...)
	return sum(dice)
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	var c Character
	c.Strength = Ability()
	c.Dexterity = Ability()
	c.Constitution = Ability()
	c.Intelligence = Ability()
	c.Wisdom = Ability()
	c.Charisma = Ability()
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}

func d6() int {
	return rand.Intn(6) + 1
}

func sum(s []int) int {
	var r int
	for _, v := range s {
		r += v
	}
	return r
}
