package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: не указана команда.")
		printHelp()
		return
	}

	if os.Args[1] == "--help" {
		printHelp()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "echo":
		echoCommand(args)
	case "ls":
		lsCommand(args)
	case "date":
		dateCommand(args)
	default:
		fmt.Printf("Ошибка: неизвестная команда '%s'\n", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Доступные команды:")
	fmt.Println("  echo <текст> - выводит текст на экран.")
	fmt.Println("  ls - выводит список файлов и папок в текущей директории.")
	fmt.Println("  date - выводит текущую дату и время.")
	fmt.Println("Используйте '--help' для отображения этой справки.")
}

func echoCommand(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func lsCommand(args []string) {
	cmd := exec.Command("ls", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}
	fmt.Print(string(output))
}

func dateCommand(args []string) {
	cmd := exec.Command("date", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
		return
	}
	fmt.Print(string(output))
}
