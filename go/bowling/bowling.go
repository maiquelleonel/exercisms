package bowling

import (
	"errors"
)

// Define the Game type here.
type Game struct {
	rolls []int
}

var ErrNegativeRoll = errors.New("negative roll is invalid")
var ErrExceededPins = errors.New("pin count exceeds pins on the lane")
var ErrGameOver = errors.New("cannot roll after game is over")
var ErrCannotScore = errors.New("cannot score yet")

func NewGame() *Game {
	return &Game{
		rolls: []int{},
	}
}

const (
	maxFrames = 20
	bonusRoll = 21
	endGame   = 22
)

func (g *Game) Roll(pins int) error {
	if pins < 0 {
		return ErrNegativeRoll
	}

	if pins > 10 {
		return ErrExceededPins
	}

	switch len(g.rolls) {

	case maxFrames:
		if g.rolls[18]+g.rolls[19] != 10 {
			return ErrGameOver
		}

	case bonusRoll:
		if g.rolls[18] != 10 {
			return ErrGameOver
		}
		if g.rolls[20] != 10 && g.rolls[20]+pins > 10 {
			return ErrExceededPins
		}
	case endGame:
		return ErrGameOver
	}

	lastRoll := len(g.rolls) - 1

	if len(g.rolls) < maxFrames &&
		len(g.rolls)%2 == 1 &&
		pins+g.rolls[lastRoll] > 10 {
		return ErrExceededPins
	}

	if pins == 10 {
		g.rolls = append(g.rolls, pins)
		if len(g.rolls) < maxFrames {
			g.rolls = append(g.rolls, 0)
		}

	} else {
		g.rolls = append(g.rolls, pins)
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if len(g.rolls) < maxFrames ||
		g.rolls[18] == 10 && len(g.rolls) != endGame ||
		g.rolls[18] != 10 && g.rolls[18]+g.rolls[19] == 10 && len(g.rolls) != bonusRoll {
		return 0, ErrCannotScore
	}

	s := 0
	for i := 0; i < maxFrames; i += 2 {
		if g.rolls[i] == 10 {
			s += 10
			if g.rolls[i+2] == 10 {
				if i == 18 {
					s += g.rolls[i+2] + g.rolls[i+3]
				} else {
					s += 10 + g.rolls[i+4]
				}
			} else {
				s += g.rolls[i+2] + g.rolls[i+3]
			}
		} else if g.rolls[i]+g.rolls[i+1] == 10 {
			s += 10 + g.rolls[i+2]
		} else {
			s += g.rolls[i] + g.rolls[i+1]
		}
	}

	return s, nil
}
