package service

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"

)

var mu sync.Mutex

type ProcessingService struct {
	repository map[int]*Account
}

func NewProcessingService(accounts map[int]*Account) *ProcessingService {
	return &ProcessingService{repository: accounts}
}

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
