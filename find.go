	package main

	import (
		"flag"
		"fmt"
		"io/fs"
		"log"
		"path/filepath"
		"strconv"
		"strings"
	)

	// Функция для проверки, соответствует ли имя файла шаблону
	func matchesPattern(filename, pattern string) bool {

		if strings.Contains(pattern, "*") {
			return strings.HasSuffix(filename, strings.TrimPrefix(pattern, "*"))
		}

		return filename == pattern
	}

	// Функция для поиска файлов в директории
	func findFiles(directory, namePattern string, minSize int64) {

		err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {

			if err != nil {
				return err
			}

			// Пропускаем директории
			if d.IsDir() {
				return nil
			}

			// Получаем информацию о файле
			fileInfo, err := d.Info()

			if err != nil {
				return err
			}

			// Проверяем размер файла
			if fileInfo.Size() < minSize {
				return nil
			}

			// Проверяем соответствие имени файла шаблону
			if matchesPattern(fileInfo.Name(), namePattern) {
				fmt.Println(path) // Печатаем путь к найденному файлу
			}

			return nil
		})

		if err != nil {
			log.Fatalf("Ошибка при обходе директории: %v", err)
		}
	}

	func main() {
		// Определение аргументов командной строки
		flagHelp := flag.Bool("h", false, "Показать справку")
		flagDirectory := flag.String("d", "", "Путь к директории для поиска")
		flagNamePattern := flag.String("n", "", "Имя файла или шаблон для поиска")
		flagSize := flag.String("s", "0", "Минимальный размер файла в байтах (по умолчанию 0)")

		flag.Parse()

		// Проверка на наличие флага помощи
		if *flagHelp {
			printHelp()

			return
		}

		// Проверка корректности аргументов
		if *flagDirectory == "" {
			log.Fatal("Ошибка: Необходимо указать путь к директории с помощью -d.")
		}

		if *flagNamePattern == "" {
			log.Fatal("Ошибка: Необходимо указать имя файла или шаблон с помощью -n.")
		}

		minSize, err := strconv.ParseInt(*flagSize, 10, 64)

		if err != nil || minSize < 0 {
			log.Fatal("Ошибка: Размер должен быть неотрицательным числом.")
		}

		// Поиск файлов в указанной директории
		findFiles(*flagDirectory, *flagNamePattern, minSize)
	}

	// Функция для вывода справки по использованию программы
	func printHelp() {
		fmt.Println("Использование: find -d <директория> -n <имя> -s <размер>")
		fmt.Println("-h: Показать справку")
		fmt.Println("-d: Путь к директории для поиска")
		fmt.Println("-n: Имя файла или шаблон для поиска")
		fmt.Println("-s: Минимальный размер файла в байтах (по умолчанию 0)")
	}
