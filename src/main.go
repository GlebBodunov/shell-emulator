package main

import (
	"flag"
	"fmt"
	"os"

	"shellemu/src/shell"
)

func main() {
	vfsPath := flag.String("vfs", "", "путь к директории с VFS")
	scriptPath := flag.String("script", "", "путь к стартовому скрипту")
	flag.Parse()
	printConfig(*vfsPath, *scriptPath)

	sh := shell.New(os.Stdout)
	if *scriptPath != "" {
		if err := sh.RunScript(*scriptPath); err != nil {
			fmt.Println("ошибка стартового скрипта:", err)
		}
	}
	if !sh.Exited() {
		sh.RunInteractive(os.Stdin)
	}
}

func printConfig(vfsPath, scriptPath string) {
	fmt.Println("[debug] параметры запуска:")
	fmt.Printf("[debug]   -vfs    = %q\n", vfsPath)
	fmt.Printf("[debug]   -script = %q\n", scriptPath)
}
