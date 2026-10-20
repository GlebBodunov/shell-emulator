package shell

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func MakePrompt(dir string) string {
	return fmt.Sprintf("%s@%s:%s$ ", Username(), hostname(), dir)
}

func Username() string {
	u, err := user.Current()
	if err != nil {
		return "user"
	}
	parts := strings.Split(u.Username, `\`)
	return parts[len(parts)-1]
}

func hostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "localhost"
	}
	return name
}
