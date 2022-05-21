package utils

import "github.com/D4vecode/personal_transactions/models"

type CalculatedData struct {
	Total float64
	AvgCredit float64
	AvgDebit float64
	NumTransactionsByMonth []TransactionsInMonth
}

type TransactionsInMonth struct {
	Month string
	NumTransactions int
}

func CalculateData(transactions []*models.Transaction) CalculatedData {
	var totalCredit, totalDebit float64
	var totalCreditTransactions, totalDebitTransactions int
	var numTransactionsInMonth []TransactionsInMonth

	for _, transaction := range transactions {
		if transaction.Type == "credit" {
			totalCredit += transaction.Amount
			totalCreditTransactions++
		} else {
			totalDebit += transaction.Amount
			totalDebitTransactions++
		}
		if len(numTransactionsInMonth) == 0 {
			numTransactionsInMonth = append(numTransactionsInMonth, TransactionsInMonth{Month: transaction.Date.Month().String(), NumTransactions: 1})
		} else {
			for i, month := range numTransactionsInMonth {
				if month.Month == transaction.Date.Month().String() {
					numTransactionsInMonth[i].NumTransactions++
					break
				} else if i == len(numTransactionsInMonth) - 1 {
					numTransactionsInMonth = append(numTransactionsInMonth, TransactionsInMonth{Month: transaction.Date.Month().String(), NumTransactions: 1})
					break
				}
			}
		}
	}

	return CalculatedData{
		Total: totalCredit - totalDebit,
		AvgCredit: totalCredit / float64(totalCreditTransactions),
		AvgDebit: -(totalDebit / float64(totalDebitTransactions)),
		NumTransactionsByMonth: numTransactionsInMonth,
	}
}