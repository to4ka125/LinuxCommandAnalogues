package main

import (
 "flag"
 "fmt"
 "os"
 "path/filepath"
)

func showHelp() {
 fmt.Println("Использование: du [опции] [путь]")
 fmt.Println("Опции:")
 fmt.Println("  -h          Показать справку.")
 fmt.Println("  -s          Вывести только итоговый размер.")
 fmt.Println("  -a          Показать размер всех файлов и каталогов.")
}

func getSize(path string, all bool) (uint64, error) {
 var totalSize uint64
 err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
  if err != nil {
   return err
  }
  if all || info.IsDir() {
   totalSize += uint64(info.Size())
  }
  return nil
 })
 return totalSize, err
}

func main() {
 showHelpFlag := flag.Bool("h", false, "Показать справку")
 summary := flag.Bool("s", false, "Вывести только итоговый размер")
 all := flag.Bool("a", false, "Показать размер всех файлов и каталогов")

 flag.Parse()

 if *showHelpFlag {
  showHelp()
  return
 }

 path := "."
 if len(flag.Args()) > 0 {
  path = flag.Args()[0]
 }

 size, err := getSize(path, *all)
 if err != nil {
  fmt.Printf("Ошибка: %s\n", err)
  return
 }

 if *summary {
  fmt.Printf("Итоговый размер: %d байт\n", size)
 } else {
  fmt.Printf("Размер каталога '%s': %d байт\n", path, size)
 }
}
