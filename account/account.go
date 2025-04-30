package account

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Принимает указатель на структуру account/Функцию превратили в метод структуры
func (acc *Account) OutputPassword() {
	color.Green("Логин:", acc.Login)
	color.Green("Пароль:", acc.Password)
	color.Green("URL:", acc.URL)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Генерирует пароль
func (acc *Account) generatePassword(length int) {
	res := make([]rune, length)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
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

func NewAccount(login, password, urlString string) (*Account, error) {
	/// Тут можно добавить проверку на валидность данных
	if login == "" {
		return nil, errors.New("логин не может быть пустым")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("некорректный URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		URL:       urlString,
	}
	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}
