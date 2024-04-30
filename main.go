package main

import (
	"fmt"
	"time"
)

func main() {
	birthdayDate := birthday()

	// Получаем текущую дату
	tNow := time.Now()
	fmt.Println("The current date: ", tNow.Format("02.01.2006"))

	// Вычисляем возраст
	age := tNow.Year() - birthdayDate.Year()

	// Проверяем, прошел ли день рождения в этом году
	if tNow.Month() < birthdayDate.Month() || (tNow.Month() == birthdayDate.Month() && tNow.Day() < birthdayDate.Day()) {
		age--
	}

	fmt.Println("Your age: ", age)
}

func birthday() time.Time {
	var birthdayStr string

	for {
		fmt.Print("Enter your birthday in the format dd.mm.yyyy: ")
		_, err := fmt.Scan(&birthdayStr)
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		// Предполагаем, что вводится формат "dd.mm.yy"
		const format = "02.01.2006"
		t, err := time.Parse(format, birthdayStr)
		if err != nil {
			fmt.Println("Incorrect date format. Please enter the date again.")
			continue
		}
		fmt.Println("Your birthday: ", t.Format("02.01.2006"))
		return t
	}
}
