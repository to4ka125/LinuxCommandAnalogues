package main

import (
 "bufio"
 "fmt"
 "os"
 "os/exec"
 "strings"
)

// showHelp выводит справочную информацию о программе
func showHelp() {
 fmt.Println("Использование: go run main.go -h <command> [options]")
 fmt.Println("Ключи:")
 fmt.Println("  -h: Показать эту справку.")
 fmt.Println("  -r: Повторить последнюю команду.")
 fmt.Println("  -l: Показать историю команд.")
 fmt.Println("  -e: Выполнить команду с аргументами.")
}

// CommandHistory хранит историю команд
var CommandHistory []string

// executeCommand выполняет переданную команду и сохраняет её в историю
func executeCommand(command string) error {
 cmd := exec.Command("bash", "-c", command)
 cmd.Stdout = os.Stdout
 cmd.Stderr = os.Stderr
 err := cmd.Run()
 if err != nil {
  return fmt.Errorf("ошибка при выполнении команды: %v", err)
 }
 CommandHistory = append(CommandHistory, command) // Сохраняем команду в историю
 return nil
}

// listHistory выводит историю команд
func listHistory() {
 fmt.Println("История команд:")
 for i, cmd := range CommandHistory {
  fmt.Printf("%d: %s\n", i+1, cmd)
 }
}

func main() {
 if len(os.Args) < 3 || os.Args[1] != "-h" {
  showHelp()
  return
 }

 switch os.Args[2] {
 case "-r":
  if len(CommandHistory) == 0 {
   fmt.Println("Ошибка: нет истории команд для повторения.")
   return
  }
  lastCommand := CommandHistory[len(CommandHistory)-1]
  fmt.Printf("Повторение команды: %s\n", lastCommand)
  err := executeCommand(lastCommand)
  if err != nil {
   fmt.Println(err)
  }
 case "-l":
  listHistory()
 case "-e":
  if len(os.Args) < 4 {
   fmt.Println("Ошибка: требуется команда для выполнения.")
   return
  }
  command := strings.Join(os.Args[3:], " ")
  err := executeCommand(command)
  if err != nil {
   fmt.Println(err)
  }
 default:
  fmt.Println("Ошибка: неизвестный ключ. Используйте -h для получения справки.")
 }
}
