package main

import (
 "fmt"
 "os"
 "strconv"
)

func main() {
 if len(os.Args) < 2 {
  os.Exit(0)
 }

 switch os.Args[1] {
 case "-h":
  printHelp()
  os.Exit(0)
 default:
  processExit(os.Args)
 }
}

func processExit(args []string) {
 if len(args) == 2 {
  os.Exit(0)
 }

 if len(args) == 3 {
  code, err := strconv.Atoi(args[2])
  if err != nil {
   fmt.Println("Ошибка: Код выхода должен быть числом.")
   os.Exit(1)
  }
  os.Exit(code)
 }

 if len(args) > 3 {
  fmt.Println("Ошибка: Неверное количество аргументов.")
  printHelp()
  os.Exit(1)
 }
}

func printHelp() {
 fmt.Println("Справка")
 fmt.Println("Использование: exit [опции]")
 fmt.Println("Опции:")
 fmt.Println(" -h     Выводит данную справку.")
}