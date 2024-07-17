package controller

import (
	"net/http"
	"sync"

	"github.com/andrew-nino/atm_v1/internal/service"

	"github.com/gin-gonic/gin"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var mu sync.Mutex

func (h *Handler) addAccount(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	nextID := intSeq()

	acc := service.Account{
		Id:      nextID(),
		Balance: 0,
	}

	h.services.Accounts[acc.Id] = &acc

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})
}
