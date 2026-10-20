package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

type Node struct {
	Name     string
	IsDir    bool
	Content  []byte
	Children map[string]*Node
}

func NewDir(name string) *Node {
	return &Node{Name: name, IsDir: true, Children: map[string]*Node{}}
}

func LoadVFS(dir string) (*Node, error) {
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%s: не является директорией", dir)
	}
	return loadDir(dir, "")
}

func loadDir(dir, name string) (*Node, error) {
	node := NewDir(name)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		child, err := loadEntry(filepath.Join(dir, entry.Name()), entry)
		if err != nil {
			return nil, err
		}
		node.Children[entry.Name()] = child
	}
	return node, nil
}

func loadEntry(path string, entry os.DirEntry) (*Node, error) {
	if entry.IsDir() {
		return loadDir(path, entry.Name())
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Node{Name: entry.Name(), Content: data}, nil
}

func SaveVFS(node *Node, dir string) error {
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return err
	}
	for name, child := range node.Children {
		if err := saveEntry(child, filepath.Join(dir, name)); err != nil {
			return err
		}
	}
	return nil
}

func saveEntry(node *Node, path string) error {
	if node.IsDir {
		return SaveVFS(node, path)
	}
	return os.WriteFile(path, node.Content, filePerm)
}
