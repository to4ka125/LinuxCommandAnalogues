package main

import (
 "flag"
 "fmt"
 "os"
)

func rmDir(p string) error {
 return os.Remove(p)
}

func isEmpty(p string) (bool, error) {
 f, err := os.Open(p)
 if err != nil {
  return false, err
 }
 defer f.Close()

 contents, err := f.Readdirnames(0)
 if err != nil {
  return false, err
 }
 return len(contents) == 0, nil
}

func main() {
 h := flag.Bool("h", false, "показать помощь")
 p := flag.Bool("p", false, "удалять родительские директории если они пустые")

 flag.Parse()

 if *h {
  fmt.Println("Использование: rmdir [-h] [-p] <директория>")
  fmt.Println("-h: показать помощь")
  fmt.Println("-p: удалять родительские директории, если они пустые")
  os.Exit(0)
 }

 if flag.NArg() < 1 {
  fmt.Println("Ошибка: директория не указана.")
  os.Exit(1)
 }

 tp := flag.Arg(0)

 if !*p {
  if err := rmDir(tp); err != nil {
   fmt.Println("Ошибка при удалении директории:", err)
   os.Exit(1)
  }
  fmt.Println("Директория успешно удалена:", tp)
  return
 }

 empty, err := isEmpty(tp)
 if err != nil {
  fmt.Println("Ошибка при проверке директории:", err)
  os.Exit(1)
 } else if empty {
  if err := rmDir(tp); err != nil {
   fmt.Println("Ошибка при удалении директории:", err)
   os.Exit(1)
  }
  fmt.Println("Директория успешно удалена:", tp)
 } else {
  fmt.Println("Ошибка: директория не пустая:", tp)
  os.Exit(1)
 }
}