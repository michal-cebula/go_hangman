package handlers

import (
	"math/rand"
	"strings"
)

var Sentences = []string{
	"He was sure the Devil created red sparkly glitter",
	"It took him a month to finish the meal",
	"The team members were hard to tell apart since they all wore their hair in a ponytail",
	"Cursive writing is the best way to build a race track",
	"Truth in advertising and dinosaurs with skateboards have much in common",
}

func AddSentence(sentence string) {
	Sentences = append(Sentences, sentence)
	removeDuplicateStr(&Sentences)
}

func AddSentences(sentences []string) {
	Sentences = append(Sentences, sentences...)
	removeDuplicateStr(&Sentences)
}

func GetRandomSentence() string {
	index := rand.Intn(len(Sentences))
	return Sentences[index]
}

func HashSentece(sentence string) []string {
	splitedSentence := SplitSentence(sentence)
	for i, char := range splitedSentence {
		if char != " " {
			splitedSentence[i] = "*"
		}
	}

	return splitedSentence
}

func CountLetters(sentence string) int {
	splittedSentence := SplitSentence(sentence)
	removeDuplicateStr(&splittedSentence)
	return len(splittedSentence)
}

func RevealHashedLetter(sentence string, hashedSentence *[]string, guess string) []string {
	splitedSentence := SplitSentence(sentence)
	for i, char := range splitedSentence {
		if char == guess {
			(*hashedSentence)[i] = splitedSentence[i]
		}
	}
	return *hashedSentence
}

func SplitSentence(sentence string) []string {
	sentence = strings.ToLower(sentence)
	return strings.Split(sentence, "")
}

func removeDuplicateStr(splitedSentence *[]string) {
	seenStr := make(map[string]bool)
	result := []string{}
	for _, sentence := range *splitedSentence {

		if value := seenStr[sentence]; !value {
			seenStr[sentence] = true
			result = append(result, sentence)
		}
	}
	*splitedSentence = result
}
