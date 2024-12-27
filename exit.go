package main

import (
 "fmt"
 "os"
 "sync"
 "time"
)

func main() {
 var wg sync.WaitGroup

 wg.Add(1) // Увеличиваем счетчик на 1

 go func() {
  defer wg.Done() // Уменьшаем счетчик на 1 при завершении
  time.Sleep(2 * time.Second) // Имитация работы
  fmt.Println("Горутина завершена.")
 }()

 // Выводим сообщение перед выходом
 fmt.Println("Ожидание завершения горутин...")
 wg.Wait() // Ожидаем завершения всех горутин

 // Указываем код выхода (0 - успешный выход)
 exitCode := 0

 // Завершаем программу с указанным кодом выхода
 os.Exit(exitCode)
}