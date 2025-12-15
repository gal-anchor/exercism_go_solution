package account

import "sync"

// Account represents a bank account.
type Account struct {
	mu      sync.Mutex
	balance int64
	open    bool
}

// Open creates a new account with an initial amount.
// Returns nil if the amount is negative.
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
		open:    true,
	}
}

// Balance returns the current balance and whether the account is open.
func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}
	return a.balance, true
}

// Deposit adds amount to the balance.
// Returns the new balance and whether the operation succeeded.
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open || a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

// Close closes the account and returns the remaining balance.
func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	a.open = false
	return a.balance, true
}
