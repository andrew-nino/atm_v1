package service

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"

)

var mu sync.Mutex
// Associates the Account structure, which implements the BankAccount interface, with map[int]*Account, which simulates a repository.
type ProcessingService struct {
	repository map[int]*Account
}

func NewProcessingService(accounts map[int]*Account) *ProcessingService {
	return &ProcessingService{repository: accounts}
}

// Using a mutex we check the existence of an account.
// If there is, we send it to the goroutine for processing. We are waiting for a possible error from the channel.
func (p *ProcessingService) DepositProcessing(id int, amount float64, resultChan chan error) {
	mu.Lock()
	acc, ok := p.repository[id]
	mu.Unlock()

	if !ok {
		resultChan <- fmt.Errorf("account not found")
		return
	}
	go func() {
		err := acc.Deposit(amount)
		resultChan <- err
		log.Printf("The deposit transaction for account %d was successful\n", id)
	}()
}

// Using a mutex we check the existence of an account.
// If there is, we send it to the goroutine for processing. We are waiting for a possible error from the channel.
func (p *ProcessingService) WithdrawProcessing(id int, amount float64, resultChan chan error) {
	mu.Lock()
	acc, ok := p.repository[id]
	mu.Unlock()

	if !ok {
		resultChan <- fmt.Errorf("account not found")
		return
	}

	go func() {
		err := acc.Withdraw(amount)
		resultChan <- err
		log.Printf("The withdraw transaction for account %d was successful\n", id)
	}()
}

// Using a mutex we check the existence of an account.
// If there is, we send it to the goroutine for processing. We are waiting for the result from the channel.
func (p *ProcessingService) BalanceProcessing(id int, resultChan chan float64) {
	mu.Lock()
	acc, ok := p.repository[id]
	mu.Unlock()

	if !ok {
		resultChan <- 0.0
		return
	}

	go func() {
		balance := acc.GetBalance()
		resultChan <- balance
		log.Printf("Receiving balance for account %d was successful.\n", id)
	}()
}
