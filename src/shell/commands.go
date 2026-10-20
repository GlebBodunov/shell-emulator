package shell

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const singleArg = 1

type command func(s *Shell, args []string) error

var commands = map[string]command{
	"ls":       cmdLs,
	"cd":       cmdCd,
	"exit":     cmdExit,
	"vfs-save": cmdVfsSave,
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

func cmdVfsSave(s *Shell, args []string) error {
	if len(args) != singleArg {
		return errors.New("vfs-save: укажите один путь для сохранения")
	}
	if _, err := os.Stat(args[0]); err == nil {
		return fmt.Errorf("vfs-save: %s: путь уже существует", args[0])
	}
	if err := SaveVFS(s.root, args[0]); err != nil {
		return fmt.Errorf("vfs-save: %w", err)
	}
	fmt.Fprintln(s.out, "VFS сохранена в", args[0])
	return nil
}
