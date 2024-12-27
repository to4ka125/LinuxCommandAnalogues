package main

import (
 "flag"
 "fmt"
 "os"
 "path/filepath"
 "strings"
)

func printHelp() {
 fmt.Println("Использование: find [-h] [-name <имя>] [-type <тип>] [-size <размер>] [-exec <команда>] <путь>")
 fmt.Println("Ключи:")
 fmt.Println("  -h           Показать эту справку")
 fmt.Println("  -name <имя> Найти файлы с указанным именем")
 fmt.Println("  -type <тип> Найти файлы указанного типа (f - файл, d - директория)")
 fmt.Println("  -size <размер> Найти файлы указанного размера (например, +100k, -1M)")
 fmt.Println("  -exec <команда> Выполнить команду для каждого найденного файла")
}

func main() {
 help := flag.Bool("h", false, "показать помощь")
 name := flag.String("name", "", "имя файла для поиска")
 fileType := flag.String("type", "", "тип файла (f - файл, d - директория)")
 size := flag.String("size", "", "размер файла (например, +100k, -1M)")
 exec := flag.String("exec", "", "команда для выполнения над найденными файлами")
 path := flag.String("path", ".", "путь для поиска")

 flag.Parse()

 if *help {
  printHelp()
  os.Exit(0)
 }

 if *name == "" && *fileType == "" && *size == "" && *exec == "" {
  fmt.Println("Ошибка: необходимо указать хотя бы один критерий поиска.")
  os.Exit(1)
 }

 err := filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
  if err != nil {
   return err
  }

  if *name != "" && !strings.Contains(info.Name(), *name) {
   return nil
  }

  if *fileType != "" {
   if (*fileType == "f" && info.IsDir()) || (*fileType == "d" && !info.IsDir()) {
    return nil
   }
  }

  if *size != "" {
   var sizeLimit int64
   var operator string
   fmt.Sscanf(*size, "%s%d", &operator, &sizeLimit)

   switch operator {
   case "+":
    if info.Size() <= sizeLimit {
     return nil
    }
   case "-":
    if info.Size() >= sizeLimit {
     return nil
    }
   default:
    if info.Size() != sizeLimit {
     return nil
    }
   }
  }

  fmt.Println(path)

  if *exec != "" {
   cmd := exec.Command(*exec, path)
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   if err := cmd.Run(); err != nil {
    fmt.Printf("Ошибка выполнения команды: %v\n", err)
   }
  }

  return nil
 })

 if err != nil {
  fmt.Printf("Ошибка при обходе файловой системы: %v\n", err)
  os.Exit(1)
 }
}