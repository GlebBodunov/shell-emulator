#!/bin/bash
# Проверка работы эмулятора: VFS с несколькими файлами.
rm -rf out
go run ./src -vfs vfs/files -script startup/stage3.txt
echo "--- содержимое сохранённой VFS на диске:"
ls -R out/saved