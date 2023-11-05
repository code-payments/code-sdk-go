package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	codesdk "github.com/code-wallet/code-sdk-go"
)

func main() {
	client := codesdk.NewWebClient()

	http.HandleFunc("/", serveStaticFiles)

	// Route to create a payment intent
	http.HandleFunc("/create-intent", func(w http.ResponseWriter, r *http.Request) {
		resp, err := func() (*codesdk.CreatePaymentRequestResponse, error) {
			// Specify payment request details
			intent, err := codesdk.NewPaymentRequestIntent(
				// Or the string "usd"
				codesdk.USD,
				// Minimum amount is $0.05 USD
				0.05,
				// Code Deposit Address or any Kin token account
				"E8otxw1CVX9bfyddKu3ZB3BVLa4VVF9J7CTPdnUwT9jR",
			)
			if err != nil {
				return nil, err
			}

			// Create a payment request intent
			return client.CreatePaymentRequest(r.Context(), intent)
		}()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		marshalledJson, _ := json.Marshal(resp)
		w.Header().Add("Content-Type", "application/json")
		w.Write(marshalledJson)
	})

	// Route to verify a payment intent
	http.HandleFunc("/verify/", func(w http.ResponseWriter, r *http.Request) {
		// Get the intent status for verification
		intentId := strings.TrimPrefix(r.URL.Path, "/verify/")
		resp, err := client.GetIntentStatus(r.Context(), intentId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		marshalledJson, _ := json.Marshal(resp)
		w.Header().Add("Content-Type", "application/json")
		w.Write(marshalledJson)
	})

	http.ListenAndServe(":3000", nil)
}

func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	if name == "/" {
		name = "/index.html"
	}

	http.ServeFile(w, r, fmt.Sprintf("./static%s", name))
}
