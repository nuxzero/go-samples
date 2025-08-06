package main

import (
	jsonStr "encoding/json"
	"fmt"
	"time"
)

const jsonString = `
{
  "object": "card",
  "id": "card_test_no1t4tnemucod0e51mo",
  "livemode": false,
  "location": "/customers/cust_test_no1t4tnemucod0e51mo/cards/card_test_no1t4tnemucod0e51mo",
  "deleted": false,
  "street1": "1448/4 Praditmanutham Road",
  "street2": null,
  "city": "Bangkok",
  "state": null,
  "phone_number": "0123456789",
  "postal_code": "10320",
  "country": "th",
  "financing": "credit",
  "bank": "Bank of the Unbanked",
  "brand": "Visa",
  "fingerprint": "XjOdjaoHRvUGRfmZacMPcJtm0U3SEIIfkA7534dQeVw=",
  "first_digits": null,
  "last_digits": "4242",
  "name": "Somchai Prasert",
  "expiration_month": 12,
  "expiration_year": 2024,
  "security_code_check": true,
  "tokenization_method": null,
  "created_at": "2019-12-31T12:59:59Z"
}
`

type Card struct {
	Object    string    `json:"object"`
	ID        string    `json:"id"`
	Livemode  bool      `json:"livemode"`
	Location  string    `json:"location"`
	Deleted   bool      `json:"deleted"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	jsonBytes := []byte(jsonString)

	var card Card
	_ = UnmarshalJSON(jsonBytes, &card)

	fmt.Println(card)

	json, err := MarshalJSON(&card)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON:", string(json))
}

func UnmarshalJSON(data []byte, card *Card) error {
	err := jsonStr.Unmarshal(data, card)

	if err != nil {
		return err
	}

	return nil
}

func MarshalJSON(card *Card) ([]byte, error) {
	return jsonStr.Marshal(card)
}
