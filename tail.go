package main

import (
 "bufio"
 "fmt"
 "io"
 "os"
 "strconv"
 "time"
)

// showHelp выводит справочную информацию о программе
func showHelp() {
 fmt.Println("Использование: tail [-n N | -c N] [-f] <file>")
 fmt.Println("Выводит последние N байт или строки из файла.")
 fmt.Println("  -n N: Выводит последние N строк файла.")
 fmt.Println("  -c N: Выводит последние N байт файла.")
 fmt.Println("  -f: Постоянно отслеживает файл на предмет новых данных.")
 fmt.Println("  -h: Показать эту справку.")
}

// readLastLines читает последние N строк из файла
func readLastLines(filePath string, lineCount int) error {
 file, err := os.Open(filePath)
 if err != nil {
  return fmt.Errorf("ошибка при открытии файла %s: %v", filePath, err)
 }
 defer file.Close()

 lines := make([]string, 0, lineCount)
 scanner := bufio.NewScanner(file)

 for scanner.Scan() {
  lines = append(lines, scanner.Text())
  if len(lines) > lineCount {
   lines = lines[1:] // Удаляем первую строку, если превышен лимит
  }
 }

 if err := scanner.Err(); err != nil {
  return fmt.Errorf("ошибка при чтении файла %s: %v", filePath, err)
 }

 for _, line := range lines {
  fmt.Println(line)
 }

 return nil
}

// readLastBytes читает последние N байт из файла
func readLastBytes(filePath string, byteCount int) error {
 file, err := os.Open(filePath)
 if err != nil {
  return fmt.Errorf("ошибка при открытии файла %s: %v", filePath, err)
 }
 defer file.Close()

 stat, err := file.Stat()
 if err != nil {
  return fmt.Errorf("ошибка при получении информации о файле %s: %v", filePath, err)
 }

 start := int64(0)
 if stat.Size() > int64(byteCount) {
  start = stat.Size() - int64(byteCount)
 }

 if _, err := file.Seek(start, io.SeekStart); err != nil {
  return fmt.Errorf("ошибка при перемещении по файлу %s: %v", filePath, err)
 }

 buffer := make([]byte, byteCount)
 n, err := file.Read(buffer)
 if err != nil && err != io.EOF {
  return fmt.Errorf("ошибка при чтении файла %s: %v", filePath, err)
 }

 fmt.Print(string(buffer[:n]))
 return nil
}

// followFile отслеживает файл на предмет новых данных
func followFile(filePath string) error {
 file, err := os.Open(filePath)
 if err != nil {
  return fmt.Errorf("ошибка при открытии файла %s: %v", filePath, err)
 }
 defer file.Close()

 if _, err := file.Seek(0, io.SeekEnd); err != nil {
  return fmt.Errorf("ошибка при перемещении по файлу %s: %v", filePath, err)
 }

 for {
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
   fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
   return fmt.Errorf("ошибка при чтении файла %s: %v", filePath, err)
  }

  time.Sleep(1 * time.Second) // Задержка перед следующей проверкой
 }
}

// main выполняет логику программы
func main() {
 if len(os.Args) < 3 {
  showHelp()
  return
 }

 var byteCount int
 var lineCount int
 var follow bool
 var filePath string

 for i := 1; i < len(os.Args); i++ {
  arg := os.Args[i]

  switch arg {
  case "-h":
   showHelp()
   return
  case "-f":
   follow = true
  case "-c":
   if i+1 < len(os.Args) {
    count, err := strconv.Atoi(os.Args[i+1])
    if err != nil || count < 0 {
     fmt.Println("Ошибка: некорректное значение для -c")
     return
    }
    byteCount = count
    i++ // Пропускаем следующий аргумент
   } else {
    fmt.Println("Ошибка: требуется значение для -c")
    return
   }
  case "-n":
   if i+1 < len(os.Args) {
    count, err := strconv.Atoi(os.Args[i+1])
    if err != nil || count < 0 {
     fmt.Println("Ошибка: некорректное значение для -n")
     return
    }
    lineCount = count
    i++ // Пропускаем следующий аргумент
   } else {
    fmt.Println("Ошибка: требуется значение для -n")
    return
   }
  default:
   filePath = arg
  }
 }

 if follow {
  err := followFile(filePath)
  if err != nil {
   fmt.Println(err)
  }
 } else if lineCount > 0 {
  err := readLastLines(filePath, lineCount)
  if err != nil {
   fmt.Println(err)
  }
 } else if byteCount > 0 {
	err := readLastBytes(filePath, byteCount)
	if err != nil {
	 fmt.Println(err)
	}
   } else {
	fmt.Println("Ошибка: требуется указать количество строк или байт для вывода.")
	showHelp()
   }
  }
  