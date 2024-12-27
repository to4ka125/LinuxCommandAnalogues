package main

import (
 "fmt"
 "os"
 "strconv"
)

// showHelp выводит справочную информацию о программе
func showHelp() {
 fmt.Println("Использование: go run main.go -h <option>")
 fmt.Println("Ключи:")
 fmt.Println("  -h: Показать эту справку.")
 fmt.Println("  -c <code>: Завершить программу с заданным кодом возврата.")
 fmt.Println("  -f: Принудительно завершить программу.")
}

func main() {
 if len(os.Args) < 2 || os.Args[1] != "-h" {
  showHelp()
  return
 }

 switch os.Args[2] {
 case "-c":
  if len(os.Args) < 4 {
   fmt.Println("Ошибка: требуется код возврата.")
   return
  }
  code, err := strconv.Atoi(os.Args[3])
  if err != nil {
   fmt.Println("Ошибка: некорректный код возврата.")
   return
  }
  fmt.Printf("Завершение программы с кодом возврата: %d\n", code)
  os.Exit(code)
 case "-f":
  fmt.Println("Принудительное завершение программы.")
  panic("Принудительное завершение программы.") // Можно использовать panic для имитации аварийного завершения
 default:
  fmt.Println("Ошибка: неизвестный ключ. Используйте -h для получения справки.")
 }
}
