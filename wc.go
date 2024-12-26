package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)
func help() {
 fmt.Println("Справка")
 fmt.Println("Использование: wc [опция] [файл]")
 fmt.Println("Опции:")
 fmt.Println("  -c  Отобразить количество байтов.")
 fmt.Println("  -l  Отобразить количество строк.")
 fmt.Println("  -w  Отобразить количество слов.")
 fmt.Println("  -h  Показать справку.")
}

func countWords(line string) int {
 return len(strings.Fields(line))
}

func wc(filePath string, countBytes, countLines, countWordsFlag bool) error {
 file, err := os.Open(filePath)
 if err != nil {
  return fmt.Errorf("ошибка при открытии файла %s: %v", filePath, err)
 }
 defer file.Close()

 var byteCount, lineCount, wordCount int
 scanner := bufio.NewScanner(file)

 for scanner.Scan() {
  line := scanner.Text()
  lineCount++
  byteCount += len(line) + 1 
  wordCount += countWords(line)
 }

 if err := scanner.Err(); err != nil {
  return fmt.Errorf("ошибка при чтении файла %s: %v", filePath, err)
 }

 // Форматированный вывод
 fmt.Printf("Файл: %s\n", filePath)
 if countBytes {
  fmt.Printf("Количество байтов: %d\n", byteCount)
 }
 if countLines {
  fmt.Printf("Количество строк: %d\n", lineCount)
 }
 if countWordsFlag {
  fmt.Printf("Количество слов: %d\n", wordCount)
 }

 return nil
}

func main() {
 if len(os.Args) < 2 {
  help()
  return
 }

 countBytes, countLines, countWordsFlag := false, false, false
 var filePath string

 for _, arg := range os.Args[1:] {
  switch arg {
  case "-h":
   help()
   return
  case "-c":
   countBytes = true
  case "-l":
   countLines = true
  case "-w":
   countWordsFlag = true
  default:
   if filePath == "" {
    filePath = arg
   } else {
    fmt.Println("Ошибка: можно указать только один файл.")
    help()
    return
   }
  }
 }

 if !countBytes && !countLines && !countWordsFlag {
  countBytes, countLines, countWordsFlag = true, true, true
 }

 if filePath == "" {
  fmt.Println("Ошибка: необходимо указать имя файла.")
  return
 }

 if err := wc(filePath, countBytes, countLines, countWordsFlag); err != nil {
  fmt.Println(err)
 }
}