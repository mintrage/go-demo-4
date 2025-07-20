package main

import (
	"fmt"
)

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	// files.WriteFile("Привет! Я файл", "file.txt")
	// login := promptData("Введите логин")
	// password := promptData("Введите пароль")
	// url := promptData("Введите URL")

	// myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	// if err != nil {
	// 	fmt.Println("Неверный формат URL или логин")
	// 	return
	// }
	// myAccount.OutputPassword()
	// fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

// func generatePassword(n int) string {
// 	res := make([]rune, n)
// 	for i := range res {
// 		res[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	return string(res)
// }
