package main

import (
 "archive/zip"
 "fmt"
 "io"
 "os"
 "path/filepath"
)

func displayHelp() {
 fmt.Println("Использование: zipextractor -f [архив.zip] -o [папка для извлечения] [-l] [--help]")
 fmt.Println("   -f   Указывает zip-архив для извлечения")
 fmt.Println("   -o   Указывает папку для извлечения файлов")
 fmt.Println("   -l   Вывести список файлов в архиве")
 fmt.Println("   --help   Показать это сообщение")
}

// / извлекает файлы из указанного ZIP-архива в заданную папку.
// Принимает имя ZIP-файла и путь к папке назначения.
// Возвращает ошибку, если что-то пошло не так (например, проблемы с открытием архива или созданием файлов).
func unzipArchive(zipFile string, destination string) error {
 reader, err := zip.OpenReader(zipFile)
 if err != nil {
  return fmt.Errorf("ошибка при открытии архива: %s", err)
 }
 defer reader.Close()

 for _, file := range reader.File {
  outputPath := filepath.Join(destination, file.Name)

  if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
   return fmt.Errorf("ошибка при создании директорий: %s", err)
  }

  inputFile, err := file.Open()
  if err != nil {
   return fmt.Errorf("ошибка при открытии файла внутри архива: %s", err)
  }
  defer inputFile.Close()

  outputFile, err := os.Create(outputPath)
  if err != nil {
   return fmt.Errorf("ошибка при создании файла: %s", err)
  }
  defer outputFile.Close()

  if _, err := io.Copy(outputFile, inputFile); err != nil {
   return fmt.Errorf("ошибка при копировании содержимого: %s", err)
  }
 }
 return nil
}

// выводит список файлов, содержащихся в указанном ZIP-архиве.
// Принимает имя ZIP-файла и возвращает ошибку, если возникли проблемы с открытием архива.
func displayFileList(zipFile string) error {
 reader, err := zip.OpenReader(zipFile)
 if err != nil {
  return fmt.Errorf("ошибка при открытии архива: %s", err)
 }
 defer reader.Close()

 fmt.Println("Список файлов в архиве:")
 for _, file := range reader.File {
  fmt.Println(file.Name)
 }
 return nil
}

func main() {
 args := os.Args

 if len(args) < 5 {
  fmt.Println("Ошибка: недостаточно аргументов.")
  displayHelp()
  return
 }

 var zipArchive string
 var outputFolder string
 var listFiles bool

 for i := 1; i < len(args); i++ {
  switch args[i] {
  case "-f":
   if i+1 < len(args) {
    zipArchive = args[i+1]
    i++
   } else {
    fmt.Println("Ошибка: не указано имя zip-файла.")
    displayHelp()
    return
   }

  case "-o":
   if i+1 < len(args) {
    outputFolder = args[i+1]
    i++
   } else {
    fmt.Println("Ошибка: не указана папка для извлечения.")
    displayHelp()
    return
   }

  case "-l":
   listFiles = true

  case "--help":
   displayHelp()
   return

  default:
   fmt.Println("Ошибка: неизвестный аргумент:", args[i])
   displayHelp()
   return
  }
 }

 if listFiles {
  if err := displayFileList(zipArchive); err != nil {
   fmt.Println("Ошибка:", err)
   return
  }
 } else {
  if err := unzipArchive(zipArchive, outputFolder); err != nil {
   fmt.Println("Ошибка:", err)
   return
  }
  fmt.Println("Файлы успешно извлечены в папку:", outputFolder)
 }
}