package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

// Принимает указатель на структуру account/Функцию превратили в метод структуры
func (acc *account) outputPassword() {
	fmt.Println("Логин:", (*acc).login)
	fmt.Println("Пароль:", (*acc).password)
	fmt.Println("URL:", (*acc).url)
}

// Генерирует пароль
func (acc *account) generatePassword(length int) {
	res := make([]rune, length)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	(*acc).password = string(res)
}

// Конструктор структуры
//
//	Если логина нет, то возвращается ошибка
//
// Если нет пароля, то генерируется пароль
// func newAccount(login, password, urlString string) (*account, error) {
// 	/// Тут можно добавить проверку на валидность данных
// 	if login == "" {
// 		return nil, errors.New("логин не может быть пустым")
// 	}

// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("некорректный URL")
// 	}

// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}

// 	if password == "" {
// 		newAcc.generatePassword(10)
// 	}

// 	return newAcc, nil
// }

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	/// Тут можно добавить проверку на валидность данных
	if login == "" {
		return nil, errors.New("логин не может быть пустым")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("некорректный URL")
	}

	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}

	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	// str := []rune("Hello, World!")
	// for _, char := range string(str) {
	// 	fmt.Println(char, string(char))
	// }

	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}
