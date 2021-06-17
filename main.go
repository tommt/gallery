package main

import (
	views "gallery.com/views"
	"github.com/gorilla/mux"
	"net/http"
)


var (
	homeView *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}


func main(){
	homeView = views.NewView("bulma","views/home.gohtml")
	contactView = views.NewView("bulma", "views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
}

func must(err error){
	if err != nil{
		panic(err)
	}
}