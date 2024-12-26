package main

import (
 "bufio"
 "fmt"
 "os"
 "os/exec"
 "strconv"
)

// Читает историю команд из файла
func readHistory(path string) ([]string, error) {
 file, err := os.Open(path)
 if err != nil {
  return nil, err
 }
 defer file.Close()

 var commands []string
 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
  commands = append(commands, scanner.Text())
 }

 if err := scanner.Err(); err != nil {
  return nil, err
 }

 return commands, nil
}

// Выполняет команду
func execute(command string) error {
 cmd := exec.Command("bash", "-c", command)
 cmd.Stdout = os.Stdout
 cmd.Stderr = os.Stderr
 return cmd.Run()
}

// Выводит справку
func showHelp() {
 fmt.Println("Использование: repeat_command -h | -n <номер>")
 fmt.Println("-h: Показать справку")
 fmt.Println("-n <номер>: Номер команды в истории для выполнения")
}

// Основная функция
func main() {
 if len(os.Args) != 3 {
  showHelp()
  os.Exit(1)
 }

 var cmdIndex int

 for i := 1; i < len(os.Args); i++ {
  switch os.Args[i] {
  case "-h":
   showHelp()
   os.Exit(0)
  case "-n":
   i++
   if i < len(os.Args) {
    var err error
    cmdIndex, err = strconv.Atoi(os.Args[i])
    if err != nil {
     fmt.Println("Ошибка: некорректный номер команды.")
     os.Exit(1)
    }
   }
  default:
   fmt.Println("Неверный аргумент.")
   os.Exit(1)
  }
 }

 historyFile := os.Getenv("HOME") + "/.bash_history" // Путь к файлу истории Bash
 history, err := readHistory(historyFile)
 if err != nil {
  fmt.Printf("Ошибка при чтении истории: %v\n", err)
  os.Exit(1)
 }

 if cmdIndex <= 0 || cmdIndex > len(history) {
  fmt.Printf("Команда с номером %d не найдена в истории.\n", cmdIndex)
  os.Exit(1)
 }

 commandToRun := history[cmdIndex-1] // Получаем команду по номеру

 fmt.Printf("Выполнение команды: %s\n", commandToRun)

 if err := execute(commandToRun); err != nil {
  fmt.Printf("Ошибка при выполнении команды: %v\n", err)
  os.Exit(1)
 }

 fmt.Println("Команда успешно выполнена.")
}
