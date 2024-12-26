package main

import (
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func usage() {
    fmt.Println("Использование: cp [опции] источник назначение")
    fmt.Println("Опции:")
    fmt.Println("  -r    Рекурсивное копирование каталогов")
    fmt.Println("  -i    Запрашивать подтверждение перед перезаписью файла")
    fmt.Println("  -v    Выводить информацию о процессе копирования")
    fmt.Println("  -h    Выводить справку по использованию команды")
}

func copyFile(src, dst string, interactive, verbose bool) error {
    if verbose {
        fmt.Printf("Копирую файл %s в %s\n", src, dst)
    }
    
    sourceFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("не удалось открыть источник: %w", err)
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("не удалось создать файл назначения: %w", err)
    }
    defer destFile.Close()

    if _, err = io.Copy(destFile, sourceFile); err != nil {
        return fmt.Errorf("ошибка при копировании: %w", err)
    }

    return nil
}

func copyDir(srcDir, dstDir string, interactive, verbose bool) error {
    if verbose {
        fmt.Printf("Копирую каталог %s в %s\n", srcDir, dstDir)
    }

    return filepath.Walk(srcDir, func(src string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        relPath, _ := filepath.Rel(srcDir, src)
        dst := filepath.Join(dstDir, relPath)

        if info.IsDir() {
            return os.MkdirAll(dst, os.ModePerm)
        }

        if interactive {
            if _, err := os.Stat(dst); err == nil {
                var response string
                fmt.Printf("Файл %s уже существует. Перезаписать? (y/n): ", dst)
                fmt.Scanln(&response)
                if response != "y" {
                    return nil
                }
            }
        }

        return copyFile(src, dst, interactive, verbose)
    })
}

func main() {
    recursive := flag.Bool("r", false, "Рекурсивное копирование каталогов")
    interactive := flag.Bool("i", false, "Запрашивать подтверждение перед перезаписью файла")
    verbose := flag.Bool("v", false, "Выводить информацию о процессе копирования")
    help := flag.Bool("h", false, "Выводить справку по использованию команды")

    flag.Parse()

    if *help {
        usage()
        return
    }

    args := flag.Args()
    if len(args) < 2 {
        fmt.Println("Ошибка: необходимо указать источник и назначение.")
        usage()
        return
    }

    src := args[0]
    dst := args[1]

    srcInfo, err := os.Stat(src)
    if err != nil {
        fmt.Printf("Ошибка: не удалось получить информацию о источнике: %s\n", err)
        return
    }

    if srcInfo.IsDir() {
        if !*recursive {
            fmt.Println("Ошибка: необходимо использовать -r для копирования каталогов.")
            return
        }
        if err := copyDir(src, dst, *interactive, *verbose); err != nil {
            fmt.Printf("Ошибка при копировании каталога: %s\n", err)
        }
    } else {
        if *interactive {
            if _, err := os.Stat(dst); err == nil {
                var response string
                fmt.Printf("Файл %s уже существует. Перезаписать? (y/n): ", dst)
                fmt.Scanln(&response)
                if response != "y" {
                    return
                }
            }
        }
        if err := copyFile(src, dst, *interactive, *verbose); err != nil {
            fmt.Printf("Ошибка при копировании файла: %s\n", err)
        }
    }
}
