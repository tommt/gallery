package main

import "fmt"
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Welcome Daju, Sanchai ho</h1>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}