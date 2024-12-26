package main

import (
 "archive/tar"
 "compress/gzip"
 "flag"
 "fmt"
 "io"
 "os"
 "path/filepath"
)

func main() {
 // Определяем флаги командной строки
 create := flag.Bool("c", false, "Создать архив")
 extract := flag.Bool("x", false, "Распаковать архив")
 archiveName := flag.String("f", "archive.tar.gz", "Имя архива (с .gz для gzip)")
 showHelp := flag.Bool("h", false, "Показать справку")

 flag.Parse()

 if *showHelp {
  printHelp()
  os.Exit(0)
 }

 if *create && *extract {
  fmt.Println("Ошибка: Флаги -c и -x не могут использоваться одновременно.")
  os.Exit(1)
 }

 if flag.NArg() == 0 {
  fmt.Println("Ошибка: Укажите файлы.")
  os.Exit(1)
 }

 files := flag.Args()

 if *create {
  createTarGz(*archiveName, files)
 } else if *extract {
  extractTarGz(*archiveName)
 } else {
  fmt.Println("Ошибка: Не указано действие (c или x).")
  os.Exit(1)
 }
}

func printHelp() {
 fmt.Println("Использование: tar [-c|-x] [-f <имя_архива>] [-h] <файлы>")
 fmt.Println("-c: создать архив")
 fmt.Println("-x: распаковать архив")
 fmt.Println("-f: имя архива (с .gz для gzip)")
 fmt.Println("-h: справка")
}

func createTarGz(filename string, files []string) {
 file, err := os.Create(filename)
 if err != nil {
  fmt.Println("Ошибка создания архива:", err)
  return
 }
 defer file.Close()

 gzipWriter := gzip.NewWriter(file)
 defer gzipWriter.Close()

 tarWriter := tar.NewWriter(gzipWriter)
 defer tarWriter.Close()

 for _, filePath := range files {
  if err := addFileToTar(tarWriter, filePath); err != nil {
   fmt.Println("Ошибка добавления файла:", err)
   return
  }
 }

 fmt.Println("Архив создан:", filename)
}

func addFileToTar(tw *tar.Writer, filePath string) error {
 info, err := os.Lstat(filePath)
 if err != nil {
  return err
 }

 header, err := tar.FileInfoHeader(info, "")
 if err != nil {
  return err
 }
 header.Name = filePath

 if err := tw.WriteHeader(header); err != nil {
  return err
 }

 if !info.IsDir() {
  file, err := os.Open(filePath)
  if err != nil {
   return err
  }
  defer file.Close()

  if _, err := io.Copy(tw, file); err != nil {
   return err
  }
 }
 return nil
}

func extractTarGz(filename string) {
 file, err := os.Open(filename)
 if err != nil {
  fmt.Println("Ошибка открытия архива:", err)
  return
 }
 defer file.Close()

 gzipReader, err := gzip.NewReader(file)
 if err != nil {
  fmt.Println("Ошибка чтения архива:", err)
  return
 }
 defer gzipReader.Close()

 tarReader := tar.NewReader(gzipReader)

 for {
  header, err := tarReader.Next()
  if err == io.EOF {
   break
  }
  if err != nil {
   fmt.Println("Ошибка чтения архива:", err)
   return
  }
  
  if err := extractFileFromTar(tarReader, header); err != nil {
   fmt.Println("Ошибка извлечения файла:", err)
   return
  }
 }

 fmt.Println("Архив распакован:", filename)
}

func extractFileFromTar(tr *tar.Reader, header *tar.Header) error {
 path := header.Name
 info := header.FileInfo()

 if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
  return err
 }

 if info.IsDir() {
  return os.MkdirAll(path, info.Mode())
 }

 file, err := os.Create(path)
 if err != nil {
  return err
 }
 defer file.Close()

 if _, err = io.Copy(file, tr); err != nil {
  return err
 }
 return nil
}
