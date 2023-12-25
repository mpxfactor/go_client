package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"mpxfactor.com/client-test/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		json := string(bodyBytes)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: 100}
	return &rate, nil
}
