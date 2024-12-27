package main

import (
 "flag"
 "fmt"
 "os"
 "strings"
 "syscall"
)

func showHelp() {
 fmt.Println("Использование: df -a -B <размер> --direct")
 fmt.Println("  -a        Показать все файловые системы.")
 fmt.Println("  -B <размер>  Указать размер в байтах (например, 1K, 1M).")
 fmt.Println("  --direct  Игнорировать кэширование и показывать прямую информацию.")
 fmt.Println("  -h        Показать справку.")
}

func getFSInfo(direct bool) ([]syscall.Statfs_t, error) {
 var fsStats []syscall.Statfs_t
 data, err := os.ReadFile("/proc/mounts")
 if err != nil {
  return nil, err
 }
 lines := strings.Split(string(data), "\n")
 for _, line := range lines {
  if line == "" {
   continue
  }
  parts := strings.Fields(line)
  if len(parts) < 2 {
   continue
  }
  path := parts[1]
  var stat syscall.Statfs_t
  if err := syscall.Statfs(path, &stat); err == nil {
   fsStats = append(fsStats, stat)
  }
 }
 return fsStats, nil
}

func main() {
 all := flag.Bool("a", false, "Показать все файловые системы")
 size := flag.String("B", "", "Указать размер в байтах (например, 1K, 1M)")
 direct := flag.Bool("direct", false, "Игнорировать кэширование и показывать прямую информацию")
 help := flag.Bool("h", false, "Показать справку")

 flag.Parse()

 if *help {
  showHelp()
  return
 }

 if *size == "" {
  fmt.Println("Ошибка: необходимо указать размер с помощью флага -B.")
  showHelp()
  return
 }

 stats, err := getFSInfo(*direct)
 if err != nil {
  fmt.Printf("Ошибка получения информации о файловых системах: %s\n", err)
  return
 }

 fmt.Printf("%-20s %-10s %-10s %-10s\n", "Файловая система", "Размер", "Использовано", "Доступно")

 for _, stat := range stats {
  total := stat.Blocks * uint64(stat.Bsize)
  free := stat.Bfree * uint64(stat.Bsize)
  if *all || total > 0 {
   fmt.Printf("%-20s %-10d %-10d %-10d\n", "/proc/mounts", total, total-free, free)
  }
 }

 if *direct {
  fmt.Println("Флаг --direct активирован. Данные получены напрямую через syscall.Statfs.")
 }
}
