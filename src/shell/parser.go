package shell

import (
	"os"
	"strings"
)

func Parse(line string) (string, []string) {
	words := strings.Fields(line)
	if len(words) == 0 {
		return "", nil
	}
	for i, word := range words {
		words[i] = os.ExpandEnv(word)
	}
	return words[0], words[1:]
}
