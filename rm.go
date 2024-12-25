package main

import (
 "fmt"
 "os"
)

func help() {
 fmt.Println("Справка: rm [опция] [путь]\nОпции:\n -R  Рекурсивно\n -f  Принудительно\n -h  Справка")
}

func remove(path string) error {
 return os.RemoveAll(path)
}

func main() {
 if len(os.Args) < 2 {
  fmt.Fprintln(os.Stderr, "rm: пропущен операнд\nИспользуйте «rm -h» для справки.")
  return
 }

 var force, recursive bool
 var files []string

 for _, arg := range os.Args[1:] {
  switch arg {
  case "-f":
   force = true
  case "-R":
   recursive = true
  case "-h":
   help()
   return
  default:
   files = append(files, arg)
  }
 }

 for _, path := range files {
  if _, err := os.Stat(path); os.IsNotExist(err) {
   fmt.Printf("Ошибка: %s не существует\n", path)
   continue
  }
  var err error
  if recursive {
   err = remove(path)
  } else {
   err = os.Remove(path)
  }
  if err != nil {
   if force {
    fmt.Printf("Ошибка при удалении %s: %v\n", path, err)
   } else {
    fmt.Printf("Ошибка: %s: %v. Используйте -f.\n", path, err)
   }
  } else {
   fmt.Printf("Удалено: %s\n", path)
  }
 }
}
