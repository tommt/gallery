package main

import (
	"fmt"
	"log"
)
import "net/http"
import "github.com/julienschmidt/httprouter"


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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main(){
	//mux := &http.ServeMux{}
	//mux.HandleFunc("/", handlerFunc)
	//http.HandleFunc("/", handlerFunc)


	//http.ListenAndServe(":3000", router)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":3000", router))
}