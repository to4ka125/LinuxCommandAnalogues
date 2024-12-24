package main

import(
	"fmt"
	"flag"
	"os"
)

func showHelp(){
	fmt.Println("Использование: cd [опции] <путь>")
	fmt.Println("-h: показать помощь")
	fmt.Println("-v: выводит текущую директорию перед выводом")
}
func changeDir(path string, verbose bool) error {
	if path == "" {
	 return fmt.Errorf("путь не может быть пустым")
	}
   
	if _, err := os.Stat(path); err != nil {
	 return fmt.Errorf("директория не найдена: %v", err)
	}
   
	if verbose {
	 if currentDir, err := os.Getwd(); err == nil {
	  fmt.Printf("Текущая директория: %s\n", currentDir)
	 }
	}
   
	if err := os.Chdir(path); err != nil {
	 return fmt.Errorf("не удалось перейти в директорию: %v", err)
	}
   
	fmt.Printf("Успешно перешли в директорию: %s\n", path)
	return nil
   }

func main (){
	verbose := flag.Bool("v", false, "выводить текущую директорию перед переходом")
	help := flag.Bool("h", false, "показать помощь")
	
	flag.Parse()
	
	if *help {
	 showHelp()
	 os.Exit(0)
	}
 
	args := flag.Args() 
 
	if len(args) == 0 {
	 fmt.Println("Ошибка: Необходимо указать путь")
	 os.Exit(1)
	}
 
	if err := changeDir(args[0], *verbose); err != nil {
	 fmt.Println("Ошибка:", err)
	 os.Exit(1)
	}
}