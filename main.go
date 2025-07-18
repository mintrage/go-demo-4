package main

import (
	"fmt"
)

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1 := account{
		login,
		password,
		url,
	}

	account2 := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(login, password, url)

}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
