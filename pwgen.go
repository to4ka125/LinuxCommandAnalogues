package main

import (
 "crypto/rand"
 "fmt"
 "math/big"
 "os"
 "strconv"
)

const (
 letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
 numbers  = "0123456789"
 symbols  = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func generatePassword(length int, includeNumbers, includeSymbols, addSpecial bool) (string, error) {
 charset := letters
 if includeNumbers {
  charset += numbers
 }
 if includeSymbols {
  charset += symbols
 }

 password := make([]byte, length)
 for i := range password {
  index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
  if err != nil {
   return "", err
  }
  password[i] = charset[index.Int64()]
 }

 if addSpecial {
  specialIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
  if err != nil {
   return "", err
  }
  password = append(password, symbols[specialIndex.Int64()])
 }

 return string(password), nil
}

func displayHelp() {
 fmt.Println("Использование: pwgen [опции] [длина]")
 fmt.Println("Опции:")
 fmt.Println("  -n          Генерировать пароли с числами.")
 fmt.Println("  -s          Генерировать пароли с символами.")
 fmt.Println("  -y          Добавить один специальный символ в пароль.")
 fmt.Println("  -h          Показать справку.")
}

func main() {
 includeNumbers := false
 includeSymbols := false
 addSpecial := false
 length := 8

 if len(os.Args) < 2 {
  password, err := generatePassword(length, includeNumbers, includeSymbols, addSpecial)
  if err != nil {
   fmt.Println("Ошибка при генерации пароля:", err)
   return
  }
  fmt.Println("Сгенерированный пароль:", password)
  return
 }

 for i := 1; i < len(os.Args); i++ {
  switch os.Args[i] {
  case "-n":
   includeNumbers = true
  case "-s":
   includeSymbols = true
  case "-y":
   addSpecial = true
  case "-h":
   displayHelp()
   return
  default:
   if lengthArg, err := strconv.Atoi(os.Args[i]); err == nil {
    length = lengthArg
   } else {
    fmt.Println("Ошибка: неизвестная опция", os.Args[i])
    displayHelp()
    return
   }
  }
 }

 password, err := generatePassword(length, includeNumbers, includeSymbols, addSpecial)
 if err != nil {
  fmt.Println("Ошибка при генерации пароля:", err)
  return
 }
 fmt.Println("Сгенерированный пароль:", password)
}