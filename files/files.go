package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) {
	// file, err := os.Open("files/files.txt") // Открываем файл и читаем по строкам
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	fmt.Println(string(data))
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
	fmt.Println("Файл успешно сохранен:", name)
}
