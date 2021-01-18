package accounts

import (
	"errors"
	"fmt"
)

// Account - export해서 쓰려면 내부 내용도 대문자로 시작 (private, public 설정)
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw, you are poor")

// NewAccount - constructor
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 복사본없이 리턴하기 위해
}

// Deposit - method -> func (receiver)
// Object를 수정하기 때문에 메소드 리시버로 포인터를 받음
// 포인터로 안받으면 복사본에다가 수정함
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance - get balance
func (a Account) Balance() int {
	return a.balance
}

// Withdraw - error의 두가지 종류 -> New Error OR nil(null)
// 그렇다고 GO에서 에러를 throw하지 않음 (별도로 구현해야함)
/*
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err) // log.Fatalln -> println & Exit(1)
	}
*/
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// go에서 자동으로 지원, struct를 호출하면 String 자동 호출
func (a Account) String() string {
	// sprint -> format of string
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
