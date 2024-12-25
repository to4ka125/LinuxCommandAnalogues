package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// getHomeDir возвращает домашнюю директорию пользователя
func getHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

// getLastCommand читает последнюю команду из файла .bash_history
func getLastCommand() (string, error) {
	// Получаем домашнюю директорию пользователя
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}

	// Путь к файлу .bash_history
	historyFilePath := filepath.Join(homeDir, ".bash_history")

	// Открываем файл .bash_history
	file, err := os.OpenFile(historyFilePath, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			// Если файл не существует, просто возвращаем nil
			return "", nil
		}
		return "", err
	}
	defer file.Close()

	// Читаем содержимое файла построчно
	var lastCmd string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastCmd = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Проверяем, начинается ли последняя команда с "./ll"
	if strings.HasPrefix(lastCmd, "./ll") {
		return "", nil // Возвращаем пустую строку, если команда не должна учитываться
	}

	return lastCmd, nil
}

// help выводит справочную информацию
func help() {
	fmt.Println("Справка")
	fmt.Println("Использование: ./!!")
	fmt.Println("Выводит последнюю введенную команду из .bash_history.")
	fmt.Println("Опции:")
	fmt.Println("  -h    Показать эту справку")
}

func main() {
	// Проверяем аргументы командной строки
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		help()
		return
	}

	cmd, err := getLastCommand()
	if err != nil {
		fmt.Println("Ошибка при получении последней команды:", err)
		return
	}

	if cmd == "" {
		fmt.Println("Ошибка: нет последней команды для отображения.")
		return
	}

	// Выводим последнюю команду
	fmt.Printf("Последняя введенная команда: %s\n", cmd)
}
