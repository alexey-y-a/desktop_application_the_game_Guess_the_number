package game

import (
	"math/rand/v2"
	"strconv"
)

const (
	MaxGuessesAmount = 10
	MaxGuess         = 300
	MinGuess         = 0
)

type Game struct {
	ComputerGuess int
	GuessesLeft   int
}

func NewGame() *Game {
	return &Game{
		ComputerGuess: rand.IntN(MaxGuess + 1),
		GuessesLeft:   MaxGuessesAmount,
	}
}

func (g *Game) StartNewGame() {
	g.ComputerGuess = rand.IntN(MaxGuess + 1)
	g.GuessesLeft = MaxGuessesAmount
}

func (g *Game) TryGuess(userGuess int) (string, string, bool) {
	if g.GuessesLeft == -1 {
		return "Для начала начните игру!", "Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!", false
	}

	if userGuess < MinGuess || userGuess > MaxGuess {
		return "Неправильное число! Введите число между 0 и 300.", "", false
	}

	if g.GuessesLeft == 1 && userGuess != g.ComputerGuess {
		g.GuessesLeft--
		return "К сожалению, вы не смогли отгадать число " + strconv.Itoa(g.ComputerGuess) +
			". Сыграйте еще раз!", "Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!", true
	}

	if userGuess == g.ComputerGuess {
		attemptsUsed := MaxGuessesAmount - g.GuessesLeft + 1
		g.GuessesLeft = -1
		return "Поздравляем! Вы отгадали число " + strconv.Itoa(g.ComputerGuess) +
			" за " + strconv.Itoa(attemptsUsed) +
			" попыток!", "Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!", true
	}

	g.GuessesLeft--
	guessHint := ""
	if userGuess > g.ComputerGuess {
		guessHint = "Ваше число больше загаданного!"
	} else {
		guessHint = "Ваше число меньше загаданного!"
	}
	return guessHint, "Попыток осталось: " + strconv.Itoa(g.GuessesLeft), false

}
