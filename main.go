package main

import "mpxfactor.com/client-test/api"

func main() {
	rate, err := api.GetRate("BTC")

	print(rate, err)
}
