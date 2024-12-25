package main

import (
 "fmt"
 "os"
 "time"
 "strings"
)

const ver = "1.0.0"

func help() {
 fmt.Println("Справка")
 fmt.Println("Использование: touch [опции] [файлы]")
 fmt.Println("Опции:")
 fmt.Println(" -h        Показать помощь.")
 fmt.Println(" -a        Изменяет только время доступа к файлу.")
 fmt.Println(" -m        Изменяет только время модификации к файлу.")
 fmt.Println(" --version Отображает текущую версию утилиты.")
}

func updateTS(fn string, at, mt time.Time) error {
 return os.Chtimes(fn, at, mt)
}

func main() {
 if len(os.Args) < 2 {
  fmt.Fprintln(os.Stderr, "touch: пропущен операнд, задающий файл")
  fmt.Fprintln(os.Stderr, "По команде «touch --help» можно получить дополнительную информацию.")
  return
 }

 var accessOnly, modifyOnly bool

 for _, arg := range os.Args[1:] {
  switch arg {
  case "-h":
   help()
   return
  case "-a":
   accessOnly = true
  case "-m":
   modifyOnly = true
  case "--version":
   fmt.Printf("touch (GNU coreutils) %s\nCopyright (C) 2022 Free Software Foundation, Inc.\nЛицензия GPLv3+: GNU GPL версии 3 или новее <https://gnu.org/licenses/gpl.html>\nЭто свободное ПО: вы можете изменять и распространять его.\nНет НИКАКИХ ГАРАНТИЙ в пределах действующего законодательства.\n\nАвторы программы — Paul Rubin, Arnold Robbins, Jim Kingdon,\nDavid MacKenzie и Randy Smith..\n", ver)
   return
  default:
   fn := arg

   file, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR, 0644)
   if err != nil {
    fmt.Fprintf(os.Stderr, "Ошибка при работе с файлом '%s': %v\n", fn, err)
    continue
   }
   file.Close()

   now := time.Now()
   at, mt := now, now

   if accessOnly && !modifyOnly {
    mt = now
   } else if modifyOnly && !accessOnly {
    at = now
   }

   err = updateTS(fn, at, mt)
   if err != nil {
    fmt.Fprintf(os.Stderr, "Ошибка при обновлении временной метки файла '%s': %v\n", fn, err)
   } else {
    fmt.Printf("Временные метки файла '%s' успешно обновлены.\n", fn)
   }
  }
 }

 if len(os.Args) == 2 && (strings.HasPrefix(os.Args[1], "-") || os.Args[1] == "--version") {
  fmt.Fprintln(os.Stderr, "touch: пропущен операнд, задающий файл")
  fmt.Fprintln(os.Stderr, "По команде «touch -h» можно получить дополнительную информацию.")
 }
}
