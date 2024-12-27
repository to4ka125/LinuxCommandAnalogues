package main

import (
 "flag"
 "fmt"
 "io/fs"
 "log"
 "path/filepath"
 "strconv"
 "strings"
)

// Функция для проверки, соответствует ли имя файла шаблону
func matchesPattern(fileName, pattern string) bool {
 if strings.Contains(pattern, "*") {
  return strings.HasSuffix(fileName, strings.TrimPrefix(pattern, "*"))
 }
 return fileName == pattern
}

// Функция для поиска файлов в директории
func searchFiles(dir, pattern string, minSize int64) {
 err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
  if err != nil {
   return err
  }
  if entry.IsDir() {
   return nil
  }

  info, err := entry.Info()
  if err != nil {
   return err
  }

  if info.Size() < minSize {
   return nil
  }

  if matchesPattern(info.Name(), pattern) {
   fmt.Println(path)
  }
  return nil
 })

 if err != nil {
  log.Fatalf("Ошибка при обходе директории: %v", err)
 }
}

func main() {
 showHelp := flag.Bool("h", false, "Показать справку")
 dirPath := flag.String("d", "", "Путь к директории для поиска")
 filePattern := flag.String("n", "", "Имя файла или шаблон для поиска")
 minFileSize := flag.String("s", "0", "Минимальный размер файла в байтах (по умолчанию 0)")

 flag.Parse()

 if *showHelp {
  printHelp()
  return
 }

 if *dirPath == "" {
  log.Fatal("Ошибка: Необходимо указать путь к директории с помощью -d.")
 }

 if *filePattern == "" {
  log.Fatal("Ошибка: Необходимо указать имя файла или шаблон с помощью -n.")
 }

 size, err := strconv.ParseInt(*minFileSize, 10, 64)
 if err != nil || size < 0 {
  log.Fatal("Ошибка: Размер должен быть неотрицательным числом.")
 }

 searchFiles(*dirPath, *filePattern, size)
}

// Функция для вывода справки по использованию программы
func printHelp() {
 fmt.Println("Использование: find -d <директория> -n <имя> -s <размер>")
 fmt.Println("-h: Показать справку")
 fmt.Println("-d: Путь к директории для поиска")
 fmt.Println("-n: Имя файла или шаблон для поиска")
 fmt.Println("-s: Минимальный размер файла в байтах (по умолчанию 0)")
}
