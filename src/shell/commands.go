package shell

import (
	"fmt"
	"strings"
)

type command func(s *Shell, args []string) error

var commands = map[string]command{
	"ls":   cmdLs,
	"cd":   cmdCd,
	"exit": cmdExit,
}

func cmdLs(s *Shell, args []string) error {
	fmt.Fprintln(s.out, "ls", strings.Join(args, " "))
	return nil
}

func cmdCd(s *Shell, args []string) error {
	fmt.Fprintln(s.out, "cd", strings.Join(args, " "))
	return nil
}

func cmdExit(s *Shell, args []string) error {
	s.exited = true
	return nil
}
