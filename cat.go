package main

import (
 "bufio"
 "flag"
 "fmt"
 "os"
)

func readFile(fn string, showAll, numNonEmpty, numAll, addDollar bool) error {
 f, err := os.Open(fn)
 if err != nil {
  return err
 }
 defer f.Close()

 scanner := bufio.NewScanner(f)
 lineNum := 1

 for scanner.Scan() {
  line := scanner.Text()
  if numAll || (numNonEmpty && line != "") {
   fmt.Printf("%d\t", lineNum)
   lineNum++
  }

  if addDollar {
   fmt.Println(line + "$")
  } else {
   fmt.Println(line)
  }

  if showAll {
   fmt.Printf("[showAll] %s\n", line)
  }
 }

 return scanner.Err()
}

func main() {
 h := flag.Bool("h", false, "показать помощь")
 showAll := flag.Bool("A", false, "выводить все символы")
 numNonEmpty := flag.Bool("b", false, "нумеровать непустые строки")
 numAll := flag.Bool("n", false, "нумеровать все строки")
 addDollar := flag.Bool("e", false, "добавить символ $ в конце каждой строки")

 flag.Parse()

 if *h {
  fmt.Println("Использование: cat [-h] [-A] [-b] [-n] [-e] <файл>")
  os.Exit(0)
 }

 if flag.NArg() < 1 {
  fmt.Println("Ошибка: файл не указан.")
  os.Exit(1)
 }

 fn := flag.Arg(0)
 if err := readFile(fn, *showAll, *numNonEmpty, *numAll, *addDollar); err != nil {
  fmt.Println("Ошибка при чтении файла:", err)
  os.Exit(1)
 }
}
