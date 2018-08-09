package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var visitors int64

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	http.HandleFunc("/hi2", handleHi2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	//if match, _ := regexp.MatchString(`\w+$`, r.FormValue("color")); !match {
	//	http.Error(w, "color is valid", http.StatusBadRequest)
	//	return
	//}
	visitorNum := atomic.AddInt64(&visitors, 1)
	//visitors++
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(visitorNum) + "!"))
}

func handleHi2(w http.ResponseWriter, r *http.Request) {
	//if match, _ := regexp.MatchString(`\w+$`, r.FormValue("color")); !match {
	//	http.Error(w, "color is valid", http.StatusBadRequest)
	//	return
	//}
	//visitors = atomic.AddInt64(&visitors, 1)
	visitors++
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome2!</h1>You are visitor number " + fmt.Sprint(visitors) + "!"))
}
