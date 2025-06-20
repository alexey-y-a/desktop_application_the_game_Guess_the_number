package game

import "testing"

func TestTryGuess(t *testing.T) {
	g := &Game{ComputerGuess: 100, GuessesLeft: MaxGuessesAmount}

	result, status, gameEnded := g.TryGuess(150)
	if result != "Ваше число больше загаданного!" {
		t.Errorf("Expected hint 'Ваше число больше загаданного!', got '%s'", result)
	}
	if status != "Попыток осталось: 9" {
		t.Errorf("Expected status 'Попыток осталось: 9', got '%s'", status)
	}
	if gameEnded {
		t.Errorf("Expected gameEnded to be false, got %v", gameEnded)
	}
	if g.GuessesLeft != 9 {
		t.Errorf("Expected GuessesLeft to be 9, got %d", g.GuessesLeft)
	}

	result, status, gameEnded = g.TryGuess(50)
	if result != "Ваше число меньше загаданного!" {
		t.Errorf("Expected hint 'Ваше число меньше загаданного!', got '%s'", result)
	}
	if status != "Попыток осталось: 8" {
		t.Errorf("Expected status 'Попыток осталось: 8', got '%s'", status)
	}
	if gameEnded {
		t.Errorf("Expected gameEnded to be false, got %v", gameEnded)
	}
	if g.GuessesLeft != 8 {
		t.Errorf("Expected GuessesLeft to be 8, got %d", g.GuessesLeft)
	}
}

func TestTryGuessCorrect(t *testing.T) {
	g := &Game{ComputerGuess: 100, GuessesLeft: MaxGuessesAmount}

	result, status, gameEnded := g.TryGuess(100)
	expectedResult := "Поздравляем! Вы отгадали число 100 за 1 попыток!"
	if result != expectedResult {
		t.Errorf("Expected result '%s', got '%s'", expectedResult, result)
	}
	if status != "Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!" {
		t.Errorf("Expected status 'Нажмите кнопку ниже...', got '%s'", status)
	}
	if !gameEnded {
		t.Errorf("Expected gameEnded to be true, got %v", gameEnded)
	}
	if g.GuessesLeft != -1 {
		t.Errorf("Expected GuessesLeft to be -1, got %d", g.GuessesLeft)
	}
}

func TestTryGuessGameOver(t *testing.T) {
	g := &Game{ComputerGuess: 100, GuessesLeft: 1}

	result, status, gameEnded := g.TryGuess(50)
	expectedResult := "К сожалению, вы не смогли отгадать число 100. Сыграйте еще раз!"
	if result != expectedResult {
		t.Errorf("Expected result '%s', got '%s'", expectedResult, result)
	}
	if status != "Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!" {
		t.Errorf("Expected status 'Нажмите кнопку ниже...', got '%s'", status)
	}
	if !gameEnded {
		t.Errorf("Expected gameEnded to be true, got %v", gameEnded)
	}
	if g.GuessesLeft != 0 {
		t.Errorf("Expected GuessesLeft to be 0, got %d", g.GuessesLeft)
	}
}

func TestTryGuessGameNotStarted(t *testing.T) {
	g := &Game{GuessesLeft: -1}

	result, status, gameEnded := g.TryGuess(50)
	if result != "Для начала начните игру!" {
		t.Errorf("Expected result 'Для начала начните игру!', got '%s'", result)
	}
	if status != "Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!" {
		t.Errorf("Expected status 'Нажмите кнопку ниже...', got '%s'", status)
	}
	if gameEnded {
		t.Errorf("Expected gameEnded to be false, got %v", gameEnded)
	}
	if g.GuessesLeft != -1 {
		t.Errorf("Expected GuessesLeft to remain -1, got %d", g.GuessesLeft)
	}
}

func TestTryGuessInvalidInput(t *testing.T) {
	g := &Game{ComputerGuess: 100, GuessesLeft: 5}

	result, status, gameEnded := g.TryGuess(301)
	if result != "Неправильное число! Введите число между 0 и 300." {
		t.Errorf("Expected result 'Неправильное число...', got '%s'", result)
	}
	if status != "" {
		t.Errorf("Expected empty status, got '%s'", status)
	}
	if gameEnded {
		t.Errorf("Expected gameEnded to be false, got %v", gameEnded)
	}
	if g.GuessesLeft != 5 {
		t.Errorf("Expected GuessesLeft to remain 5, got %d", g.GuessesLeft)
	}
}

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g.GuessesLeft != MaxGuessesAmount {
		t.Errorf("Expected GuessesLeft to be %d, got %d", MaxGuessesAmount, g.GuessesLeft)
	}
	if g.ComputerGuess < MinGuess || g.ComputerGuess > MaxGuess {
		t.Errorf("Expected ComputerGuess to be in range [%d, %d], got %d", MinGuess, MaxGuess, g.ComputerGuess)
	}
}

func TestStartNewGame(t *testing.T) {
	g := &Game{ComputerGuess: 100, GuessesLeft: 2}
	g.StartNewGame()
	if g.GuessesLeft != MaxGuessesAmount {
		t.Errorf("Expected GuessesLeft to be %d, got %d", MaxGuessesAmount, g.GuessesLeft)
	}
	if g.ComputerGuess < MinGuess || g.ComputerGuess > MaxGuess {
		t.Errorf("Expected ComputerGuess to be in range [%d, %d], got %d", MinGuess, MaxGuess, g.ComputerGuess)
	}
}
