package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

/*
 * Complete the 'totalTransactions' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER locationId
 *  2. STRING transactionType
 *
 * Base url for copy/paste
 * https://jsonmock.hackerrank.com/api/transactions/search?txnType=<transactionType>&page=<pageNumber>
 */

type LocationObject struct {
	ID      int32  `json:"id"`
	Address string `json:"address"`
	City    string `json:"cit"`
	ZipCode int32  `json:"zipCode"`
}

type TransactionObject struct {
	ID        int32          `json:"id"`
	Timestamp int            `json:"timestamp"`
	UserID    int            `json:"userId"`
	Username  string         `json:"userName"`
	TxnType   string         `json:"txnType"`
	Amount    string         `json:"amount"`
	Location  LocationObject `json:"location"`
	IP        string         `json:"ip"`
}

type HttpResponseBody struct {
	Page       int32               `json:"page"`
	PerPage    int32               `json:"per_page"`
	Total      int32               `json:"total"`
	TotalPages int32               `json:"total_pages"`
	Data       []TransactionObject `json:"data"`
}

func getListTransactionsByPage(transactionType string, page int) (jsonBody HttpResponseBody, err error) {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/transactions/search?txnType=%s&page=%d", transactionType, page)
	resp, err := http.Get(url)
	if err != nil {
		return HttpResponseBody{}, err
	}
	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return HttpResponseBody{}, err
	}

	return jsonBody, nil
}

func getListTransactions(transactionType string) ([]TransactionObject, error) {
	var list []TransactionObject

	resp, err := getListTransactionsByPage(transactionType, 1)
	if err != nil {
		return []TransactionObject{}, err
	}
	totalPages := resp.TotalPages
	list = append(list, resp.Data...)

	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 2; i <= int(totalPages); i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			resp, err := getListTransactionsByPage(transactionType, page)
			if err != nil {
				return
			}
			lock.Lock()
			list = append(list, resp.Data...)
			lock.Unlock()
		}(i)
	}
	wg.Wait()
	return list, nil
}

func currencyToFloat64(s string) float64 {
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.
	}
	return num
}

func totalTransactions(locationId int32, transactionType string) [][]int32 {
	txns, err := getListTransactions(transactionType)
	if err != nil {
		return [][]int32{{-1, -1}}
	}

	userIDs := []int{}
	userTotalAmount := make(map[int]float64)
	for _, txn := range txns {
		if txn.Location.ID == locationId {
			amount := currencyToFloat64(txn.Amount)
			userID := txn.UserID
			if _, ok := userTotalAmount[userID]; !ok {
				userIDs = append(userIDs, userID)
				userTotalAmount[userID] = amount
			} else {
				userTotalAmount[userID] += amount
			}
		}
	}
	sort.Ints(userIDs)

	var out [][]int32
	for _, id := range userIDs {
		amount := int32(math.Floor(userTotalAmount[id]))
		userID := int32(id)
		out = append(out, []int32{userID, amount})
	}
	return out
}

func main() {
	txns := totalTransactions(1, "debit")
	fmt.Println(txns)
}
