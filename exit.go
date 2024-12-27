package main

import (
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
