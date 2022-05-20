package utils

import (
	"fmt"
	"github.com/D4vecode/personal_transactions/models"
	"math"
	"strconv"
	"strings"
	"time"
)

func getTransactionType(transactionType string) (string, float64) {
	amount, err := strconv.ParseFloat(transactionType, 64)

	if err != nil {
		panic(err)
	}
	
	if amount < 0 {
		return "debit", math.Abs(float64(amount))
	}
	return "credit", float64(amount)
}

func parseDate(date string) (time.Time, error) {
	dateNow := time.Now()
	splitedDate := strings.Split(date, "/")
	newDate := strconv.Itoa(dateNow.Year()) + "-" + splitedDate[0] + "-" + splitedDate[1] + "T00:00:00Z"
	parsedDate, err := time.Parse(time.RFC3339, newDate)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

func ProcessData(transactionsArray [][]string) []*models.Transaction {
	transactionsArray = transactionsArray[1:]
	transactions := []*models.Transaction{}

	for _, transaction := range transactionsArray {
		id, _ := strconv.Atoi(transaction[0])
		transactionType, amount := getTransactionType(transaction[2])
		date, err := parseDate(transaction[1])

		if err != nil {
			fmt.Println(err)
			break
		}

		processedTransaction := &models.Transaction{
			ID:        id,
			Date:      date,
			Type:      transactionType,
			Amount:    amount,
		}
		transactions = append(transactions, processedTransaction)

	}
	return transactions
}
