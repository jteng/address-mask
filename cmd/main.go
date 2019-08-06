package main

import (
	"addressmask"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ServeMe)
	r.Handle("/deposits/accounts/{accountReferenceId}/address/{addressId}", NewAddressHandler())
	srv := &http.Server{
		WriteTimeout: 1 * time.Second,
		Handler:      r,
		Addr:         ":3333",
	}
	log.Print("starting.....")
	log.Fatal(srv.ListenAndServe())
}

func ServeMe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

type AddressHandler struct {
	AddressBook []addressmask.Address
	randGen     *rand.Rand
}

func NewAddressHandler() *AddressHandler {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return &AddressHandler{
		randGen: r1,
		AddressBook: []addressmask.Address{
			{AddressLine1: "line1", AddressLine2: "line2", City: "Phila"},
			{AddressLine1: "addrLine1", AddressLine2: "addrLine2", City: "Wilmington", State: "DE"},
			{AddressLine1: "l1", AddressLine2: "l2", City: "Trenton", State: "NJ"},
			{AddressLine1: "1", AddressLine2: "2", City: "McLean", State: "VA"},
		},
	}
}

func (h *AddressHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	waitTime := time.Duration(h.randGen.Intn(50)) * time.Millisecond
	log.Printf("wait %v", waitTime)
	time.Sleep(waitTime)
	idx := h.randGen.Intn(len(h.AddressBook) - 1)
	addr := h.AddressBook[idx]
	data, err := json.Marshal(addr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
	return
}
