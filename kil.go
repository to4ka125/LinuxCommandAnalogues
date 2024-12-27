package main

import (
 "fmt"
 "os"
 "strconv"
 "syscall"
)

func displayHelp() {
 fmt.Println("Использование: kill <ID> [<сигнал>]")
 fmt.Println("Параметры:")
 fmt.Println("  <ID>      Указать ID процесса для завершения.")
 fmt.Println("  <сигнал>  Указать сигнал для отправки (по умолчанию SIGTERM).")
 fmt.Println("  -h        Показать справку.")
}

func isValidProcessID(id int) bool {
 return id > 0
}

func resolveSignal(sig string) (syscall.Signal, bool) {
 switch sig {
 case "SIGTERM":
  return syscall.SIGTERM, true
 case "SIGKILL":
  return syscall.SIGKILL, true
 default:
  return 0, false
 }
}

func main() {
 if len(os.Args) < 2 {
  displayHelp()
  os.Exit(1)
 }

 var procID int
 var procSignal syscall.Signal = syscall.SIGTERM // По умолчанию SIGTERM

 // Получаем ID процесса из аргументов
 var err error
 procID, err = strconv.Atoi(os.Args[1])
 if err != nil || !isValidProcessID(procID) {
  fmt.Println("Ошибка: Неверный ID процесса. Он должен быть положительным числом.")
  displayHelp()
  os.Exit(1)
 }

 // Проверяем наличие второго аргумента для сигнала
 if len(os.Args) == 3 {
  signalStr := os.Args[2]
  var valid bool
  procSignal, valid = resolveSignal(signalStr)

  if !valid {
   fmt.Println("Ошибка: Неверный сигнал. Допустимые сигналы: SIGTERM, SIGKILL.")
   displayHelp()
   os.Exit(1)
  }
 }

 err = syscall.Kill(procID, procSignal)

 if err != nil {
  fmt.Printf("Ошибка при отправке сигнала %s процессу с ID %d: %v\n", procSignal, procID, err)
  os.Exit(1)
 }
 fmt.Printf("Сигнал %s успешно отправлен процессу с ID %d\n", procSignal, procID)
}