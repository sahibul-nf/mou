package transaction

import "time"

type TransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatTransactions(transactions []Transaction) []TransactionFormatter {

	if len(transactions) == 0 {
		return []TransactionFormatter{}
	}

	var transactionsFormatter []TransactionFormatter

	for _, transaction := range transactions {
		transactionFormatter := FormatTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transactions Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transactions.ID
	formatter.Amount = transactions.Amount
	formatter.Status = transactions.Status
	formatter.CreatedAt = transactions.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transactions.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transactions.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transactions.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactionsFormatter []UserTransactionFormatter

	for _, userTransaction := range transactions {
		userTransactionFormatter := FormatUserTransaction(userTransaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransactionFormatter)
	}

	return userTransactionsFormatter
}
