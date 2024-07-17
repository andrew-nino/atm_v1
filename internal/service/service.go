package service

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Service struct {
	BankAccount
	Accounts map[int]*Account
}

func NewService(repository map[int]*Account) *Service {
	return &Service{
		BankAccount: NewAccountService(),
		Accounts:    repository,
	}
}
