package main

import (
	"fmt"
	"github.com/gorilla/mux"
)
import "net/http"


func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "To get in touch please send email to <a href=\"mailto:atombombmount@gmail.com\">atombombmount@gmail.com</a>")
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}