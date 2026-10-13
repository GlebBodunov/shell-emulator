package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"shellemu/src/shell"
)

func TestRunScriptSkipsComments(t *testing.T) {
	path := filepath.Join(t.TempDir(), "start.txt")
	text := "// комментарий\n\nbad\nls x\nexit\nls y\n"
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	sh := shell.New(&out)
	if err := sh.RunScript(path); err != nil {
		t.Fatal(err)
	}
	result := out.String()
	if strings.Contains(result, "комментарий") || !strings.Contains(result, "ls x") {
		t.Errorf("неверный вывод: %q", result)
	}
	if strings.Contains(result, "ls y") {
		t.Errorf("команды после exit не должны выполняться: %q", result)
	}
}

func TestRunScriptMissingFile(t *testing.T) {
	sh := shell.New(&bytes.Buffer{})
	if err := sh.RunScript("no_such_file.txt"); err == nil {
		t.Error("ожидалась ошибка")
	}
}
