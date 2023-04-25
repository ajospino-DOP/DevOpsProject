package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	name string
	value string
}

var savedInfo []*Info = []*Info{}

func index(w http.ResponseWriter, r *http.Request) {
	if(r.Method != http.MethodGet){
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	fmt.Fprintf(w, "Welcome visitor")
}

func info(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
	case http.MethodGet:
		fmt.Fprintf(w, "Welcome to the X application, simply send values to save with your request or reload this page to see past values")
		listInfo(w, r)
	case http.MethodPost:
		sendInfo(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}	

func listInfo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(savedInfo)
}

func sendInfo(w http.ResponseWriter, r *http.Request){
	newInfo := &Info{}
	err := json.NewDecoder(r.Body).Decode(newInfo)
	if(err != nil){
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	savedInfo = append(savedInfo, newInfo)
}