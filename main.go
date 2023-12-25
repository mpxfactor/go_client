package main

import (
	"fmt"

	"mpxfactor.com/client-test/api"
)

func main() {
	rate, err := api.GetRate("BTC")

	if err == nil {
		fmt.Printf("The rate for %v is %.2f", rate.Currency, rate.Price)
	} else {
		fmt.Printf("Something went wrong")
	}
}
