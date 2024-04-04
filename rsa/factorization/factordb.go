package factorization

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/charmbracelet/log"
)

type response struct {
	Id interface{} `json:"id"`
	Status string `json:"status"`
	Factors [][]interface{} `json:"factors"`
}

func FactorDBFactorize(n *big.Int) (*big.Int, *big.Int) {
	url := "http://factordb.com/api/?query=" + n.String()

	r, err := http.Get(url)
	if err != nil {
		log.Error("Failed to fetch data from factordb.com", "error", err)
		return &big.Int{}, &big.Int{}
	}

	defer r.Body.Close()

	var response response
	err = json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		log.Error("Failed to decode response from factordb.com", "error", err)
		return &big.Int{}, &big.Int{}
	}

	if len(response.Factors) == 0 {
		return &big.Int{}, &big.Int{}
	}

	var factors []*big.Int
	for _, fwn := range response.Factors {
		factor, sucs := big.NewInt(0).SetString(fwn[0].(string), 10)
		if !sucs {
			log.Error("Failed to parse factor", "factor", factor)
			return &big.Int{}, &big.Int{}
		}

		times := int(fwn[1].(float64))

		for i := 0; i < times; i++ {
			factors = append(factors, factor)
		}
	}

	return factors[0], factors[1]
}