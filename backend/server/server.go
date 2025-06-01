package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
	"io"
	"log"
	"net/http"
)

func main() {
	stripe.Key = "sk_test_51ROH3hP1jm8N24y4KsOvySph8UwEO08W3Ep3n5W1rhm8IMwxHloduk6OoPEk5ZYaeDEmzdEnfcYbjFrrFKiLxL6400SSBtTS5m"
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)
	log.Println("Server started")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	//fmt.Fprint(w, "Endpoint Hit!")

	var req struct {
		ProductID string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address1"`
		Address2  string `json:"address2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductID)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response struct {
		ClientSecret string `json:"clientSecret"`
	}
	response.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = io.Copy(w, &buf)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(req.ProductID)

	//fmt.Println(req.FirstName)
	//fmt.Println(req.LastName)
	//fmt.Println(req.Address1)
	//fmt.Println(req.Address2)
	//fmt.Println(req.City)
	//fmt.Println(req.State)
	//fmt.Println(req.Zip)
	//fmt.Println(req.Country)
	//fmt.Fprint(w, "Endpoint Hit!")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	var response []byte = []byte("Server OK! and running")

	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}
	return 0
}
