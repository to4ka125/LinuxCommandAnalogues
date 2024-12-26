package main

import (
 "bufio"
 "fmt"
 "os"
 "os/exec"
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
  line := scanner.Text()
  if line != "" { // игнорируем пустые строки
   history = append(history, line)
  }
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
 fmt.Println("Использование: repeat_command -h | !!")
 fmt.Println("-h: Показать справку")
 fmt.Println("!!: Выполнить последнюю команду из истории")
}

// Основная функция
func main() {
 if len(os.Args) > 2 {
  printHelp()
  os.Exit(1)
 }

 if len(os.Args) == 2 && os.Args[1] == "-h" {
  printHelp()
  os.Exit(0)
 }

 historyFile := os.Getenv("HOME") + "/.bash_history" // Путь к файлу истории Bash
 history, err := readHistory(historyFile)
 if err != nil {
  fmt.Printf("Ошибка при чтении истории: %v\n", err)
  os.Exit(1)
 }

 if len(history) == 0 {
  fmt.Println("История команд пуста.")
  os.Exit(1)
 }

 commandToExecute := history[len(history)-1] // Получаем последнюю команду

 fmt.Printf("Выполнение команды: %s\n", commandToExecute)

 err = executeCommand(commandToExecute)
 if err != nil {
  fmt.Printf("Ошибка при выполнении команды: %v\n", err)
  os.Exit(1)
 }
}
