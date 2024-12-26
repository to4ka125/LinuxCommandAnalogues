package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyFile(src, dst string, preserveAttrs bool) error {
	sourceFile, err := os.Open(src)

	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)

	if err != nil {
		return err
	}

	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile) // Копируем содержимое файла

	if err != nil {
		return err
	}

	if preserveAttrs {
		sourceInfo, err := sourceFile.Stat()

		if err != nil {
			return err
		}

		err = os.Chmod(dst, sourceInfo.Mode()) // Сохраняем права доступа

		if err != nil {
			return err
		}
	}

	return nil
}

func copyDirectory(src string, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm) // Создаем директорию назначения

	if err != nil {
		return err
	}

	files, err := os.ReadDir(src) // Читаем содержимое директории

	if err != nil {
		return err
	}

	for _, file := range files {

		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dst, file.Name())

		if file.IsDir() {
			err = copyDirectory(srcPath, dstPath) // Рекурсивно копируем поддиректории

			if err != nil {
				return err
			}

		} else {
			err = copyFile(srcPath, dstPath, false) // Копируем файлы (без сохранения атрибутов)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func handleArgs(args []string) (string, string, string, error) {

	if len(args) < 4 {
		return "", "", "", fmt.Errorf("недостаточно аргументов")
	}

	flag := args[1]

	src := args[2]

	dst := args[3]

	if flag != "-a" && flag != "-b" && flag != "-d" {
		return "", "", "", fmt.Errorf("некорректный флаг: %s", flag)
	}
	return flag, src, dst, nil
}

func main() {
	flag, src, dst, err := handleArgs(os.Args)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Использование: cp [-a | -b | -d] <source> <destination>")

		return
	}

	switch flag {

	case "-a":
		err = copyFile(src, dst, true) // Копирование с сохранением атрибутов

	case "-b":
		err = copyFile(src, dst, false) // Бинарное копирование (без сохранения атрибутов)

	case "-d":
		err = copyDirectory(src, dst) // Копирование директории

	default:
		err = fmt.Errorf("недопустимый флаг: %s", flag)

	}

	if err != nil {
		fmt.Println("Ошибка:", err)

		return
	}
	fmt.Println("Копирование завершено успешно.")
}
