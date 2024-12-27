package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Определение ключей
	physical := flag.Bool("P", false, "Использовать физический путь (без символических ссылок)")
	logical := flag.Bool("L", false, "Использовать логический путь (учитывая символические ссылки)")
	version := flag.Bool("v", false, "Показать версию программы")
	help := flag.Bool("h", false, "Показать справку")

	// Разбор аргументов командной строки
	flag.Parse()

	// Обработка ключа -v (версия)
	if *version {
		fmt.Println("Версия программы: 1.0.0")
		return
	}

	// Обработка ключа -h (справка)
	if *help {
		fmt.Println("Использование: cd [-P] [-L] <путь>")
		fmt.Println("  -P: Использовать физический путь (без символических ссылок)")
		fmt.Println("  -L: Использовать логический путь (учитывая символические ссылки)")
		fmt.Println("  -v: Показать версию программы")
		fmt.Println("  -h: Показать это сообщение справки")
		fmt.Println("\nДля перехода в директорию выполните: cd /путь/к/директории")
		return
	}

	// Проверка наличия пути
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Ошибка: не указан путь. Для перехода в директорию выполните: cd /путь/к/директории")
		return
	}

	path := args[0]

	// Обработка пути в зависимости от ключей
	var finalPath string
	if *physical {
		// Преобразование в физический путь
		absPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			fmt.Printf("Ошибка: не удалось преобразовать путь '%s' в физический: %v", path, err)
			return
		}
		finalPath = absPath
	} else if *logical {
		// Использование логического пути
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Ошибка: не удалось преобразовать путь '%s' в логический: %v", path, err)
			return
		}
		finalPath = absPath
	} else {
		// Если ключи не указаны, просто проверяем существование пути
		finalPath = path
	}

	// Проверка существования директории
	if _, err := os.Stat(finalPath); os.IsNotExist(err) {
		fmt.Printf("Ошибка: директория '%s' не существует. Убедитесь, что путь указан правильно.\n", finalPath)
		return
	}

	// Вывод команды для смены директории
	fmt.Printf("Чтобы перейти в директорию, выполните: cd %s", finalPath)
}

