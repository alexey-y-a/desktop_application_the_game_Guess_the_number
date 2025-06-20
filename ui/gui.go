package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"game/internal/game"
	"strconv"
)

func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Угадай число!")

	g := &game.Game{GuessesLeft: -1}

	computerGuessLabel := widget.NewLabel("Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!")

	startGameButton := widget.NewButton("Загадать новое число!", func() {
		g.StartNewGame()
		computerGuessLabel.SetText("Компьютер загадал число! Попыток осталось: " + strconv.Itoa(g.GuessesLeft))
	})

	guessDisplay := widget.NewLabel("Введите ваше первое число!")

	userGuessInput := widget.NewEntry()
	userGuessInput.SetPlaceHolder("Введите ваше число: ")

	tryGuessButton := widget.NewButton("Попробовать", func() {
		userGuess, err := strconv.Atoi(userGuessInput.Text)
		if err != nil {
			guessDisplay.SetText("Пожалуйста, введите число!")
			return
		}

		guessResult, statusText, gameEnded := g.TryGuess(userGuess)
		guessDisplay.SetText(guessResult)
		if statusText != "" {
			computerGuessLabel.SetText(statusText)
		}
		if gameEnded {
			g.GuessesLeft = -1
		}
		userGuessInput.SetText("")
	})

	userGuessInput.OnSubmitted = func(text string) {
		tryGuessButton.OnTapped()
	}

	myWindow.SetContent(
		container.NewCenter(
			container.NewVBox(
				computerGuessLabel,
				startGameButton,
				guessDisplay,
				userGuessInput,
				tryGuessButton,
			),
		),
	)

	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()
}
