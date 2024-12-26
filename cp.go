package main

import (
 "fmt"
 "io"
 "os"
 "path/filepath"
 "strings"
)

// showHelp выводит справку по использованию программы
func showHelp() {
 fmt.Println("Справка")
 fmt.Println("Использование:")
 fmt.Println("cp [опции] [откуда] [куда]")
 fmt.Println()
 fmt.Println("Опции:")
 fmt.Println("  -h          Показать эту справку")
 fmt.Println("  -a          Копировать рекурсивно и сохранять атрибуты")
 fmt.Println("  -i          Спрашивать перед тем как переписывать")
 fmt.Println("  -v          Показать версию программы")
}

// showVersion выводит информацию о версии программы
func showVersion() {
 fmt.Println("cp (GNU coreutils) 9.1")
 fmt.Println("Copyright (C) 2022 Free Software Foundation, Inc.")
 fmt.Println("Лицензия: MIT")
}

// copyFile выполняет копирование файла
func copyFile(src, dst string, preserveAttrs bool) error {
 sourceFile, err := os.Open(src)
 if err != nil {
  return err
 }
 defer sourceFile.Close()

 destFile, err := os.Create(dst)
 if err != nil {
  return err
 }
 defer destFile.Close()

 if _, err = io.Copy(destFile, sourceFile); err != nil {
  return err
 }

 if preserveAttrs {
  srcInfo, err := sourceFile.Stat()
  if err != nil {
   return err
  }
  return os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime())
 }

 return nil
}

// copyDir выполняет рекурсивное копирование каталога
func copyDir(src, dst string, preserveAttrs bool) error {
 return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
  if err != nil {
   return err
  }

  relPath, err := filepath.Rel(src, path)
  if err != nil {
   return err
  }
  destPath := filepath.Join(dst, relPath)

  if info.IsDir() {
   return os.MkdirAll(destPath, os.ModePerm)
  }

  return copyFile(path, destPath, preserveAttrs)
 })
}

func main() {
 if len(os.Args) < 2 {
  showHelp()
  return
 }

 var interactive, preserveAttrs bool

 for _, arg := range os.Args[1:] {
  switch arg {
  case "-h":
   showHelp()
   return
  case "-v":
   showVersion()
   return
  case "-a":
   preserveAttrs = true
  case "-i":
   interactive = true
  default:
   if strings.HasPrefix(arg, "-") {
    fmt.Printf("Неизвестный аргумент: %s\n", arg)
    showHelp()
    return
   }
  }
 }

 if len(os.Args) < 3 {
  fmt.Println("Ошибка: Не указаны источник и назначение.")
  showHelp()
  return
 }

 src := os.Args[len(os.Args)-2]
 dst := os.Args[len(os.Args)-1]

 if interactive {
  fmt.Printf("Вы действительно хотите скопировать %s в %s? (y/n): ", src, dst)
  var response string
  fmt.Scanln(&response)
  if response != "y" {
   fmt.Println("Копирование отменено.")
   return
  }
 }

 srcInfo, err := os.Stat(src)
 if err != nil {
  fmt.Printf("Ошибка: %v\n", err)
  return
 }

 if srcInfo.IsDir() {
  err = copyDir(src, dst, preserveAttrs)
 } else {
  err = copyFile(src, dst, preserveAttrs)
 }

 if err != nil {
  fmt.Printf("Ошибка при копировании: %v\n", err)
 }
}
