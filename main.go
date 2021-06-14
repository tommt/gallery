package main

import (
	"fmt"
	"github.com/gorilla/mux"
)
import "net/http"

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/"{
		fmt.Println("Some one looked")
		fmt.Fprintf(w, "<h1>Welcome Daju, Sanchai ho</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(w, "To get in touch please send email to <a href=\"mailto:atombombmount@gmail.com\">atombombmount@gmail.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>Page Not Found 404 :(</h1>")
	}


}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", r)
}