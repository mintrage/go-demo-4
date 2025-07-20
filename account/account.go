package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	// fmt.Println(acc)
	// fmt.Println(acc.login, acc.password, acc.url)
	color.Cyan(acc.login)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// 1. Если логина нет, то ошибка
// 2. Если нет пароля, то генерим
func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login:    login,
			password: password,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, err
}
