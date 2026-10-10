#!/bin/zsh

# Проверяем, передана ли ссылка
if [ -z "$1" ]; then
    echo "❌ Ошибка: Передай ссылку на задачу!"
    echo "Использование: ./init_task.sh <ссылка>"
    exit 1
fi

URL=$1

# Вытаскиваем название задачи из конца URL
TASK_NAME=$(echo "$URL" | sed 's/\/$//' | awk -F'/' '{print $NF}')

if [ -z "$TASK_NAME" ]; then
    echo "❌ Ошибка: Не удалось распарсить имя задачи."
    exit 1
fi

# Путь к папке и файлу
TARGET_DIR="coderun/$TASK_NAME"
FILE_PATH="$TARGET_DIR/main.go"

# Создаем папку, если её нет
mkdir -p "$TARGET_DIR"

# Записываем только ссылку-комментарий (перезапишет файл, если он был)
echo "// $URL" > "$FILE_PATH"

echo "🚀 Успешно! Создан файл: $FILE_PATH"

# Открываем в VS Code (если он есть)
if command -v code &> /dev/null; then
    code "$FILE_PATH"
fi
