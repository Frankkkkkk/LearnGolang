package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct{
	Id int `json:"Id"`
	Tittle string `jason:"Tittle"`
	Desc string `jason:"Desc"`
	Content string `jason:"Content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	articles :=Articles{
		Article{Tittle:"Test Tittle",Desc:"Test Desc", Content:"Test Content"},
		Article{Tittle:"Tittle",Desc:"Desc", Content:"Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprint(w, "key:"+key)
}

func homePage(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"Homepage Endpoint Hit")
	fmt.Println("Endpoint Hit: homePage")
}


func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/all",returnAllArticles)
	http.HandleFunc("/articles/{id}",returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func main(){
	handleRequest()
}
