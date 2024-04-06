package account

import "sync"

// Define the Account type here.
type Account struct {
	mu      sync.Mutex
	balance int64
	isOpen  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, isOpen: true}
}

func (a *Account) Balance() (int64, bool) {
	if a.isOpen {
		return a.balance, true
	}
	return 0, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	var ret int64 = 0
	var ok bool = false
	if a.isOpen && a.balance > 0 {
		a.mu.Lock()
		a.balance += amount
		ret, ok = a.balance, true
		a.mu.Unlock()
	}
	return ret, ok
}

func (a *Account) Close() (int64, bool) {
	var ret int64 = 0
	var ok bool = false
	a.mu.Lock()
	if a.isOpen {
		curBalance := a.balance
		a.balance = 0
		a.isOpen = false
		ret, ok = curBalance, true
	}
	defer a.mu.Unlock()
	return ret, ok
}
