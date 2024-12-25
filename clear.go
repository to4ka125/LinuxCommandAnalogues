// clear project main.go
package main

import (
	"fmt"
	"os"
)

// Функция для вывода справки по использованию программы
func help() {
	fmt.Println("Справка")
	fmt.Println("Использование: clear [опция]")
	fmt.Println("Опции:")
	fmt.Println(" -h      Выводит  справку.")
	fmt.Println(" -v      Выводит версию программы.")
}

// Функция для очистки экрана
func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape код для очистки экрана
}

// main - точка входа программы
func main() {
	if len(os.Args) < 2 {
		// Если аргументы не указаны, просто очищаем экран
		clearScreen()
		return
	}

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			help() // Выводим справочную информацию
			return
		case "-v":
			fmt.Println("clear версия 1.0.0") // Выводим версию программы
			return
		default:
			fmt.Printf("Неизвестный аргумент: %s\n", os.Args[i])
			return
		}
	}

	// Если не было указано никаких опций, просто очищаем экран
	clearScreen()
}
