package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"james": {
		AuthToken: "212EFR",
		Username:  "James",
	},
	"vega": {
		AuthToken: "839IOU",
		Username:  "Vega",
	},
	"chris": {
		AuthToken: "643MRF",
		Username:  "Chris",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"james": {
		Coins:    70,
		Username: "James",
	},
	"vega": {
		Coins:    100,
		Username: "Vega",
	},
	"chris": {
		Coins:    50,
		Username: "Chris",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
