package handlers

type Game struct {
	PlayerName     string
	Sentence       string
	HashedSentence []string
	Tries          int
	MaxTries       int
	Guesses        []string
}

func NewGame(playerName string) Game {
	sentence := GetRandomSentence()

	return Game{
		PlayerName: playerName,
		Sentence:   sentence,
		Tries:      0,
		MaxTries:   calculateMaxTries(sentence),
		Guesses:    []string{},
	}
}

func AvailableTries(game Game) int {
	return game.MaxTries - game.Tries
}

func AddTry(game *Game) {
	game.Tries = (*game).Tries + 1
}

func AddGuess(game *Game, guess string) {
	game.Guesses = append((*game).Guesses, guess)
}

func AddHashSentece(game *Game) {
	game.HashedSentence = HashSentece((*game).Sentence)
}

func calculateMaxTries(sentence string) int {
	unique_letters := CountLetters(sentence)
	return unique_letters + unique_letters/6
}
