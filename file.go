package main

import (
 "flag"
 "fmt"
 "os"
 "path/filepath"
)

func getFileType(name string) string {
 ext := filepath.Ext(name)

 switch ext {
 case ".txt":
  return "Текстовый файл"
 case ".jpg", ".jpeg":
  return "Изображение JPEG"
 case ".png":
  return "Изображение PNG"
 case ".go":
  return "Файл Go"
 case ".pdf":
  return "PDF документ"
 case ".zip":
  return "ZIP архив"
 case ".tar":
  return "TAR архив"
 default:
  return "Неизвестный тип файла"
 }
}

func main() {
 h := flag.Bool("h", false, "показать помощь")
 flag.Parse()

 if *h {
  fmt.Println("Использование: file [-h] <файл1> <файл2> <файл3>")
  os.Exit(0)
 }

 if flag.NArg() < 1 {
  fmt.Println("Ошибка: необходимо указать хотя бы один файл.")
  os.Exit(1)
 }

 for _, f := range flag.Args() {
  if _, err := os.Stat(f); err != nil {
   fmt.Printf("Ошибка: файл %s не найден.\n", f)
   continue
  }
  fmt.Printf("%s: %s\n", f, getFileType(f))
 }
}