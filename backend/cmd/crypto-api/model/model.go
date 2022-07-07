package model

// In the first part of the file we are mapping requests and responses to their JSON payload.
type CryptoRequest struct {
	Symbol string
}

type CryptoCurrencyResp struct {
	Id          string `json:"id"`
	FullName    string `json:"fullName"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurency"`
}

type CryptoCurrenciesResp struct {
	Currencies []CryptoCurrencyResp `json:"currencies"`
}
