package service

import (
	"fmt"
	"sync"
)

type Account struct {
	Id      int `json:"id"`
	Balance float64 `json:"balance"`
	mutex   sync.Mutex
}

func NewAccountService() BankAccount{
	return &Account{}
}

func (acc *Account) Deposit(amount float64) error {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	acc.Balance += amount
	return nil
   }
   
   func (acc *Account) Withdraw(amount float64) error {
	acc.mutex.Lock()
	defer acc.mutex.Unlock()
	if acc.Balance < amount {
	 return fmt.Errorf("insufficient funds")
	}
	acc.Balance -= amount
	return nil
   }
   
   func (acc *Account) GetBalance() float64 {
	return acc.Balance
   }