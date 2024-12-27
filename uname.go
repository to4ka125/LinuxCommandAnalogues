package main

import (
 "fmt"
 "os"
 "runtime"
)

type SystemInfo struct {
 Kernel string
 Host   string
 OS     string
}

func NewSystemInfo() *SystemInfo {
 return &SystemInfo{
  Kernel: runtime.GOOS,
  Host:   "localhost",
  OS:     runtime.GOARCH,
 }
}

func (si *SystemInfo) Show() {
 fmt.Printf("Имя ядра: %s\nИмя узла: %s\nОперационная система: %s\n", si.Kernel, si.Host, si.OS)
}

func PrintHelp() {
 fmt.Println("Использование:")
 fmt.Println("  uname -a                   Показать всю информацию о системе")
 fmt.Println("  uname -s                   Показать имя ядра")
 fmt.Println("  uname -n                   Показать имя узла")
 fmt.Println("  uname -h                   Показать справку")
}

func ValidateArgs(args []string) error {
 if len(args) != 2 {
  return fmt.Errorf("недостаточно или слишком много аргументов: ожидается ровно 1 аргумент")
 }
 switch args[1] {
 case "-a", "-s", "-n", "-h":
  return nil
 default:
  return fmt.Errorf("неизвестный аргумент: %s. Допустимые аргументы: -a, -s, -n", args[1])
 }
}

func main() {
 args := os.Args
 if err := ValidateArgs(args); err != nil {
  PrintHelp()
  fmt.Println(err)
  return
 }
 sysInfo := NewSystemInfo()
 switch args[1] {
 case "-a":
  sysInfo.Show()
 case "-s":
  fmt.Println(sysInfo.Kernel)
 case "-n":
  fmt.Println(sysInfo.Host)
 case "-h":
  PrintHelp()
 }
}
