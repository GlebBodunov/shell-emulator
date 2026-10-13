package shell

import (
	"fmt"
	"os"
	"strings"
)

const commentPrefix = "//"

func (s *Shell) RunScript(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, commentPrefix) {
			continue
		}
		fmt.Fprintln(s.out, MakePrompt(s.cwd)+line)
		if err := s.Execute(line); err != nil {
			fmt.Fprintln(s.out, err)
		}
		if s.exited {
			return nil
		}
	}
	return nil
}
