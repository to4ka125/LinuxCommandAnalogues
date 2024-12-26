package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// help выводит справку по использованию программы
func help() {
	fmt.Println("Справка")
	fmt.Println("Использование:")
	fmt.Println("cp [опции] [откуда] [куда]")
	fmt.Println()
	fmt.Println("Опции:")
	fmt.Println("  -h          Показать эту справку")
	fmt.Println("  -a          Копировать рекурсивно и сохранять атрибуты")
	fmt.Println("  -i          Спрашивать перед тем как переписывать")
	fmt.Println("  -v          Показать версию программы")
}

// vers выводит информацию о версии программы
func vers() {
	fmt.Println("cp (GNU coreutils) 9.1")
	fmt.Println("Copyright (C) 2022 Free Software Foundation, Inc.")
	fmt.Println("Лицензия: MIT")
}

// copyFile выполняет копирование файла
func copyFile(source, destination string, archive bool) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	if archive {
		// Сохранение атрибутов (например, времени модификации)
		srcInfo, err := srcFile.Stat()
		if err != nil {
			return err
		}
		err = os.Chtimes(destination, srcInfo.ModTime(), srcInfo.ModTime())
		if err != nil {
			return err
		}
	}

	return nil
}

// copyDirectory выполняет рекурсивное копирование каталога
func copyDirectory(source, destination string, archive bool) error {
	return filepath.Walk(source, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Определяем путь назначения
		relPath, err := filepath.Rel(source, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destination, relPath)

		if info.IsDir() {
			// Создаем директорию назначения
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// Копируем файл
		return copyFile(srcPath, destPath, archive)
	})
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	var interactive bool
	var archive bool

	// Обработка аргументов командной строки
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h":
			help()
			return
		case "-v":
			vers()
			return
		case "-a":
			archive = true
		case "-i":
			interactive = true
		default:
			// Если это не опция, то это должно быть либо источник, либо назначение
			if strings.HasPrefix(arg, "-") {
				fmt.Printf("Неизвестный аргумент: %s\n", arg)
				help()
				return
			}
		}
	}

	// Проверяем, что указаны как минимум два аргумента (источник и назначение)
	if len(os.Args) < 3 {
		fmt.Println("Ошибка: Не указаны источник и назначение.")
		help()
		return
	}

	// Получаем источник и назначение
	source := os.Args[len(os.Args)-2]
	destination := os.Args[len(os.Args)-1]

	// Логика интерактивного режима
	if interactive {
		fmt.Printf("Вы действительно хотите скопировать %s в %s? (y/n): ", source, destination)
		var response string
		fmt.Scanln(&response)
		if response != "y" {
			fmt.Println("Копирование отменено.")
			return
		}
	}

	// Проверяем, является ли источник каталогом или файлом
	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	if sourceInfo.IsDir() {
		// Если источник - это каталог, выполняем рекурсивное копирование
		err = copyDirectory(source, destination, archive)
	} else {
		// Если источник - это файл, выполняем копирование файла
		err = copyFile(source, destination, archive)
	}

	if err != nil {
		fmt.Printf("Ошибка при копировании: %v\n", err)
	}
}

