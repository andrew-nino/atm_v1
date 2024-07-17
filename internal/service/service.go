package service

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Processing interface {
	DepositProcessing(int, float64, chan error)
	WithdrawProcessing(int, float64, chan error)
	BalanceProcessing(int, chan float64)
}

type Service struct {
	BankAccount
	Processing
	Accounts map[int]*Account
}

func NewService(repository map[int]*Account) *Service {
	return &Service{
		BankAccount: NewAccountService(),
		Processing:  NewProcessingService(repository),
		Accounts:    repository,
	}
}
