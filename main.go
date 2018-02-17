package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"repayment"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessing index")
	readme, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(readme))
}

func generatePlan(w http.ResponseWriter, r *http.Request) {
	log.Println("Invoking API: generate-plan")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	p := repayment.PVPlanJSON(body)
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/generate-plan", generatePlan)
	port := ":8080"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
