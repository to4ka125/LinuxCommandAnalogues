package main

import (
 "fmt"
 "os"
 "runtime"
)

func printArchitecture(args []string) {
 arch := runtime.GOARCH
 fmt.Printf("Архитектура: %s\n", arch)

 if len(args) > 0 {
  fmt.Println("Параметры:")
  for _, arg := range args {
   fmt.Println("-", arg)
  }
 }
}

func showHelp() {
 fmt.Println("Справка")
 fmt.Println("Использование: arch [опция]")
 fmt.Println("Опции:")
 fmt.Println("  -h Показать справку.")
}

func main() {
 if len(os.Args) < 2 {
  printArchitecture(nil)
  return
 }

 switch os.Args[1] {
 case "arch":
  printArchitecture(os.Args[2:])
 case "-h":
  showHelp()
 default:
  fmt.Println("Неизвестная команда. Используйте 'arch -h' для получения справки.")
 }
}
