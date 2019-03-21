package main

import (
	"fmt"
	"net/http"
)

var stroredValuse = make(map[string]string)

func main() {
	http.HandleFunc("/insertValue", insertValue)
	http.HandleFunc("/getValue", getValueForKey)
	http.ListenAndServe(":8666", nil)
}

func insertValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	stroredValuse[key] = value
	fmt.Fprintf(w, "Value was stored in map.")
}

func getValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, stroredValuse[r.URL.Query().Get("key")])
}
