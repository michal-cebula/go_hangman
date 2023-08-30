package main

import (
	"fmt"
	"go_project/handlers"
	"reflect"
)

func main() {
	fmt.Println("Insert player name:")
	playerName := handlers.GetPlayerName()

	game := handlers.NewGame(playerName)
	handlers.AddHashSentece(&game)

	runGame(&game)
	addSentence()
}

func runGame(game *handlers.Game) {
	for game.Tries != game.MaxTries {
		if !compareOutputSentence(game) {
			handlers.PerformTurn(game)
		} else {
			fmt.Println(handlers.WinnerMessage(*game))
			return
		}
	}
	fmt.Println(game.HashedSentence)
	fmt.Println(handlers.LooserMessage(*game))
}

func addSentence() {
	fmt.Println("Do you want to add a new sentence? (y/n)")
	answer := handlers.PlayerInput()

	if answer == "y" {
		fmt.Println("Type your sentence:")
		sentence := handlers.PlayerInput()
		handlers.AddSentence(sentence)
	} else if answer == "n" {
		fmt.Println("Ok. Untill next time!")
	}
}

func compareOutputSentence(game *handlers.Game) bool {
	return reflect.DeepEqual(game.HashedSentence, handlers.SplitSentence(game.Sentence))
}
