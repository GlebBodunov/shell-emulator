package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"shellemu/src/shell"
)

func TestLoadVFS(t *testing.T) {
	root, err := shell.LoadVFS("../vfs/deep")
	if err != nil {
		t.Fatal(err)
	}
	old := root.Children["docs"].Children["work"].Children["archive"].Children["old.txt"]
	if old == nil || old.IsDir {
		t.Fatal("файл docs/work/archive/old.txt не загружен")
	}
	if string(old.Content) != "старый файл\n" {
		t.Errorf("неверное содержимое: %q", old.Content)
	}
}

func TestLoadVFSMissing(t *testing.T) {
	if _, err := shell.LoadVFS("../vfs/no_such_dir"); err == nil {
		t.Error("ожидалась ошибка")
	}
}

func TestLoadVFSNotDir(t *testing.T) {
	if _, err := shell.LoadVFS("../vfs/minimal/hello.txt"); err == nil {
		t.Error("ожидалась ошибка: это файл, а не папка")
	}
}

func TestVfsSave(t *testing.T) {
	sh := shell.New(&bytes.Buffer{})
	if err := sh.MountVFS("../vfs/files"); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(t.TempDir(), "saved")
	if err := sh.Execute("vfs-save " + target); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(target, "notes.txt"))
	if err != nil {
		t.Fatalf("файл не сохранён: %v", err)
	}
	original, _ := os.ReadFile("../vfs/files/notes.txt")
	if string(data) != string(original) {
		t.Errorf("содержимое отличается: %q", data)
	}
}

func TestVfsSaveErrors(t *testing.T) {
	sh := shell.New(&bytes.Buffer{})
	existing := t.TempDir()
	for _, line := range []string{"vfs-save", "vfs-save a b", "vfs-save " + existing} {
		if err := sh.Execute(line); err == nil {
			t.Errorf("%s: ожидалась ошибка", line)
		}
	}
}
