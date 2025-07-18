package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}
	return &account{
		login:    login,
		password: password,
		url:      urlString,
	}, nil

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

func main() {
	// fmt.Println(generatePassword(12))
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL")
		return
	}
	myAccount.generatePassword(12)
	myAccount.outputPassword()

	fmt.Println(myAccount)
	// outputPassword(&myAccount)

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

// func generatePassword(n int) string {
// 	res := make([]rune, n)
// 	for i := range res {
// 		res[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	return string(res)
// }
