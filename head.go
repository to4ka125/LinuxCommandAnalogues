package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
)

// showHelp выводит справочную информацию о программе
func showHelp() {
 fmt.Println("Использование: head [-c N | -n N] [-q] <file1> [<file2> ...]")
 fmt.Println("Выводит первые N байт или строки из файла.")
 fmt.Println("  -c N: Выводит первые N байт файла.")
 fmt.Println("  -n N: Выводит первые N строк файла.")
 fmt.Println("  -q: Не выводить заголовки файлов.")
 fmt.Println("  -h: Показать эту справку.")
}

// readFile выводит первые N байт или строк из файла
func readFile(path string, bytes int, lines int, quiet bool) error {
 file, err := os.Open(path)
 if err != nil {
  return fmt.Errorf("ошибка при открытии файла %s: %v", path, err)
 }
 defer file.Close()

 if lines > 0 {
  if !quiet {
   fmt.Printf("==> %s <==\n", path)
  }
  scanner := bufio.NewScanner(file)
  for i := 0; i < lines && scanner.Scan(); i++ {
   fmt.Println(scanner.Text())
  }
 } else if bytes > 0 {
  if !quiet {
   fmt.Printf("==> %s <==\n", path)
  }
  buffer := make([]byte, bytes)
  n, err := file.Read(buffer)
  if err != nil && err.Error() != "EOF" {
   return fmt.Errorf("ошибка при чтении файла %s: %v", path, err)
  }
  fmt.Print(string(buffer[:n]))
 }
 return nil
}

// main выполняет логику программы
func main() {
 if len(os.Args) < 3 {
  showHelp()
  return
 }

 var bytes, lines int
 var quiet bool
 var paths []string

 for i := 1; i < len(os.Args); i++ {
  arg := os.Args[i]
  switch arg {
  case "-h":
   showHelp()
   return
  case "-q":
   quiet = true
  case "-c":
   if i+1 < len(os.Args) {
    if count, err := strconv.Atoi(os.Args[i+1]); err == nil && count >= 0 {
     bytes = count
     i++
    } else {
     fmt.Println("Ошибка: некорректное значение для -c")
     return
    }
   } else {
    fmt.Println("Ошибка: требуется значение для -c")
    return
   }
  case "-n":
   if i+1 < len(os.Args) {
    if count, err := strconv.Atoi(os.Args[i+1]); err == nil && count >= 0 {
     lines = count
     i++
    } else {
     fmt.Println("Ошибка: некорректное значение для -n")
     return
    }
   } else {
    fmt.Println("Ошибка: требуется значение для -n")
    return
   }
  default:
   paths = append(paths, arg)
  }
 }

 for _, path := range paths {
  if err := readFile(path, bytes, lines, quiet); err != nil {
   fmt.Println(err)
  }
 }
}
