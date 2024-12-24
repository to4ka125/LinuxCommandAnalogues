package main

import (
 "flag"
 "fmt"
 "os"
 "strconv"
)

func createDir(path string, mode int, verbose bool) error {
 err := os.MkdirAll(path, os.FileMode(mode)) 
 if err != nil {
  return err
 }
 if verbose {
  fmt.Printf("Директория '%s' создана.\n", path)
 }
 return nil
}

func main() {

 help := flag.Bool("h", false, "показать помощь")
 createParents := flag.Bool("p", false, "создать родительские директории")
 modeStr := flag.String("m", "0755", "режим доступа (octal)")
 verbose := flag.Bool("v", false, "включить подробный вывод")

 flag.Parse() 

 if *help { 
  fmt.Println("Использование: mkdir [-h] [-p] [-m <режим>] [-v] <путь>")
  fmt.Println("-h: показать помощь")
  fmt.Println("-p: создавать родительские директории")
  fmt.Println("-m: режим доступа (octal, например 0755)")
  fmt.Println("-v: включить подробный вывод")
  os.Exit(0)
 }

 if flag.NArg() < 1 { 
  fmt.Println("Ошибка: путь не указан. Укажите путь для создания директории.")
  os.Exit(1)
 }

 targetPath := flag.Arg(0) 

 mode, err := strconv.ParseInt(*modeStr, 8, 32) 
 if err != nil || mode < 0 || mode > 0777 {
  fmt.Println("Ошибка: Некорректный режим доступа. Используйте octal значение (например, 0755).")
  os.Exit(1)
 }

 if *createParents {
  if err := createDir(targetPath, int(mode), *verbose); err != nil {
   fmt.Println("Ошибка при создании директории:", err)
   os.Exit(1)
  }
 } else {
  if err := os.Mkdir(targetPath, os.FileMode(int(mode))); err != nil {
   fmt.Println("Ошибка при создании директории:", err)
   os.Exit(1)
  }
  if *verbose {
   fmt.Printf("Директория '%s' создана.\n", targetPath)
  }
 }

 fmt.Println("Директория успешно создана:", targetPath)
}
