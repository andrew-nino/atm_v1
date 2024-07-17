package controller

import (
	"net/http"
	"strconv"

	"github.com/andrew-nino/atm_v1/internal/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)
// Closing function
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type inputData struct {
	Amount float64 `json:"amount"`
}

// Adding a new account to map is an imitation of a repository.
func (h *Handler) addAccount(c *gin.Context) {
// The closure is used as a sequence number for identification.
	nextID := intSeq()

	acc := service.Account{
		Id:      nextID(),
		Balance: 0,
	}
	h.services.Accounts[acc.Id] = &acc
	log.Printf("New account created with ID: %d", acc.Id)
	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
}

// Depositing funds. The received data is checked for correctness and sent to the goroutine for processing.
// The result of the operation is received through the channel.
func (h *Handler) deposit(c *gin.Context) {

	paramStr := c.Param("id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}

	var input inputData
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	resultChan := make(chan error)
	go h.services.DepositProcessing(id, input.Amount, resultChan)
	err = <-resultChan
	if err != nil {
		log.Println("Deposit error:", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "The deposit transaction was successful."})
}

// Withdraw funds. The received data is checked for correctness and sent to the goroutine for processing.
// The result of the operation is received through the channel.
func (h *Handler) withdraw(c *gin.Context) {

	paramStr := c.Param("id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}

	var input inputData
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	resultChan := make(chan error)
	go h.services.WithdrawProcessing(id, input.Amount, resultChan)
	err = <-resultChan
	if err != nil {
		log.Println("Withdraw error:", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "The withdraw transaction was successful."})
}

// Getting balance. The received data is checked for correctness and sent to the goroutine for processing.
// The result of the operation is received through the channel.
func (h *Handler) getBalance(c *gin.Context) {

	paramStr := c.Param("id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}
	id, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}

	resultChan := make(chan float64)
	go h.services.BalanceProcessing(id, resultChan)
	balance := <-resultChan

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
