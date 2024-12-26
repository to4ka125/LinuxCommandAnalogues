package main

import (
 "fmt"
 "io/ioutil"
 "os"
 "os/user"
 "strconv"
 "strings"
)

func displayHelp() {
 fmt.Println("Использование: ps -h | -a | -u | -g")
 fmt.Println("Опции:")
 fmt.Println("-h  Выводит справку")
 fmt.Println("-a  Показать информацию о всех процессах.")
 fmt.Println("-u  Показать процессы текущего пользователя.")
 fmt.Println("-g  Показать все процессы, включая фоновые.")
 fmt.Println("Использование: ps [атрибут]")
}

type ProcInfo struct {
 User    string
 PID     int
 CPU     float64
 Mem     float64
 VSZ     int
 RSS     int
 TTY     string
 STAT    string
 Start   string
 Time    string
 Command string
}

func fetchProcesses() ([]ProcInfo, error) {
 var procList []ProcInfo

 files, err := ioutil.ReadDir("/proc")
 if err != nil {
  return nil, err
 }
 for _, file := range files {
  if file.IsDir() && isNumeric(file.Name()) {
   pid := file.Name()
   statPath := fmt.Sprintf("/proc/%s/stat", pid)
   statusPath := fmt.Sprintf("/proc/%s/status", pid)

   stat, err := ioutil.ReadFile(statPath)
   if err != nil {
    continue
   }
   status, err := ioutil.ReadFile(statusPath)
   if err != nil {
    continue
   }

   fields := strings.Fields(string(stat))
   if len(fields) > 1 {
    command := fields[1]
    pidInt, _ := strconv.Atoi(pid)

    tty := fields[6]  
    stat := fields[3]  
    time := fields[13] 

    userName := extractUserFromStatus(status)
    cpu := calculateCPUUsage(fields)
    mem := calculateMemoryUsage(fields)
    vsz := extractIntField(fields[22]) 
    rss := extractIntField(fields[23])

    procList = append(procList, ProcInfo{
     User:    userName,
     PID:     pidInt,
     CPU:     cpu,
     Mem:     mem,
     VSZ:     vsz,
     RSS:     rss,
     TTY:     tty,
     STAT:    stat,
     Start:   time,
     Time:    time,
     Command: command,
    })
   }
  }
 }
 return procList, nil
}

func isNumeric(s string) bool {
 _, err := strconv.Atoi(s)
 return err == nil
}

func extractUserFromStatus(status []byte) string {
 lines := strings.Split(string(status), "\n")
 for _, line := range lines {
  if strings.HasPrefix(line, "Uid:") {
   fields := strings.Fields(line)
   if len(fields) > 2 {
    userID := fields[1]
    uidInt, _ := strconv.Atoi(userID)
    userInfo, err := user.LookupId(strconv.Itoa(uidInt)) 
    if err == nil {
     return userInfo.Username
    }
   }
  }
 }
 return "unknown"
}

func calculateCPUUsage(fields []string) float64 {
 return 0.0 
}

func calculateMemoryUsage(fields []string) float64 {
 memUsage := extractIntField(fields[23]) 
 return float64(memUsage) / 1024.0      
}

func extractIntField(field string) int {
 value, _ := strconv.Atoi(field)
 return value
}

func main() {
 if len(os.Args) > 1 {
  switch os.Args[1] {
  case "-h":
   displayHelp() 
   return
  case "-a":
   procList, err := fetchProcesses()
   if err != nil {
    fmt.Println("Ошибка при получении процессов:", err)
    return
   }
   fmt.Println("PID\tTTY\tTIME\tCMD")
   for _, proc := range procList {
    fmt.Printf("%d\t%s\t%s\t%s\n", proc.PID, proc.TTY, proc.Time, proc.Command)
   }
   return
  case "-g":
   procList, err := fetchProcesses()
   if err != nil {
    fmt.Println("Ошибка при получении процессов:", err)
    return
   }
   fmt.Println("PID\tTTY\tSTAT\tTIME\tCOMMAND")
   for _, proc := range procList {
    fmt.Printf("%d\t%s\t%s\t%s\t%s\n", proc.PID, proc.TTY, proc.STAT, proc.Time, proc.Command)
   }
   return
  case "-u":
   procList, err := fetchProcesses()
   if err != nil {
    fmt.Println("Ошибка при получении процессов:", err)
    return
   }
   fmt.Println("USER\tPID\t%CPU\t%MEM\tVSZ\tRSS\tTTY\tSTAT\tSTART\tTIME\tCOMMAND")
   for _, proc := range procList {
    fmt.Printf("%s\t%d\t%.2f\t%.2f\t%d\t%d\t%s\t%s\t%s\t%s\t%s\n",
     proc.User, proc.PID,
13:48


proc.CPU, proc.Mem, proc.VSZ, proc.RSS, proc.TTY, proc.STAT, proc.Start, proc.Time, proc.Command)
   }
   return
  }
 }

 procList, err := fetchProcesses()
 if err != nil {
  fmt.Println("Ошибка при получении процессов:", err)
  return
 }
 fmt.Println("PID\tCommand")
 for _, proc := range procList {
  fmt.Printf("%d\t%s\n", proc.PID, proc.Command)
 }
}