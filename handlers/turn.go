package handlers

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func GetPlayerName() string {
	playerName := PlayerInput()
	fmt.Println("Your player name is:", playerName)

	return playerName
}

func PerformTurn(game *Game) {
	fmt.Println("Your available tries are:", AvailableTries(*game))
	fmt.Println("Your sentence is: ", game.HashedSentence)

	startGuessing(game)
}

func WinnerMessage(game Game) string {
	return fmt.Sprintf("Congratulations %s! You guessed the sentece: \n%s", game.PlayerName, game.Sentence)
}

func LooserMessage(game Game) string {
	return fmt.Sprintf("Sorry %s! You did not guessed the sentece: \n%s", game.PlayerName, game.Sentence)
}

func startGuessing(game *Game) {
	guess, err := getGuess()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Try again.")
		startGuessing(game)
	}

	handleGuess(guess, game, &game.HashedSentence)
}

func getGuess() (string, error) {
	fmt.Println("Make a guess:")
	guess := PlayerInput()
	formatGuess(&guess)

	if isValid(guess) {
		return guess, nil
	}

	return "", errors.New("input is invalid, you can type only single letter")
}

func handleGuess(guess string, game *Game, hashedSentence *[]string) {
	AddTry(game)
	AddGuess(game, guess)
	RevealHashedLetter(game.Sentence, (hashedSentence), guess)

	fmt.Println("Your guess is:", guess)
}

func isValid(guess string) bool {
	if len(guess) != 1 {
		fmt.Printf("%d", len(guess))
		return false
	}
	return isLetter(guess)
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatGuess(guess *string) {
	*guess = strings.ToLower(*guess)
}
