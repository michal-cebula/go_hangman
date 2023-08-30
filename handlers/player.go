package handlers

import (
	"bufio"
	"os"
	"strings"
)

func PlayerInput() string {
	var playerInput string

	input := bufio.NewReader(os.Stdin)
	playerInput, _ = input.ReadString('\n')
	playerInput = strings.TrimSuffix(playerInput, "\n")

	return playerInput
}
