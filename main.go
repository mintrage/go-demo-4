package main

import (
	"fmt"
)

func main() {
	// fmt.Println(generatePassword(12))
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или логин")
		return
	}
	// myAccount.generatePassword(12)
	myAccount.acc.outputPassword()

	fmt.Println(myAccount)
	// outputPassword(&myAccount)

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
