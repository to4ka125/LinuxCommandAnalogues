package main

import (
 "flag"
 "fmt"
 "io/fs"
 "os"
 "path/filepath"
 "sort"
 "strings"
 "time"
)

type FileDetails struct {
	Name    string
	Size    int64
	ModTime time.Time
	Mode    fs.FileMode
	IsDir   bool
   }
   
   func displayFileInfo(file FileDetails, showHidden bool, longFormat bool) {
	if !showHidden && file.Name[0] == '.' {
	 return
	}
   
	if longFormat {
	 fmt.Printf("%s %d %s %s\n", file.Name, file.Size, file.ModTime.Format(time.RFC3339), file.Mode)
	} else {
	 fmt.Println(file.Name)
	}
   }
   
   func listDirectory(dirPath string, showHidden bool, longFormat bool, reverseOrder bool, recursive bool) error {
	var files []FileDetails
   
	err := filepath.WalkDir(dirPath, func(path string, entry fs.DirEntry, err error) error {
	 if err != nil {
	  return err
	 }
   
	 if !showHidden && strings.HasPrefix(entry.Name(), ".") {
	  return nil
	 }
   
	 fileInfo, err := entry.Info()
	 if err != nil {
	  return err
	 }
   
	 files = append(files, FileDetails{
	  Name:    fileInfo.Name(),
	  Size:    fileInfo.Size(),
	  ModTime: fileInfo.ModTime(),
	  Mode:    fileInfo.Mode(),
	  IsDir:   fileInfo.IsDir(),
	 })
   
	 return nil
	})
   
	if err != nil {
	 return err
	}
   
	sort.Slice(files, func(i, j int) bool {
	 if reverseOrder {
	  i, j = j, i
	 }
	 return files[i].Name < files[j].Name
	})
   
	for _, file := range files {
	 displayFileInfo(file, showHidden, longFormat)
	}
   
	return nil
   }
   
   func main() {
	all := flag.Bool("a", false, "включить скрытые файлы")
	long := flag.Bool("l", false, "длинный формат")
	reverse := flag.Bool("r", false, "обратить порядок")
	recursive := flag.Bool("R", false, "рекурсивный вывод")
	help := flag.Bool("h", false, "показать помощь")
   
	flag.Parse()
   
	if *help {
	 fmt.Println("Использование: ls [опции] [путь]")
	 fmt.Println("-a: включить скрытые файлы")
	 fmt.Println("-l: длинный формат")
	 fmt.Println("-r: обратить порядок")
	 fmt.Println("-R: рекурсивный вывод")
	 fmt.Println("-h: помощь")
   
	 os.Exit(0)
	}
   
	dirPath := "."
   
	if len(flag.Args()) > 0 {
	 dirPath = flag.Args()[0]
	}
   
	err := listDirectory(dirPath, *all, *long, *reverse, *recursive)
   
	if err != nil {
	 fmt.Println("Ошибка:", err)
	}
   }