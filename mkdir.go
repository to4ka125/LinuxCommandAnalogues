package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// createDirectory создает директорию с указанным путем и режимом доступа.
// Если verbose установлен в true, выводит сообщение об успешном создании.
// Использует os.MkdirAll для создания всех необходимых родительских директорий, если флаг -p установлен.
func createDirectory(path string, mode int, verbose bool) error {

	err := os.MkdirAll(path, os.FileMode(mode)) // Создаем директорию и все родительские директории, если необходимо

	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Директория '%s' создана.\n", path)
	}

	return nil
}

func main() {

	// help флаг для вывода справки
	help := flag.Bool("h", false, "показать помощь")

	// createParents флаг для создания родительских директорий
	createParents := flag.Bool("p", false, "создать родительские директории при необходимости")

	// mode флаг для указания режима доступа к директории (octal)
	mode := flag.String("m", "0755", "режим доступа для создаваемых директорий (octal)")

	// verbose флаг для подробного вывода информации
	verbose := flag.Bool("v", false, "включить подробный вывод")

	flag.Parse() // Разбираем флаги

	if *help { // Если флаг -h установлен, выводим справку
		fmt.Println("Использование: mkdir [-h] [-p] [-m <режим>] [-v] <путь>")
		fmt.Println("-h: показать помощь")
		fmt.Println("-p: создавать родительские директории")
		fmt.Println("-m: режим доступа (octal, например 0755)")
		fmt.Println("-v: включить подробный вывод")

		os.Exit(0)
	}

	if flag.NArg() < 1 { // Проверяем наличие аргументов командной строки (путь к директории)

		fmt.Println("Ошибка: путь не указан. Укажите путь для создания директории.")

		os.Exit(1)
	}

	targetPath := flag.Arg(0) // Получаем путь к директории из аргументов

	modeInt, err := strconv.ParseInt(*mode, 8, 32) // Преобразуем строковое значение режима доступа в целое число

	if err != nil || modeInt < 0 || modeInt > 0777 {
		fmt.Println("Ошибка: Некорректный режим доступа. Используйте octal значение (например, 0755).")

		os.Exit(1)
	}

	// Создаем директорию, обрабатывая флаги -p и -v
	if *createParents {

		err := createDirectory(targetPath, int(modeInt), *verbose)

		if err != nil {
			fmt.Println("Ошибка при создании директории:", err)

			os.Exit(1)
		}

	} else {
		err := os.Mkdir(targetPath, os.FileMode(int(modeInt))) // Создаем директорию без создания родительских

		if err != nil {
			fmt.Println("Ошибка при создании директории:", err)

			os.Exit(1)
		}

		if *verbose {
			fmt.Printf("Директория '%s' создана.\n", targetPath)
		}

	}

	fmt.Println("Директория успешно создана:", targetPath)
}
