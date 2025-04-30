package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	// str := []rune("Hello, World!")
	// for _, char := range string(str) {
	// 	fmt.Println(char, string(char))
	// }

	// files.ReadFile("test.txt")
	// files.WriteFile("Test GO", "test.txt")

	createAccount()
}

func createAccount() {
	login := promptData("Введите логин:")
	password := promptData("Введите пароль:")
	url := promptData("Введите URL:")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Ошибка при сериализации в байты:", err)
		return
	}
	files.WriteFile(file, "account.json")
}

func promptData(prompt string) string {
	fmt.Println(prompt)
	var result string
	fmt.Scanln(&result)
	return result
}
