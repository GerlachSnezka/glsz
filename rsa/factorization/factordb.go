package factorization

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
)

type response struct {
	Id int `json:"id"`
	Status string `json:"status"`
	Factors [][]interface{} `json:"factors"`
}

func FactorDBFactorize(n *big.Int) (*big.Int, *big.Int) {
	url := "http://factordb.com/api/?query=" + n.String()

	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch data from factordb.com, error: %v\n", err)
		return &big.Int{}, &big.Int{}
	}

	defer r.Body.Close()

	var response response
	err = json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		fmt.Printf("Failed to decode response from factordb.com, error: %v\n", err)
		return &big.Int{}, &big.Int{}
	}

	if len(response.Factors) == 0 {
		return &big.Int{}, &big.Int{}
	}

	var factors []*big.Int
	for _, fwn := range response.Factors {
		factor, sucs := big.NewInt(0).SetString(fwn[0].(string), 10)
		if !sucs {
			fmt.Printf("Failed to parse factor %d", factor)
			return &big.Int{}, &big.Int{}
		}

		times := int(fwn[1].(float64))

		for i := 0; i < times; i++ {
			factors = append(factors, factor)
		}
	}

	return factors[0], factors[1]
}