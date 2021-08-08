package handler

import (
	"moyu/helper"
	"moyu/transaction"
	"moyu/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter dan mapping input struct
// panggil service, input struct sebagai parameter
// service, dengan campaign id bisa panggil repo
// repo mencari data transaction suatu campaign

type transactionHanler struct {
	service transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHanler {
	return &transactionHanler{transactionService}
}

func (h *transactionHanler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormatter := transaction.FormatCampaignTransactions(transactions)

	response := helper.APIResponse("Successfuly to get campaign's transaction", "success", http.StatusOK, transactionsFormatter)
	c.JSON(http.StatusOK, response)
}

// GetUserTransaction
// handler
// ambil nilai user dari jwt atau middleware
// service
// repository => ambil data transaction (preload campaign)

func (h *transactionHanler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transaction", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userTransactionsFormatter := transaction.FormatUserTransactions(transactions)

	response := helper.APIResponse("Successfuly to get user's transaction", "success", http.StatusOK, userTransactionsFormatter)
	c.JSON(http.StatusOK, response)
}

//	ada input user jumlah amount
// handler tangkap input user -> mapping ke input struct
// call service for transaction, panggil midtran dan daftarkan transaksi
// call repo -> create new transaction data

func (h *transactionHanler) CreateNewTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorFormatter := helper.ErrorValidationFormat(err)

		errorMessage := gin.H{"errors": errorFormatter}

		response := helper.APIResponse("Create new transaction is failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Create new transaction is failed", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Successfuly to create new transaction", "success", http.StatusOK, transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHanler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", "error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
