package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

func printHelp() {
 fmt.Println("Использование: nl [опции] [файл]")
 fmt.Println("Опции:")
 fmt.Println("  -h        Показать справку.")
 fmt.Println("  -b <mode> Установить режим нумерации строк (a - все строки, t - только ненумерованные).")
 fmt.Println("  -n <mode> Установить формат номера (ln - номер строки, rn - номер с пробелами).")
 fmt.Println("  -w <width> Установить ширину номера строки.")
}

func numberLines(scanner *bufio.Scanner, mode string, width int) {
 lineNumber := 1
 for scanner.Scan() {
  line := scanner.Text()
  if mode == "a" || (mode == "t" && strings.TrimSpace(line) != "") {
   fmt.Printf("%*d  %s\n", width, lineNumber, line)
   lineNumber++
  } else {
   fmt.Println(line)
  }
 }
}

func main() {
 if len(os.Args) < 2 {
  printHelp()
  return
 }

 mode := "a"
 width := 6

 for i := 1; i < len(os.Args); i++ {
  switch os.Args[i] {
  case "-h":
   printHelp()
   return
  case "-b":
   if i+1 < len(os.Args) {
    mode = os.Args[i+1]
    i++
   } else {
    fmt.Println("Ошибка: отсутствует аргумент для -b.")
    return
   }
  case "-n":
   if i+1 < len(os.Args) {
    if os.Args[i+1] == "ln" {
     width = 6
    } else if os.Args[i+1] == "rn" {
     width = 8
    }
    i++
   } else {
    fmt.Println("Ошибка: отсутствует аргумент для -n.")
    return
   }
  case "-w":
   if i+1 < len(os.Args) {
    var err error
    width, err = strconv.Atoi(os.Args[i+1])
    if err != nil {
     fmt.Println("Ошибка: ширина должна быть числом.")
     return
    }
    i++
   } else {
    fmt.Println("Ошибка: отсутствует аргумент для -w.")
    return
   }
  default:
   file, err := os.Open(os.Args[i])
   if err != nil {
    fmt.Println("Ошибка при открытии файла:", err)
    return
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)
   numberLines(scanner, mode, width)
   return
  }
 }

 printHelp()
}