#!/bin/bash
# Проверка работы эмулятора: минимальная VFS (один файл).
rm -rf out
go run ./src -vfs vfs/minimal -script startup/stage3.txt
echo "--- содержимое сохранённой VFS на диске:"
ls -R out/saved