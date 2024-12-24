package main

import (
	"flag"
	"fmt"
	"os"
)

// Функция для вывода справки по использованию программы
func printHelp() {
	fmt.Println("Использование: cd [опции] <путь>")
	fmt.Println("-h: показать помощь")
	fmt.Println("-v: выводить текущую директорию перед переходом")
}

// Функция для изменения директории
func changeDirectory(path string, verbose bool) error {
	if path == "" {
		return fmt.Errorf("путь не может быть пустым")
	}

	_, err := os.Stat(path) // Проверка существования директории

	if err != nil {
		return fmt.Errorf("директория не найдена: %v", err)
	}

	if verbose { // Если включен verbose режим
		currentDir, _ := os.Getwd()
		fmt.Printf("Текущая директория: %s\n", currentDir)
	}

	err = os.Chdir(path) // Пытаемся изменить директорию

	if err != nil {
		return fmt.Errorf("не удалось перейти в директорию: %v", err)
	}

	// Добавлено сообщение об успешном переходе
	fmt.Printf("Успешно перешли в директорию: %s\n", path)

	return nil
}

func main() {
	// Определяем флаги командной строки
	verbose := flag.Bool("v", false, "выводить текущую директорию перед переходом")
	help := flag.Bool("h", false, "показать помощь")

	flag.Parse()

	if *help {
		printHelp()

		os.Exit(0)
	}

	args := flag.Args() // Получаем аргументы командной строки

	if len(args) == 0 { // Проверка на наличие аргументов
		fmt.Println("Ошибка: Необходимо указать путь")

		os.Exit(1)
	}

	path := args[0] // Получаем путь из аргументов

	err := changeDirectory(path, *verbose) // Пытаемся изменить директорию

	if err != nil { // Если произошла ошибка при изменении директории
		fmt.Println("Ошибка:", err)

		os.Exit(1)
	}
}
