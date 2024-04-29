package main

import (
	"fmt"
	"time"
)

func main() {
	var birthdayDate string
	for {
		fmt.Print("Enter your birthday in the format dd.mm.yyyy: ")
		_, err := fmt.Scanln(&birthdayDate)
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		// Предполагаем, что вводится формат "dd.mm.yy"
		const format = "02.01.2006"
		t, err := time.Parse(format, birthdayDate)
		if err != nil {
			fmt.Println("Incorrect date format. Please enter the date again.")
			continue
		}
		fmt.Println("Your birthday: ", t.Format("02.01.2006"))

		// Получаем текущую дату
		tNow := time.Now()
		fmt.Println("The current date: ", tNow.Format("02.01.2006"))

		// Вычисляем возраст
		age := tNow.Year() - t.Year()

		// Проверяем, прошел ли день рождения в этом году
		if tNow.YearDay() < t.YearDay() {
			age--
		}

		fmt.Println("Your age: ", age)

		break
	}
}
