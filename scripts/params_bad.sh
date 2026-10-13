#!/bin/bash
# Запуск с несуществующим стартовым скриптом: ожидается сообщение об ошибке.
echo exit | go run ./src -script startup/missing.txt