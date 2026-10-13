package shell

import (
	"bufio"
	"fmt"
	"io"
)

type Shell struct {
	cwd    string
	out    io.Writer
	exited bool
}

func New(out io.Writer) *Shell {
	return &Shell{cwd: "~", out: out}
}

func (s *Shell) Exited() bool {
	return s.exited
}

func (s *Shell) Execute(line string) error {
	name, args := Parse(line)
	if name == "" {
		return nil
	}
	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("%s: команда не найдена", name)
	}
	return cmd(s, args)
}

func (s *Shell) RunInteractive(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for !s.exited {
		fmt.Fprint(s.out, MakePrompt(s.cwd))
		if !scanner.Scan() {
			fmt.Fprintln(s.out)
			return
		}
		if err := s.Execute(scanner.Text()); err != nil {
			fmt.Fprintln(s.out, err)
		}
	}
}
