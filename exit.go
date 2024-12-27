package main

import (
<<<<<<< HEAD
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
=======
 "flag"
 "fmt"
 "os"
 "strconv"
)

func showHelp() {
 fmt.Println("Использование: exit [опции] [код выхода]")
 fmt.Println("Опции:")
 fmt.Println("  -h          Показать справку.")
 fmt.Println("  -f          Принудительно завершить программу.")
 fmt.Println("  -n          Указать код выхода (по умолчанию 0).")
}

func main() {
 showHelpFlag := flag.Bool("h", false, "Показать справку")
 forceExit := flag.Bool("f", false, "Принудительно завершить программу")
 exitCode := flag.Int("n", 0, "Код выхода (по умолчанию 0)")

 flag.Parse()

 if *showHelpFlag {
  showHelp()
  return
 }

 if *forceExit {
  fmt.Println("Принудительное завершение программы...")
  os.Exit(*exitCode)
 } else {
  fmt.Printf("Завершение программы с кодом выхода: %d\n", *exitCode)
  os.Exit(*exitCode)
 }
}
>>>>>>> d60f4e32d4f5580607be7741b0039bb8841c14ab
