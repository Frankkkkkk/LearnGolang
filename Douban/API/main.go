package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",Index)
	router.HandleFunc("/todos",TodoIndex)
	router.HandleFunc("/todos/{todoId}",TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprint(w, "Todo show:",todoId)

}