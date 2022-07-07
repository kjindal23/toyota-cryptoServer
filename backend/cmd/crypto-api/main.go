package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kjindal23/toyota-cryptoServer/backend/cmd/crypto-api/model"
)

func main() {
	log.Println("starting Crypto API server")
	//create a new router
	router := mux.NewRouter()
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/currency/{symbol}", getCryptoPrice).Methods("GET")
	router.HandleFunc("/currencies/all", getCryptosPrices).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func getCryptoPrice(w http.ResponseWriter, r *http.Request) {
	log.Println("entering getCryptoCurrency end point")

	requestVal := mux.Vars(r)
	currencySymbol := requestVal["symbol"]

	if !isSymbolValid(currencySymbol) {
		w.Write([]byte(`{"error": {"message": "bad request"}}`))

		return
	}

	url := fmt.Sprintf("https://api.hitbtc.com/api/2/public/ticker/" + currencySymbol)

	resp, err := http.Get(url)
	if err != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))
	}

	defer resp.Body.Close()

	responseData, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))
	}

	var cryptoCurrencyResponse model.CryptoCurrencyResp

	//Unmarshal response
	err = json.Unmarshal(responseData, &cryptoCurrencyResponse)
	if err != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// marshal response before sending
	jsonResponse, err := json.Marshal(cryptoCurrencyResponse)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func getCryptosPrices(w http.ResponseWriter, r *http.Request) {
	log.Println("entering getCryptoCurrencies end point")

	url := "https://api.hitbtc.com/api/2/public/ticker"

	resp, err := http.Get(url)
	if err != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))
	}

	defer resp.Body.Close()

	responseData, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))
	}

	var cryptoResponses model.CryptoCurrenciesResp

	//Unmarshal response
	err = json.Unmarshal(responseData, &cryptoResponses.Currencies)
	if err != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// marshal response before sending
	jsonResponse, err := json.Marshal(cryptoResponses)
	if err != nil {
		w.Write([]byte(`{"error": {"message": "internal server error"}}`))

		return
	}

	w.Write(jsonResponse)
}

func isSymbolValid(symbol string) bool {
	validcurrency := []string{"BTCUSD", "ETHBTC"}
	for _, valid := range validcurrency {
		if valid == symbol {
			return true
		}
	}
	return false
}
