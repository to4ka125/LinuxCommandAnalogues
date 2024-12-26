package main

import (
 "bufio"
 "fmt"
 "os"
 "os/exec"
 "strings"
)

func printHelp() {
 fmt.Println("Использование: go run main.go [команда]")
 fmt.Println("Команда: !! - выполняет последнюю введённую команду.")
 fmt.Println("Флаги:")
 fmt.Println("  -h, --help  Показать эту справку")
}

func main() {
 if len(os.Args) > 1 {
  if os.Args[1] == "-h" || os.Args[1] == "--help" {
   printHelp()
   return
  }
 }

 historyFile := os.Getenv("HOME") + "/.bash_history"
 file, err := os.Open(historyFile)
 if err != nil {
  fmt.Println("Ошибка при открытии файла истории:", err)
  return
 }
 defer file.Close()

 var lastCommand string
 scanner := bufio.NewScanner(file)

 for scanner.Scan() {
  lastCommand = scanner.Text() // Считываем последнюю команду
 }

 if err := scanner.Err(); err != nil {
  fmt.Println("Ошибка при чтении файла истории:", err)
  return
 }

 if lastCommand == "" {
  fmt.Println("История команд пуста.")
  return
 }

 fmt.Println("Выполнение последней команды:", lastCommand)
 cmd := exec.Command("bash", "-c", lastCommand) // Выполняем команду в оболочке
 cmd.Stdout = os.Stdout
 cmd.Stderr = os.Stderr

 if err := cmd.Run(); err != nil {
  fmt.Println("Ошибка при выполнении команды:", err)
 }
}