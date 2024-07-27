package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GeneratePassword() string {
	numbers := GenerateInt()
	strings := GenerateString(15)

	myRune := []rune(numbers + strings) //здесь храним то что мы генерировали ранее

	lenRune := len(myRune) - 1

	var result string

	for i := 0; i < 20; i++ {
		result += string(myRune[rand.Intn(lenRune)+1])
	}
	return result
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateInt() string {
	password := rand.Int()
	return strconv.Itoa(password)
}

func checkCommand() {
	fmt.Println("Привет, я генератор паролей! Напиши 1, если хочешь сделать пароль.")
	getPass := 0
	fmt.Scanln(&getPass)

	if getPass == 1 {
		fmt.Println("Начинаю генерацию")
		GeneratePassword()
	} else {
		fmt.Println("Попробуйте ещё раз..")
		checkCommand()
	}
}
