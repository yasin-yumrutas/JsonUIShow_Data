// Bu Go kodu, basit bir JSON API servisi oluşturur.
// /api/items yolunda GET isteği geldiğinde, örnek öğeleri JSON formatında döndürür.

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Item 1", Price: 100},
	{ID: 2, Name: "Item 2", Price: 200},
	// ... diğer öğeler
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/api/items", getItems)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
