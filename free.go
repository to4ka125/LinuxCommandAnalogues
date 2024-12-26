package main

import (
 "fmt"
 "os"
 "syscall"
)

const (
 KB = 1024
 MB = 1024 * KB
 GB = 1024 * MB
)

func showHelp() {
 fmt.Println("Справка")
 fmt.Println("Использование: free [опции]")
 fmt.Println("Опции:")
 fmt.Println("  -b     Отображает свободную память в байтах.")
 fmt.Println("  -m     Отображает свободную память в мегабайтах.")
 fmt.Println("  -g     Отображает свободную память в гигабайтах.")
 fmt.Println("  -h     Выводит справку.")
}

func getMemoryStats() (total, used, free uint64) {
 var sysInfo syscall.Sysinfo_t
 syscall.Sysinfo(&sysInfo)

 total = sysInfo.Totalram * uint64(sysInfo.Unit)
 free = sysInfo.Freeram * uint64(sysInfo.Unit)
 used = total - free
 return
}

func displayMemory(total, used, free uint64, unit string) {
 var format string
 switch unit {
 case "b":
  format = "%d B"
 case "m":
  format = "%d MB"
 case "g":
  format = "%d GB"
 default:
  if total >= GB {
   format = "%d GB"
  } else {
   format = "%d MB"
  }
 }
 fmt.Printf("033[33mMem: "+format+" total, "+format+" used, "+format+" free033[0m\n",
  total/parseUnit(unit), used/parseUnit(unit), free/parseUnit(unit))
}

func parseUnit(unit string) uint64 {
 switch unit {
 case "b":
  return 1
 case "m":
  return MB
 case "g":
  return GB
 default:
  return MB // по умолчанию показываем в мегабайтах
 }
}

func main() {
 var unit string

 if len(os.Args) > 2 {
  showHelp()
  return
 }

 if len(os.Args) == 2 {
  switch os.Args[1] {
  case "-b", "-m", "-g":
   unit = os.Args[1][1:] // извлекаем символ после "-"
  case "-h":
   showHelp()
   return
  default:
   fmt.Fprintln(os.Stderr, "Неизвестная опция:", os.Args[1])
   showHelp()
   return
  }
 }

 total, used, free := getMemoryStats()
 displayMemory(total, used, free, unit)
}
