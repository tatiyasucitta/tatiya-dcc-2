package main

import (
	"fmt"
	"net/http"
)
func pageOne(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Ini page utama\nhalo ges")
}

func pageTwo(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Sekarang page 2\nHEHEHEHEHEHEHHE")
}

func main(){
	http.HandleFunc("/", pageOne)
	http.HandleFunc("/page-two", pageTwo)

	http.ListenAndServe("", nil)
}