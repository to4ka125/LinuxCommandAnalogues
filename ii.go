package main

import (
 "flag"
 "fmt"
 "os"
 "os/exec"
 "strings"
)

func main() {
 h := flag.Bool("h", false, "Выводит справку")
 flag.Parse()

 if *h {
  fmt.Println("Использование: go run main.go [опции]")
  fmt.Println("Опции:")
  fmt.Println("  -h    Выводит справку")
  return
 }

 cmd := getLastCmd()
 if cmd == "" {
  fmt.Println("Нет доступных команд в истории.")
  return
 }

 fmt.Printf("Выполняется: %s\n", cmd)
 execCmd := exec.Command("sh", "-c", cmd)
 output, err := execCmd.CombinedOutput()
 if err != nil {
  fmt.Printf("Ошибка: %s\n", err)
  return
 }

 fmt.Println(string(output))
}

func getLastCmd() string {
 histFile := os.Getenv("HISTFILE")
 if histFile == "" {
  return ""
 }

 data, err := os.ReadFile(histFile)
 if err != nil {
  return ""
 }

 lines := strings.Split(string(data), "\n")
 if len(lines) == 0 {
  return ""
 }

 for i := len(lines) - 1; i >= 0; i-- {
  if strings.TrimSpace(lines[i]) != "" {
   return lines[i]
  }
 }
 return ""
}
