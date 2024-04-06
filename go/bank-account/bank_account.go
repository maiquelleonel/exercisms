package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	sync.RWMutex
	balance int64
	isOpen  bool
}

func Open(amount int64) *Account {

	if amount < 0 {
		return nil
	}

	return &Account{balance: amount, isOpen: true}
}

func (a *Account) Balance() (ret int64, ok bool) {

	a.RLock()
	defer a.RUnlock()

	if !a.isOpen {
		return
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (ret int64, ok bool) {

	a.Lock()
	defer a.Unlock()

	if a.isOpen && a.balance+amount >= 0 {
		a.balance += amount
		return a.balance, true
	}
	return

}

func (a *Account) Close() (ret int64, ok bool) {

	a.Lock()
	defer a.Unlock()

	if a.isOpen {
		curBalance := a.balance
		a.balance, a.isOpen = 0, false
		return curBalance, true
	}
	return

}
