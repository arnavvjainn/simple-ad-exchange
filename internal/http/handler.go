package http

import (
	"net/http"
	"encoding/json"
	"simple-ad-exchange/internal/auction"
)

func handleBidRequest(w http.ResponseWriter, r *http.Request) {
	var bidRequest auction.BidRequest
	if err := json.NewDecoder(r.Body).Decode(&bidRequest); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
	response := auction.PerformAuction(bidRequest)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
