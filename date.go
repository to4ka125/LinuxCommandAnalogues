package main

import (
 "bufio"
 "fmt"
 "os"
 "time"
)

func displayHelp() {
 fmt.Println("Использование: mydate [опции]")
 fmt.Println("Опции:")
 fmt.Println("  -d <дата>  Печатает указанную дату")
 fmt.Println("  -f <файл>  Читает даты из файла и печатает их")
 fmt.Println("  -r <файл>  Читает дату из файла и печатает её в стандартном формате")
 fmt.Println("  -h         Показать эту справку")
}

func processDate(dateInput string) {
 if parsedDate, err := time.Parse(time.RFC3339, dateInput); err == nil {
  fmt.Println("Указанная дата:", parsedDate.Format(time.RFC1123))
 } else {
  fmt.Printf("Ошибка: Неверный формат даты '%s'. Используйте формат RFC3339.\n", dateInput)
 }
}

func processFile(filePath string) {
 file, err := os.Open(filePath)
 if err != nil {
  fmt.Printf("Ошибка при открытии файла '%s': %v\n", filePath, err)
  return
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
  processDate(scanner.Text())
 }
 if err := scanner.Err(); err != nil {
  fmt.Printf("Ошибка при чтении файла: %v\n", err)
 }
}

func main() {
 if len(os.Args) < 3 {
  displayHelp()
  return
 }

 switch os.Args[1] {
 case "-d":
  processDate(os.Args[2])
 case "-f", "-r":
  processFile(os.Args[2])
 case "-h":
  displayHelp()
 default:
  displayHelp()
 }
}