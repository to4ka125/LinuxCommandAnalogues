package main

import (
 "flag"
 "fmt"
 "os"
 "os/exec"
 "strings"
)

func main() {
 // Определяем флаг для справки
 help := flag.Bool("h", false, "Выводит справку")
 flag.Parse()

 // Если запрашивается справка, выводим её
 if *help {
  fmt.Println("Использование: go run main.go [опции]")
  fmt.Println("Опции:")
  fmt.Println("  -h    Выводит справку")
  return
 }

 // Получаем последнюю введенную команду из истории
 lastCommand := getLastCommand()
 if lastCommand == "" {
  fmt.Println("Нет доступных команд в истории.")
  return
 }

 // Выполняем последнюю команду
 fmt.Printf("Выполняется: %s\n", lastCommand)
 cmd := exec.Command("sh", "-c", lastCommand)
 output, err := cmd.CombinedOutput()
 if err != nil {
  fmt.Printf("Ошибка: %s\n", err)
  return
 }

 // Выводим результат выполнения команды
 fmt.Println(string(output))
}

// getLastCommand возвращает последнюю команду из истории
func getLastCommand() string {
 historyFile := os.Getenv("HISTFILE") // Обычно это ~/.bash_history
 if historyFile == "" {
  return ""
 }

 data, err := os.ReadFile(historyFile)
 if err != nil {
  return ""
 }

 lines := strings.Split(string(data), "\n")
 if len(lines) == 0 {
  return ""
 }

 // Возвращаем последнюю непустую строку (команду)
 for i := len(lines) - 1; i >= 0; i-- {
  if strings.TrimSpace(lines[i]) != "" {
   return lines[i]
  }
 }
 return ""
}