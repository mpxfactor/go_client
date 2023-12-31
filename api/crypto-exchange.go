package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

	var response CexResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	currencyPrice, _ := strconv.ParseFloat(response.Last, 2)
	// rate := datatypes.Rate{Currency: currency, Price: float64(response.Bid)}
	rate := datatypes.Rate{Currency: currency, Price: currencyPrice}
	return &rate, nil
}
