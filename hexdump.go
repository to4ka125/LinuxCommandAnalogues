package main

import (
 "flag"
 "fmt"
 "io/ioutil"
 "log"
)

// Функция для печати содержимого в шестнадцатеричном формате
func printHexDump(data []byte, totalBytes int) {
 lines := (totalBytes + 15) / 16 // Количество строк для вывода

 for i := 0; i < lines; i++ {
  start := i * 16
  end := start + 16
  if end > totalBytes {
   end = totalBytes
  }

  // Печатаем адрес строки
  fmt.Printf("%08x  ", start)

  // Печатаем байты в шестнадцатеричном формате
  for j := start; j < end; j++ {
   if j < len(data) {
    fmt.Printf("%02x ", data[j])
   } else {
    fmt.Print("   ") // Заполняем пробелами, если байтов не хватает
   }
  }

  // Печатаем символы, соответствующие байтам
  fmt.Print(" |")
  for j := start; j < end; j++ {
   if j < len(data) {
    if data[j] >= 32 && data[j] <= 126 {
     fmt.Printf("%c", data[j]) // Печатаемые символы
    } else {
     fmt.Print(".") // Непечатаемые символы
    }
   }
  }
  fmt.Println("|")
 }
}

func main() {
 // Определение аргументов командной строки
 showHelp := flag.Bool("h", false, "Показать справку")
 filePath := flag.String("f", "", "Путь к файлу для чтения")
 byteCount := flag.Int("n", -1, "Количество байтов для чтения (по умолчанию - весь файл)")

 flag.Parse()

 // Проверка на наличие флага помощи
 if *showHelp {
  fmt.Println("Использование: hexdump -f <файл> -n <количество>")
  fmt.Println("-h: Показать справку")
  fmt.Println("-f: Путь к файлу для чтения")
  fmt.Println("-n: Количество байтов для чтения (по умолчанию - весь файл)")
  return
 }

 // Проверка корректности аргументов
 if *filePath == "" {
  log.Fatal("Ошибка: Необходимо указать путь к файлу с помощью -f.")
 }

 if *byteCount < -1 {
  log.Fatal("Ошибка: Количество байтов должно быть неотрицательным числом или -1 для чтения всего файла.")
 }

 // Чтение файла
 data, err := ioutil.ReadFile(*filePath)
 if err != nil {
  log.Fatalf("Ошибка при чтении файла: %v", err)
 }

 numBytes := len(data)
 if *byteCount != -1 && *byteCount < numBytes {
  numBytes = *byteCount
 }

 printHexDump(data[:numBytes], numBytes)
}
