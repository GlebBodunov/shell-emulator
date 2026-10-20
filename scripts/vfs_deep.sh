#!/bin/bash
# Проверка работы эмулятора: VFS с 4 уровнями вложенности.
rm -rf out
go run ./src -vfs vfs/deep -script startup/stage3.txt
echo "--- содержимое сохранённой VFS на диске:"
ls -R out/saved