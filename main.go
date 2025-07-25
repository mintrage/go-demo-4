package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// output.PrintError(1)
	// output.PrintError("sd")
	fmt.Println("___Менеджер паролей___")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	//vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

// func getMenu() int {
// 	var variant int
// 	fmt.Println("Выберите вариант:")
// 	fmt.Println("1. Создать аккаунт")
// 	fmt.Println("2. Найти аккаунт")
// 	fmt.Println("3. Удалить аккаунт")
// 	fmt.Println("4. Выход")
// 	fmt.Scanln(&variant)
// 	return variant
// }

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}

}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или логин")
		return
	}
	vault.AddAccount(*myAccount)
}

// func promptData(prompt string) string {
// 	fmt.Print(prompt + ": ")
// 	var res string
// 	fmt.Scanln(&res)
// 	return res
// }

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
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
