package main

import (
 "bufio"
 "fmt"
 "os"
 "time"
)

func showHelp() {
 fmt.Println("Использование: date [опции]")
 fmt.Println("Опции:")
 fmt.Println("  -d <date>  Печатает указанную дату")
 fmt.Println("  -f <file>  Читает даты из файла и печатает их")
 fmt.Println("  -r <file>  Читает дату из файла и печатает её в стандартном формате")
 fmt.Println("  -h         Показать эту справку")
}

func printDate(dateStr string) {
 date, err := time.Parse(time.RFC3339, dateStr)
 if err != nil {
  fmt.Printf("Ошибка: Неверный формат даты '%s'. Используйте формат RFC3339.\n", dateStr)
  return
 }
 fmt.Println("Указанная дата:", date.Format(time.RFC1123))
}

func processFile(filename string) {
 file, err := os.Open(filename)
 if err != nil {
  fmt.Printf("Ошибка при открытии файла '%s': %v\n", filename, err)
  return
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
  printDate(scanner.Text())
 }

 if err := scanner.Err(); err != nil {
  fmt.Printf("Ошибка при чтении файла: %v\n", err)
 }
}
func main() {
 if len(os.Args) < 3 {
  showHelp()
  return
 }

 switch os.Args[1] {
 case "-d":
  printDate(os.Args[2])
 case "-f":
  processFile(os.Args[2])
 case "-r":
  processFile(os.Args[2]) 
 case "-h":
  showHelp()
 default:
  showHelp()
 }
}
