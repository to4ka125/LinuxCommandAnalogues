package main

import (
 "archive/zip"
 "fmt"
 "io"
 "os"
 "path/filepath"
)

func displayHelp() {
 fmt.Println("Опции:")
 fmt.Println("-h Показать помощь.")
 fmt.Println("-v Включить подробный вывод.")
 fmt.Println("-d Удалить указанные файлы после архивации.")
 fmt.Println("-u Обновить существующий архив, добавив новые файлы.")
 fmt.Println("-p Рекурсивно сжать содержимое каталога.")
 fmt.Println("Использование: zip [опция] [имя.zip] [файлы для архивации...]")
}

func main() {
 if len(os.Args) < 3 {
  displayHelp()
  return
 }

 var verbose bool
 var archiveName string
 var filesToZip []string
 var deleteAfterArchiving bool
 var updateArchive bool
 var recursiveCompress bool

 for _, arg := range os.Args[1:] {
  switch arg {
  case "-h":
   displayHelp()
   return
  case "-v":
   verbose = true
  case "-d":
   deleteAfterArchiving = true
  case "-u":
   updateArchive = true
  case "-p":
   recursiveCompress = true
  default:
   if archiveName == "" {
    archiveName = arg
   } else {
    filesToZip = append(filesToZip, arg)
   }
  }
 }

 if len(filesToZip) == 0 {
  fmt.Println("Ошибка: Не указаны файлы для архивации.")
  displayHelp()
  return
 }

 var zipWriter *zip.Writer
 var err error

 if updateArchive {
  existingArchive, err := os.OpenFile(archiveName, os.O_RDWR|os.O_CREATE, 0666)
  if err != nil {
   fmt.Println("Ошибка при открытии архива:", err)
   return
  }
  defer existingArchive.Close()

  zipWriter = zip.NewWriter(existingArchive)
  defer zipWriter.Close()
 } else {
  outputZip, err := os.Create(archiveName)
  if err != nil {
   fmt.Println("Ошибка при создании архива:", err)
   return
  }
  defer outputZip.Close()

  zipWriter = zip.NewWriter(outputZip)
  defer zipWriter.Close()
 }

 for _, file := range filesToZip {
  if recursiveCompress && isDirectory(file) {
   err = filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
    if err != nil {
     return err
    }
    if !info.IsDir() {
     addFileToZip(zipWriter, path, file)
     if deleteAfterArchiving {
      os.Remove(path)
     }
    }
    return nil
   })
   if err != nil {
    fmt.Println("Ошибка при рекурсивном сжатии:", err)
    return
   }
  } else {
   addFileToZip(zipWriter, file, "")
   if deleteAfterArchiving {
    os.Remove(file)
   }
  }
 }

 if verbose {
  fmt.Printf("Архивирование завершено: %s\n", archiveName)
 }
}

func addFileToZip(zipWriter *zip.Writer, filePath string, basePath string) {
 inputFile, err := os.Open(filePath)
 if err != nil {
  fmt.Println("Ошибка при открытии файла:", err)
  return
 }
 defer inputFile.Close()

 var zipEntry string
 if basePath != "" {
  zipEntry = filepath.Join(filepath.Base(basePath), filepath.Base(filePath))
 } else {
  zipEntry = filepath.Base(filePath)
 }

 entryWriter, err := zipWriter.Create(zipEntry)
 if err != nil {
  fmt.Println("Ошибка при создании записи в архиве:", err)
  return
 }

 if _, err := io.Copy(entryWriter, inputFile); err != nil {
  fmt.Println("Ошибка при копировании файла в архив:", err)
  return
 }
}

func isDirectory(path string) bool {
 info, err := os.Stat(path)
 return err == nil && info.IsDir()
}