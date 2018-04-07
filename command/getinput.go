package command

import (
	"bufio"
	"os"
	"strings"
)

func GetInput(wordsLength int) []string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	split := strings.Split(input, " ")
	var ar []string
	for i := 0; i < wordsLength; i++ {
		ar = append(ar, split[i])
	}
	return ar
}
