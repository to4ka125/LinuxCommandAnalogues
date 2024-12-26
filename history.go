package main

import (
 "bufio"
 "fmt"
 "os"
 "path/filepath"
 "strings"
)

// help выводит справку по использованию команды.
func help() {
 fmt.Println("Справка")
 fmt.Println("Использование: history [опция]")
 fmt.Println("Опции:")
 fmt.Println("  -c              Очистить историю.")
 fmt.Println("  -d [n]          Удалить строку с номером n из истории.")
 fmt.Println("  -a [команда]    Добавить новую команду в историю.")
 fmt.Println("  -h              Показать эту справку.")
}

// historyFile указывает путь к файлу с историей команд.
var historyFile = filepath.Join(getHomeDir(), ".bash_history")

// cmdHistory хранит историю команд в памяти.
var cmdHistory []string

// getHomeDir возвращает путь к домашнему каталогу пользователя.
func getHomeDir() string {
 homeDir, err := os.UserHomeDir()
 if err != nil {
  fmt.Println("Ошибка при получении домашнего каталога:", err)
  os.Exit(1)
 }
 return homeDir
}

// loadHistory загружает историю команд из файла.
func loadHistory() error {
 file, err := os.OpenFile(historyFile, os.O_RDONLY, 0644)
 if err != nil {
  if os.IsNotExist(err) {
   return nil
  }
  return err
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
  cmdHistory = append(cmdHistory, scanner.Text())
 }
 return scanner.Err()
}

// saveHistory сохраняет историю команд в файл.
func saveHistory() error {
 file, err := os.Create(historyFile)
 if err != nil {
  return err
 }
 defer file.Close()

 for _, cmd := range cmdHistory {
  _, err := file.WriteString(cmd + "\n")
  if err != nil {
   return err
  }
 }
 return nil
}

// clearHistory очищает историю.
func clearHistory() {
 cmdHistory = []string{}
 fmt.Println("История очищена.")
 saveHistory() // Сохраняем изменения в файле
}

// modifyHistory изменяет историю: удаляет строку или добавляет новую команду.
func modifyHistory(action string, lineNum int, command string) {
 if action == "delete" {
  if lineNum < 1 || lineNum > len(cmdHistory) {
   fmt.Printf("Ошибка: номер строки %d вне диапазона\n", lineNum)
   return
  }
  cmdHistory = append(cmdHistory[:lineNum-1], cmdHistory[lineNum:]...) // Удаляем строку.
  fmt.Printf("Строка %d удалена из истории.\n", lineNum)
 } else if action == "add" {
  cmdHistory = append(cmdHistory, command) // Добавляем новую команду.
  fmt.Printf("Команда '%s' добавлена в историю.\n", command)
 }

 saveHistory() // Сохраняем изменения в файле
}

// printHistory выводит всю историю команд на экран.
func printHistory() {
 if len(cmdHistory) == 0 {
  fmt.Println("История пуста.")
  return
 }
 for i, cmd := range cmdHistory {
  fmt.Printf("%d  %s\n", i+1, cmd) // Формат вывода с номером строки.
 }
}

// main функция
func main() {
 if err := loadHistory(); err != nil {
  fmt.Println("Ошибка при загрузке истории:", err)
  return
 }

 if len(os.Args) == 1 {
  printHistory()
  return
 }

 switch os.Args[1] {
 case "-h":
  help()
 case "-c":
  clearHistory()
 case "-d":
  if len(os.Args) < 3 {
   fmt.Println("Ошибка: необходимо указать номер строки для удаления.")
   return
  }
  var lineNum int
  fmt.Sscanf(os.Args[2], "%d", &lineNum)
  modifyHistory("delete", lineNum, "")
 case "-a":
  if len(os.Args) < 3 {
   fmt.Println("Ошибка: необходимо указать команду для добавления в историю.")
   return
  }
  command := strings.Join(os.Args[2:], " ")
  modifyHistory("add", 0, command)
 default:
  fmt.Println("Неизвестная опция:", os.Args[1])
 }
}