package main

import (
 "flag"
 "fmt"
 "os"
 "strconv"
)

func main() {
 lineCountArg := flag.String("n", "25", "Количество пустых строк (целое число)")
 showHelp := flag.Bool("h", false, "Показать справку")

 flag.Parse()

 if *showHelp {
  fmt.Println("Использование: clear [-h] [-n <количество_строк>]")
  os.Exit(0)
 }

 lineCount, err := strconv.Atoi(*lineCountArg)
 if err != nil || lineCount < 0 {
  fmt.Println("Ошибка: Неверный параметр -n. Используйте неотрицательное целое число.")
  os.Exit(1)
 }

 for i := 0; i < lineCount; i++ {
  fmt.Println()
 }
}