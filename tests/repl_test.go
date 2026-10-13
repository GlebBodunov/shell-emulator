package tests

import (
	"bytes"
	"strings"
	"testing"

	"shellemu/src/shell"
)

func TestParseSplitsBySpaces(t *testing.T) {
	name, args := shell.Parse("  ls   -l  /docs ")
	if name != "ls" || strings.Join(args, ",") != "-l,/docs" {
		t.Errorf("получено: %q %q", name, args)
	}
}

func TestParseExpandsEnv(t *testing.T) {
	t.Setenv("MY_VAR", "hello")
	_, args := shell.Parse("cd $MY_VAR ${MY_VAR}/x")
	if args[0] != "hello" || args[1] != "hello/x" {
		t.Errorf("получено: %q", args)
	}
}

func TestStubPrintsArgs(t *testing.T) {
	var out bytes.Buffer
	sh := shell.New(&out)
	if err := sh.Execute("ls a b"); err != nil {
		t.Fatal(err)
	}
	if out.String() != "ls a b\n" {
		t.Errorf("получено: %q", out.String())
	}
}

func TestUnknownCommand(t *testing.T) {
	sh := shell.New(&bytes.Buffer{})
	err := sh.Execute("abc")
	if err == nil || !strings.Contains(err.Error(), "не найдена") {
		t.Errorf("ожидалась ошибка, получено: %v", err)
	}
}

func TestExit(t *testing.T) {
	var out bytes.Buffer
	sh := shell.New(&out)
	sh.RunInteractive(strings.NewReader("exit\nls\n"))
	if !sh.Exited() || strings.Contains(out.String(), "ls \n") {
		t.Errorf("REPL не завершился: %q", out.String())
	}
}
