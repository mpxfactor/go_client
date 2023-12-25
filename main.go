package main

import (
	"fmt"
	"sync"

	"mpxfactor.com/client-test/api"
)

func main() {
	currencyArr := []string{"BTC", "ETH", "BCH"}
	var waitgroup sync.WaitGroup

	for _, currency := range currencyArr {
		waitgroup.Add(1)

		go func() {
			getCurrencyData(currency)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	} else {
		fmt.Printf("Something went wrong")
	}
}
