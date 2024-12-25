package main

import (
 "archive/zip"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "strings"
)

const versionInfo = "This is Zip 3.0 (July 5th 2008), by Info-ZIP. 
Currently maintained by E. Gordon. Please send bug reports to 
the authors using the web page at www.info-zip.org; see README for details."


func createZip(archiveName string, sources []string, recursive bool) error {
 zipFile, err := os.Create(archiveName)
 if err != nil {
  return fmt.Errorf("ошибка при создании zip-файла: %v", err)
 }
 defer zipFile.Close()

 zipWriter := zip.NewWriter(zipFile)
 defer zipWriter.Close()

 for _, source := range sources {
  if err := addToZip(zipWriter, source, recursive); err != nil {
   return err
  }
 }
 return nil
}

func addToZip(zipWriter *zip.Writer, source string, recursive bool) error {
 info, err := os.Stat(source)
 if os.IsNotExist(err) {
  return fmt.Errorf("файл или директория %s не существует", source)
 }

 if info.IsDir() {
  if recursive {
   return filepath.Walk(source, func(filePath string, fileInfo os.FileInfo, err error) error {
    if err != nil {
     return err
    }
    if fileInfo.IsDir() {
     return nil
    }
    return addFileToZip(zipWriter, filePath, source)
   })
  }
  return addFileToZip(zipWriter, source, filepath.Dir(source))
 }

 return addFileToZip(zipWriter, source, filepath.Dir(source))
}

func addFileToZip(zipWriter *zip.Writer, filePath, baseDir string) error {
 file, err := os.Open(filePath)
 if err != nil {
  return fmt.Errorf("ошибка при открытии файла %s: %v", filePath, err)
 }
 defer file.Close()

 relativePath := strings.TrimPrefix(filePath, baseDir+string(os.PathSeparator))
 writer, err := zipWriter.Create(relativePath)
 if err != nil {
  return fmt.Errorf("ошибка при создании заголовка zip-файла: %v", err)
 }

 if _, err := io.Copy(writer, file); err != nil {
  return fmt.Errorf("ошибка при копировании содержимого файла %s: %v", filePath, err)
 }

 return nil
}
func deleteFromZip(archiveName string, filesToDelete []string) error {
 r, err := zip.OpenReader(archiveName)
 if err != nil {
  return fmt.Errorf("ошибка при открытии zip-файла: %v", err)
 }
 defer r.Close()

 tempZipFile, err := os.Create(archiveName + ".tmp")
 if err != nil {
  return fmt.Errorf("ошибка при создании временного zip-файла: %v", err)
 }
 defer tempZipFile.Close()

 zipWriter := zip.NewWriter(tempZipFile)
 defer zipWriter.Close()

 for _, f := range r.File {
  if !contains(filesToDelete, f.Name) {
   fileReader, err := f.Open()
   if err != nil {
    return fmt.Errorf("ошибка при открытии файла %s из архива: %v", f.Name, err)
   }
   defer fileReader.Close()

   writer, err := zipWriter.Create(f.Name)
   if err != nil {
    return fmt.Errorf("ошибка при создании файла %s в новом архиве: %v", f.Name, err)
   }

   if _, err := io.Copy(writer, fileReader); err != nil {
    return fmt.Errorf("ошибка при копировании содержимого файла %s: %v", f.Name, err)
   }
  }
 }

 if err := os.Remove(archiveName); err != nil {
  return fmt.Errorf("ошибка при удалении старого zip-файла: %v", err)
 }
 if err := os.Rename(archiveName+".tmp", archiveName); err != nil {
  return fmt.Errorf("ошибка при переименовании временного zip-файла: %v", err)
 }

 return nil
}

func contains(slice []string, item string) bool {
 for _, s := range slice {
  if s == item {
   return true
  }
 }
 return false
}

func help() {
 fmt.Println("Справка")
 fmt.Println("Использование: zip [опции] [файлы/каталоги]")
 fmt.Println("Опции:")
 fmt.Println("  --version  Показать информацию о версии")
 fmt.Println("  -d         Удалить указанные файлы из архива")
 fmt.Println("  -r         Рекурсивно добавлять директории")
}

func main() {
 if len(os.Args) < 2 {
  help()
  return
 }

 switch os.Args[1] {
 case "--version":
  fmt.Println(versionInfo)
 case "-d":
  if len(os.Args) < 3 {
   fmt.Println("Не указаны файлы для удаления.")
   return
  }
  err := deleteFromZip(os.Args[2], os.Args[3:])
  if err != nil {
   fmt.Printf("Ошибка: %v\n", err)
  }
 default:
  files := os.Args[1:]
  err := createZip("archive.zip", files, false) 
  if err != nil {
   fmt.Printf("Ошибка: %v\n", err)
  }
 }
}