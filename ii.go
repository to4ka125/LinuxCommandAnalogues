package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec" // Импортируем пакет os/exec для выполнения команд
	"strconv"
)

// Функция для чтения истории команд из файла
func readHistory(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var history []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		history = append(history, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return history, nil
}

// Функция для выполнения команды
func executeCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Функция для вывода справки
func printHelp() {
	fmt.Println("Использование: repeat_command -h | -n <номер>")
	fmt.Println("-h: Показать справку")
}

// Основная функция
func main() {
	if len(os.Args) != 3 {
		printHelp()
		os.Exit(1)
	}

	var commandNumber int

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-h":
			printHelp()
			os.Exit(0)
		default:
			fmt.Println("Неверный аргумент.")
			os.Exit(1)
		}
	}

	// Чтение истории команд из файла
	historyFile := os.Getenv("HOME") + "/.bash_history" // Путь к файлу истории Bash
	history, err := readHistory(historyFile)
	if err != nil {
		fmt.Printf("Ошибка при чтении истории: %v\n", err)
		os.Exit(1)
	}

	// Проверяем, существует ли команда с указанным номером
	if commandNumber <= 0 || commandNumber > len(history) {
		fmt.Printf("Команда с номером %d не найдена в истории.\n", commandNumber)
		os.Exit(1)
	}

	commandToExecute := history[commandNumber-1] // Получаем команду по номеру

	fmt.Printf("Выполнение команды: %s\n", commandToExecute)

	err = executeCommand(commandToExecute)
	if err != nil {
		fmt.Printf("Ошибка при выполнении команды: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Команда успешно выполнена.")
}
